package Route

import (
	RouterURL "MarkDownAPI/Helper/RouteStore"
	"MarkDownAPI/Package/Controller"

	"github.com/gorilla/mux"
)

func CustomRouter(Router *mux.Router) {
	Router.HandleFunc(RouterURL.AddFileURL, Controller.AddFile).Methods("POST")
	Router.HandleFunc(RouterURL.GetFileURL, Controller.GetFileByID).Methods("GET")
	Router.HandleFunc(RouterURL.GetAllFilesURL, Controller.GetAllFileByID).Methods("GET")
	Router.HandleFunc(RouterURL.GetRenderedFilesURL, Controller.GetRenderedFileByID).Methods("GET")
}
