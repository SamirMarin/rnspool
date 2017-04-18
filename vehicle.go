package main

import (
	"encoding/json"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"net/http"
	"fmt"
)

func handleVehicle(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "POST":
		err = handleVehiclePost(w, r)
	default:
		err = handleVehicleDefault(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handleVehiclePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var vehicle data.Vehicle
	json.Unmarshal(body, &vehicle)
	fmt.Println("what", vehicle.DriverId, vehicle.Uuid)
	var session data.Session
	session = data.Session{Uuid: vehicle.Uuid, UserId: vehicle.DriverId}
	var isValidSess bool
	fmt.Println("its a ssess", session)
	isValidSess, err = session.Check()
	fmt.Println("there boool come one now:", isValidSess)
	if err != nil {
		return
	}
	if isValidSess {
		err = vehicle.Create()
		if err != nil {
			return
		}
		var marshalJsonByte []byte
		marshalJsonByte, err = json.MarshalIndent(&vehicle, "", "\t\t")
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshalJsonByte)
		return
	}
	w.WriteHeader(400)
	return
}
func handleVehicleDefault(w http.ResponseWriter, r *http.Request) (err error) {
	http.NotFound(w, r)
	return
}
