package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", handleLandingPage)
	http.HandleFunc("/signup/", handleSignUp)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/vehicle", handleVehicle)
	http.HandleFunc("/setupride/", handleSetUpRide)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
