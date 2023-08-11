package restrouteshandlers

import (
	"net/http"
	//"time"
	Request "pkg/request_handlers"
	//"github.com/dgrijalva/jwt-go"
)


func HandleRESTRequest(w http.ResponseWriter, r *http.Request) {
    // You can switch based on the HTTP method (GET, POST, PUT, DELETE, etc.)
    switch r.Method {
	case http.MethodPatch:
		handlePATCHRequest(w, r)
    case http.MethodGet:
        handleGETRequest(w, r)
    case http.MethodPost:
		handlePOSTRequest(w, r)
    case http.MethodPut:
        handlePUTRequest(w, r)
    case http.MethodDelete:
        handleDELETERequest(w, r)
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func handlePATCHRequest(w http.ResponseWriter, r *http.Request) {
    // Handle GET request logic
    // You can fetch data from a database or perform some processing
    // Then write the response using w.Write() or other methods
}

func handleGETRequest(w http.ResponseWriter, r *http.Request) {
    // Handle GET request logic
    // You can fetch data from a database or perform some processing
    // Then write the response using w.Write() or other methods
}

func handlePOSTRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/auth/user":
		Request.Auth_Token(w, r)
	case "/auth/token":
		Request.GenerateJWT(w, r)
	case "/resource/getinfo":
		Request.GetCustomerInfo(w, r)
    case "/upload/images":
        Request.UploadImges(w, r)
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func handlePUTRequest(w http.ResponseWriter, r *http.Request) {
    // Handle PUT request logic
    // Parse the request body, update data in the database, and handle errors
    // Write a response indicating success or failure
}

func handleDELETERequest(w http.ResponseWriter, r *http.Request) {
    // Handle DELETE request logic
    // Extract parameters from the request, delete data from the database, handle errors
    // Write a response indicating success or failure
}
