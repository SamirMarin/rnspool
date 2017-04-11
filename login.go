package main

import (
	"encoding/json"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"net/http"
)

func handleLogin(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {

	case "GET":
		err = handleLoginGet(w, r)
	case "POST":
		err = handleLoginPost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

//TODO: TO IMPLEMENT
func handleLoginGet(w http.ResponseWriter, r *http.Request) (err error) {
	return

}
func handleLoginPost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var userAuth data.User
	json.Unmarshal(body, userAuth)
	var user data.User
	user, err = data.UserByEmail(userAuth.Email)
	if err != nil {
		return
	}
	// check if valid pass
	if user.Password == data.Encrypt(userAuth.Password) {
		var session data.Session
		session, err = user.CreateSession()
		if err != nil {
			return
		}
		var marshalJsonByte []byte
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
