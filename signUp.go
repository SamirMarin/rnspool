package main

import (
	"encoding/json"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"net/http"
	"fmt"
)

func handleSignUp(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {

	case "GET":
		err = handleSignUpGet(w, r)
	case "POST":
		err = handleSignUpPost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handleSignUpGet(w http.ResponseWriter, r *http.Request) (err error) {
	return err

}
func handleSignUpPost(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/signup/driver" && r.URL.Path != "/signup/rider" {
		http.NotFound(w, r)
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var user data.User
	json.Unmarshal(body, &user)
	fmt.Println(user)
	err = user.Create()
	if err != nil {
		return
	}
	if r.URL.Path == "/signUp/driver" {
		driver := data.Driver{UserId: user.Id}
		err = driver.Create()
		if err != nil {
			return
		}

	} else {
		rider := data.Rider{UserId: user.Id}
		err = rider.Create()
		if err != nil {
			return
		}
	}
	w.WriteHeader(200)
	return
}
