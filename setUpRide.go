package main

import (
	"net/http"
	"github.com/SamirMarin/rnspool/data"
	"encoding/json"
	"github.com/SamirMarin/rnspool/controllerLogic"
)

func handleSetUpRide(w http.ResponseWriter, r *http.Request){
	var err error
	switch r.Method {
	case "POST":
		err = handleSetUpRidePost(w, r)
	default:
		err = handleSetUpRideDefault(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSetUpRidePost(w http.ResponseWriter, r *http.Request) (err error) {
	if r.URL.Path != "/setupride/offered" && r.URL.Path != "/setupride/needed" {
		http.NotFound(w, r)
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var ride data.Ride
	json.Unmarshal(body, &ride)
	var session data.Session
	session = data.Session{Uuid: ride.Uuid, UserId: ride.UserId}
	var isValidSees bool
	isValidSees, err = session.Check()
	if err != nil {
		return
	}
	if isValidSees {
		var location data.Location
		location = data.Location{City: ride.City, Province: ride.Province, Country: ride.Province}
		err = location.Create()
		err = ride.Create(location.Id)
		if err != nil {
			return
		}
		if r.URL.Path == "/setupride/offered" {
			var vehicle data.Vehicle
			vehicle, err =data.VehicleByUserId(ride.UserId, ride.VehicleMake, ride.VehicleModel, ride.VehicleYear)
			err = ride.CreateRideOffered(vehicle.Id)
			if err != nil {
				return
			}
		} else {
			err = ride.CreateRideNeeded()
			if err != nil {
				return
			}
		}
		//we need to make API call to save routes and legs..
		var routes []data.Route
		routes, err = controllerLogic.ObtainRoutes(ride.StartDescrip, ride.EndDescrip)
		if err != nil {
			return
		}
		for _, route := range routes {
			err = route.Create()
			if err != nil {
				return
			}
			err = data.CreateRideHasRouteByIds(ride.Id, route.Id)
			if err != nil {
				return
			}
			for _, leg := range route.Legs {
				err = leg.Create(route.Id)
				if err != nil {
					return
				}
			}
		}
	}
	return
}

func handleSetUpRideDefault(w http.ResponseWriter, r *http.Request) (err error) {
	http.NotFound(w, r)
	return err
}
