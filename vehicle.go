package main

import (
	"net/http"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"encoding/json"
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
	json.Unmarshal(body, vehicle)
	session := data.Session{Uuid: vehicle.Uuid}
	var isValidSess bool
	isValidSess,err = session.Check()
	if err != nil {
		return
	}
	if isValidSess{
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
