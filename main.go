package main

import (
	"net/http"
	"fmt"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", handleLandingPage)
	http.HandleFunc("/signup/", handleSignUp)
	http.HandleFunc("/login", handleLogin)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
