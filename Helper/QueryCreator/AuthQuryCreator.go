package AuthQueryCreator

import "MarkDownAPI/Helper/StructStore"

// ===== ToKen =====

func GetToken(UserID string) string {
	var query string = "SELECT Token, DateTime FROM Token_Storage where UserID = " + UserID

	return query
}

func AddToken(Token string, UserID string, DateTime string) string {
	var query string = "INSERT INTO Token_Storage (UserID, Token, DateTime, ISLoggedIN) VALUES  ('" + UserID + "', '" + Token + "', " + DateTime + ", 1)"
	return query
}

//--

// ====== IS User Exits =====

func IsUserExistsQueryCreator(SiginUPStruct StructStore.APISignUP, LoginStruct StructStore.APILogin, IsLogin bool) string {
	baseString := "SELECT COUNT(*) FROM "
	if IsLogin == true {
		LoginQuery := baseString + "User_Cred where UserName = '" + LoginStruct.UserName
		return LoginQuery
	} else {
		SignUPQuery := baseString + "User_Data where Email = '" + SiginUPStruct.Email
		return SignUPQuery
	}
}

//--

// ===== USER =====

func AddUserDATA(SiginUPStruct StructStore.APISignUP, UserID string) string {
	var query string = "INSERT INTO User_Data (Userid, FirstName, LastName, Email, Title)  VALUES  ('" + UserID + "', '" + SiginUPStruct.FirstName + "', '" + SiginUPStruct.LastName + "', '" + SiginUPStruct.FirstName + "', '" + SiginUPStruct.Email + "', '" + SiginUPStruct.JobTitle + "')"
	return query
}

func AddUserCRED(SiginUPStruct StructStore.APISignUP) string {
	var query string = "INSERT INTO User_Cred (UserName, Password) VALUES  ('" + SiginUPStruct.UserName + "', '" + SiginUPStruct.Password + "'); SELECT LAST_INSERT_ID();"
	return query
}

func CheckUserCRED(UserName string) string {
	var query string = "SELECT Password, Userid FROM User_Cred WHERE UserName = '" + UserName + "'"
	return query
}

func GetUserID(UserName string) string {
	var query string = "SELECT Userid FROM User_Cred WHERE UserName = '" + UserName + "'"
	return query
}

//--
