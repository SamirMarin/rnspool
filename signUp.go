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
	case "POST":
		err = handleSignUpPost(w, r)
	default:
		err = handleSignUpDefault(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
	fmt.Println(user.Email)
	fmt.Println(user.Password)

	var location data.Location
	location = data.Location{City: user.City, Province: user.Province, Country: user.Country}
	err = location.Create()
	if err != nil {
		return
	}

	var address data.Address
	address = data.Address{AptNum: user.AptNum, HouseNum: user.HouseNum, Street: user.Street,
		PostalCode: user.PostalCode, LocationId: location.Id}
	err = address.Create()
	if err != nil {
		return
	}
	user.AddressId = address.Id
	err = user.Create()
	if err != nil {
		return
	}
	if r.URL.Path == "/signup/driver" {
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

func handleSignUpDefault(w http.ResponseWriter, r *http.Request) (err error) {
	http.NotFound(w, r)
	return err
}
