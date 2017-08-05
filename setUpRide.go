package main

import (
	"net/http"
	"github.com/SamirMarin/rnspool/data"
	"encoding/json"
	"github.com/SamirMarin/rnspool/controllerLogic"
	"fmt"
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
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/setupride/offered" && r.URL.Path != "/setupride/needed" {
		http.NotFound(w, r)
		return
	}
	fmt.Println("not failing the path")
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var ride data.Ride
	json.Unmarshal(body, &ride)
	fmt.Println(ride.Uuid)
	fmt.Println(ride.StartDescrip)
	fmt.Println(ride.EndDescrip)
	fmt.Println(ride.AvailableSeats)
	fmt.Println(ride.TimeLeaving)
	fmt.Println(ride.TimePickUp)
	fmt.Println(ride.UserId)
	var session data.Session
	session = data.Session{Uuid: ride.Uuid, UserId: ride.UserId}
	var isValidSees bool
	isValidSees, err = session.Check()
	fmt.Println(isValidSees)
	if err != nil {
		return
	}
	if isValidSees {
		var location data.Location
		location = data.Location{City: ride.City, Province: ride.Province, Country: ride.Province}
		err = location.Create()
		fmt.Println(location.Id)
		err = ride.Create(location.Id)
		if err != nil {
			fmt.Println("GOING TO FUCKEN ERROR", err)
			return
		}
		if r.URL.Path == "/setupride/offered" {
			var vehicle data.Vehicle
			vehicle, err =data.VehicleByUserId(ride.UserId, ride.VehicleMake, ride.VehicleModel, ride.VehicleYear)
			if err != nil {
				fmt.Println("GOING TO error at Vehicle by user ID", err)
				return
			}
			err = ride.CreateRideOffered(vehicle.Id)
			if err != nil {
				fmt.Println("GOING TO error at create ride offered", err)
				return
			}
		} else {
			err = ride.CreateRideNeeded()
			if err != nil {
				fmt.Println("GOING TO error at create ride needed", err)
				return
			}
		}
		//we need to make API call to save routes and legs..
		var routes []data.Route
		routes, err = controllerLogic.ObtainRoutes(ride.StartDescrip, ride.EndDescrip)
		if err != nil {
			fmt.Println("GOING TO error at obtainRoutes", err)
			return
		}
		for _, route := range routes {
			fmt.Println("here now get ready")
			fmt.Println(route)
			err = route.Create()
			if err != nil {
				fmt.Println("GOING TO error at create route", err)
				return
			}
			fmt.Println("routeID here it comes")
			fmt.Println(route.Id)
			err = data.CreateRideHasRouteByIds(ride.Id, route.Id)
			if err != nil {
				fmt.Println("GOING TO error at create ride has route by ids", err)
				return
			}
			for _, leg := range route.Legs {
				err = leg.Create(route.Id)
				if err != nil {
					fmt.Println("GOING TO error at create leg", err)
					return
				}
			}
		}
		var marshalJsonByte []byte
		marshalJsonByte, err = json.MarshalIndent(&ride, "", "\t\t")
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

func handleSetUpRideDefault(w http.ResponseWriter, r *http.Request) (err error) {
	http.NotFound(w, r)
	return err
}
