package main

import (
	"MarkDownAPI/Package/Route"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	MuxRouter := mux.NewRouter()

	Route.CustomRouter(MuxRouter)

	http.Handle("/", MuxRouter)

	log.Fatal(http.ListenAndServe("localhost:9030", MuxRouter))

}
