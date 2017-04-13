package main

import (
	"encoding/json"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"net/http"
	"fmt"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {

	case "POST":
		err = handleLoginPost(w, r)
	default:
		err = handleLoginDefault(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//TODO: TO IMPLEMENT
func handleLoginPost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var userAuth data.User
	json.Unmarshal(body, &userAuth)
	var user data.User
	user, err = data.UserByEmail(userAuth.Email)
	fmt.Println(userAuth.Email)
	if err != nil {
		fmt.Println("error when obtaining user")
		return
	}
	// check if valid pass
	fmt.Println("obtained pass", user.Password)
	fmt.Println("received pass", userAuth.Password)
	if user.Password == data.Encrypt(userAuth.Password) {
		var session data.Session
		session, err = user.CreateSession()
		if err != nil {
			fmt.Print("internal error in createSEssion")
			return
		}
		var marshalJsonByte []byte
		if err != nil {
			return
		}
		marshalJsonByte, err = json.MarshalIndent(&session, "", "\t\t")
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalJsonByte)
		return
	}
	//the server cannot or will not process the request due to something that is perceived to be a client error
	w.WriteHeader(400)
	return

}
func handleLoginDefault(w http.ResponseWriter, r *http.Request) (err error) {
	http.NotFound(w, r)
	return

}
