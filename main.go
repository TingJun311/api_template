package main

import (
	"fmt"
	"net/http"
	Config "pkg/config"
	"github.com/gorilla/mux"
	Rest "pkg/rest_routes_handlers"
	Auth "pkg/globals/authentications"
	GraphQL "pkg/graphql_routes_handlers"
)

// GET request is not allow, since POST request is more secure

func main() {
    router := mux.NewRouter()

    http.Handle("/", router)
	serverFile()

    // Configure REST API routes
	router.HandleFunc("/auth/user", Rest.HandleRESTRequest).Methods("POST")
	router.HandleFunc("/auth/token", Rest.HandleRESTRequest).Methods("POST")
    router.HandleFunc("/resource/getinfo",  Auth.Authenticate(Rest.HandleRESTRequest)).Methods("POST")
	router.HandleFunc("/upload/images", Auth.Authenticate(Rest.HandleRESTRequest)).Methods("POST")
    router.HandleFunc("/graphql/resource/info",  Auth.Authenticate(GraphQL.HandleGraphQLRequest)).Methods("POST")

    fmt.Println("Server started at :8080")
    if err :=  http.ListenAndServe(":8080", nil); err != nil {

	}
}

func serverFile() {
	fs := http.FileServer(http.Dir(Config.FILES_DIR + "/test/"))
	fs2 := http.FileServer(http.Dir(Config.FILES_DIR + "/test22/"))

	http.Handle("/icons/", http.StripPrefix("/icons/", fs))
	http.Handle("/images/", http.StripPrefix("/images/", fs2))
}
