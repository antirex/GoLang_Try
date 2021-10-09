package main

import (
	"Rest_api/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", handler.CreateUserEndpoint)
	http.HandleFunc("/users/", handler.GetUserByIDEndpoint)
	http.HandleFunc("/posts", handler.CreatePostEndpoint)
	http.HandleFunc("/posts/", handler.GetPostByIDEndpoint)
	http.HandleFunc("/posts/users/", handler.GetUsersPostByIdEndpoint)
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
