package main

import (
	"MarkDownAPI/Package/Route"
	"MarkDownAPI/Package/Utility"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	Utility.InitialiseDatabaseConnection()

	Utility.InitialiseRedisConn()

	MuxRouter := mux.NewRouter()

	Route.CustomRouter(MuxRouter)

	http.Handle("/", MuxRouter)

	log.Fatal(http.ListenAndServe("localhost:9030", MuxRouter))

	defer Utility.TerminateDatabaseConnection()

}
