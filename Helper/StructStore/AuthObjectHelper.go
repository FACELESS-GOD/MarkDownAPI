package StructStore

// Login/SiginUP Structs:

type APILogin struct {
	UserName string `json : "username"`
	Password string `json : "password"`
}

type APISignUP struct {
	Name      string `json : "name"`
	Email     string `json : "email"`
	JobTitle  string `json : "title"`
	UserName  string `json : "username"`
	FirstName string `json : "firstname"`
	LastName  string `json : "lastname"`
	Password  string `json : "password"`
}


type TokenStore struct {
	UserID string
	Token string
	DateTime string
}
