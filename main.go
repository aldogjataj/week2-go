package main

import (
	"acme/api"
	"fmt"
	"io"
	"net/http"
)

func main() {
	//using a multiplexer instead
	router := http.NewServeMux()

	router.HandleFunc("GET /", rootHandler)
	router.HandleFunc("GET /api/users", api.GetUsers)
	router.HandleFunc("POST /api/users", api.CreateUser)
	router.HandleFunc("GET /api/users/{id}", api.GetSingleUser)
	router.HandleFunc("DELETE /api/users/{id}", api.DeleteUser)
	router.HandleFunc("PUT /api/users/{id}", api.UpdateUser)

	fmt.Println("server on port 3000")
	err := http.ListenAndServe(":3000", CorsMiddleware(router))

	if err != nil {
		fmt.Println("error starting server", err)
	}
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving route /")
	io.WriteString(w, "Hello Aldo")
}

// func simple_main() {
// 	http.HandleFunc("/", rootHandler)
// 	http.HandleFunc("/api/users", getUsers)

// 	fmt.Println("server on port 3000")
// 	err := http.ListenAndServe(":3000", nil)

// 	if err != nil {
// 		fmt.Println("error starting server", err)
// 	}
// }
