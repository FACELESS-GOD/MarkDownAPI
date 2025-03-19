package Route

import (
	RouterURL "MarkDownAPI/Helper/RouteStore"
	"MarkDownAPI/Package/Controller"

	"github.com/gorilla/mux"
)

func CustomRouter(Router *mux.Router) {
	Router.HandleFunc(RouterURL.SiginUPURL, Controller.SignUp).Methods("GET")
	Router.HandleFunc(RouterURL.SiginUPURL, Controller.Login).Methods("GET")
}
