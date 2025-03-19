package Controller

import (
	"MarkDownAPI/Helper/StructStore"
	"MarkDownAPI/Package/Model"
	"MarkDownAPI/Package/Utility"
	"encoding/json"
	"net/http"
)

func SignUp(writer http.ResponseWriter, Req *http.Request) {
	Response := StructStore.GenericMessageReponse{}

	CurrUser := &StructStore.APISignUP{}

	Utility.ParseBody(Req, CurrUser)

	IsUserAdded, err := Model.AddUser(*CurrUser)
	if err != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}

	if IsUserAdded == false {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	} else {
		ValidOperationResponse(writer, Response, "SignUp was successfull")
		return
	}

}

func Login(writer http.ResponseWriter, Req *http.Request) {
	Response := StructStore.GenericMessageReponse{}

	CurrCred := &StructStore.APILogin{}

	Utility.ParseBody(Req, CurrCred)

	TokenStore, IsValid, err := Model.VerifyUser(*CurrCred)
	if err != nil {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	}

	if IsValid == false {
		InvalidOperationResponse(writer, Response, "SignUp was Unsuccessfull")
		return
	} else {
		writer.Header().Set("Authorization", "Bearer "+TokenStore.Token)
		ValidOperationResponse(writer, Response, "SignUp was successfull")
		return
	}

}

// Response Generator function

func ValidOperationResponse(writer http.ResponseWriter, Response StructStore.GenericMessageReponse, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func InvalidOperationResponse(writer http.ResponseWriter, Response StructStore.GenericMessageReponse, message string) {
	Response.Message = message
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(res)
}

func ErrorResponse(writer http.ResponseWriter, Response StructStore.GenericMessageReponse, err error) {
	Response.Message = "Unable to Process"
	res, _ := json.Marshal(Response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	writer.Write(res)

}

//--

// Token
func IsloggedIN(req *http.Request) (bool, error) {
	tokenString := req.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]
	isValid, err := Model.VerifyToken(tokenString)

	if err != nil {
		return false, err
	}
	return isValid, nil

}

//--
