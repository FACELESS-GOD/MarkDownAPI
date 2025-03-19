package Model

import (
	"MarkDownAPI/Helper"
	AuthQueryCreator "MarkDownAPI/Helper/QueryCreator"
	"MarkDownAPI/Helper/StructStore"
	"MarkDownAPI/Package/Utility"
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func IsUserExits(SiginUPStructHelper StructStore.APISignUP, LoginStruct StructStore.APILogin, IsLogin bool) (bool, error) {

	var count int 
	query := AuthQueryCreator.IsUserExistsQueryCreator(SiginUPStructHelper,LoginStruct,IsLogin)

	execute, err := Utility.DatabaseInstace.Query(query)
	if err != nil {
		return false, err
	}

	for execute.Next() {
		err := execute.Scan(&count)
		if err != nil {
			return false,  err
		}
	}

	if  count>0 {
		return true, nil
	} else {
		return false, nil
	}

}

func AddUser(SiginUPStructHelper StructStore.APISignUP) (bool, error) {
	IsUserExist, err := IsUserExits(SiginUPStructHelper, StructStore.APILogin{}, false)

	if err != nil {
		return false, err
	}

	if IsUserExist == true {
		return false, nil
	} else {
		IsCredAdded, UserID, err := AddUserCRED(SiginUPStructHelper)
		if err != nil {
			return false, err
		}

		if IsCredAdded == true {
			go AddUserData(SiginUPStructHelper, UserID)
			return true, nil
		} else {
			return false, nil
		}

	}

}

func AddUserData(SiginUPStructHelper StructStore.APISignUP, UserID string) (bool, error) {
	query := AuthQueryCreator.AddUserDATA(SiginUPStructHelper, UserID)

	execute, err := Utility.DatabaseInstace.Query(query)
	if err != nil {
		return false, err
	}

	for execute.Next() {
	}
	return true, nil

}

func AddUserCRED(SiginUPStructHelper StructStore.APISignUP) (bool, string, error) {
	query := AuthQueryCreator.AddUserCRED(SiginUPStructHelper)
	var UserID string

	execute, err := Utility.DatabaseInstace.Query(query)
	if err != nil {
		return false, "", err
	}

	for execute.Next() {
		err := execute.Scan(&UserID)
		if err != nil {
			return false, "", err
		}
	}

	return true, UserID, nil
}

func VerifyUser(LoginStruct StructStore.APILogin) (StructStore.TokenStore, bool, error) {

	TokenStore := StructStore.TokenStore{}

	IsUserExist, err := IsUserExits(StructStore.APISignUP{}, LoginStruct, true)

	if err != nil {
		return TokenStore, false, err
	}

	if IsUserExist != true {
		return TokenStore, false, nil
	} else {
		query := AuthQueryCreator.CheckUserCRED(LoginStruct.UserName)

		execute, err := Utility.DatabaseInstace.Query(query)

		if err != nil {
			return TokenStore, false, err
		}

		var password string
		var UserID int

		for execute.Next() {
			err := execute.Scan(&password, &UserID)

			if err != nil {
				return TokenStore, false, err
			}
		}

		if password == LoginStruct.Password {

			TokenStore, IsTokenExist, err := GetFromDBToken(strconv.Itoa(UserID))

			if IsTokenExist == true {
				return TokenStore, true, nil
			}

			Token, err := CreateToken(LoginStruct.UserName)
			if err != nil {
				return TokenStore, false, err
			}

			TokenStore.UserID = strconv.Itoa(UserID)

			TokenStore.Token = Token

			currentTime := time.Now()

			TokenStore.DateTime = currentTime.Format("2017-09-07 17:06:06")

			go AddToDBToken(TokenStore)

			return TokenStore, true, nil

		} else {
			return TokenStore, false, nil
		}

	}
}

//--

// ===== Token =====

func AddToDBToken(TokenStore StructStore.TokenStore) (bool, error) {
	query := AuthQueryCreator.AddToken(TokenStore.Token, TokenStore.UserID, TokenStore.DateTime)
	execute, err := Utility.DatabaseInstace.Query(query)
	if err != nil {
		return false, err
	}

	for execute.Next() {
	}

	return true, nil
}

func GetFromDBToken(UserID string) (StructStore.TokenStore, bool, error) {

	TokenStore := StructStore.TokenStore{}

	query := AuthQueryCreator.GetToken(UserID)

	execute, err := Utility.DatabaseInstace.Query(query)

	if err != nil {
		return TokenStore, false, err
	}
	var Token string
	var DateTime time.Time

	for execute.Next() {
		err := execute.Scan(&Token, &DateTime)
		if err != nil {
			return TokenStore, false, err
		}
	}
	currTime := time.Now()
	diff := currTime.Sub(currTime)

	if diff.Hours() < 2 {
		TokenStore.Token = Token
		TokenStore.DateTime = DateTime.String()
		TokenStore.UserID = UserID
		return TokenStore, true, nil
	} else {
		return TokenStore, false, nil
	}

}

func CreateToken(UserName string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"UserName": UserName,
			"exp":      time.Now().Add(time.Hour * 1).Unix(),
		})
	tokenString, err := token.SignedString(Helper.SecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(TokenString string) (bool, error) {

	token, err := jwt.Parse(
		TokenString,
		func(token *jwt.Token) (interface{}, error) {
			return Helper.SecretKey, nil
		})

	if err != nil {
		log.Printf("Invalid Token")
		return false, err
	}

	if token.Valid != true {
		return false, err
	}
	return true, nil
}

//--
