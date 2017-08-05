package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"github.com/SamirMarin/rnspool/data"
	"fmt"
)

func TestHandleDriverSetUpRidePost(t *testing.T) {
	writer = httptest.NewRecorder()
	reader, err := makeJsonReader("testJsonData/drivers/driverSetUpRide.json")

	if err != nil {
		t.Fatal("error making json reader")
	}
	request , err := http.NewRequest("POST", "/setupride/offered", reader)
	if err != nil {
		t.Fatal("Error forming the request")
	}
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Resoponse code is %v", writer.Code)
	}
	var ride data.Ride
	json.Unmarshal(writer.Body.Bytes(), &ride)
	if ride.Id == 0 {
		t.Errorf("Cannot retrieve JSON ride")
	}
	fmt.Println(ride)
}

func TestHandleRiderSetUpRidePost(t *testing.T){
	writer = httptest.NewRecorder()
	reader, err := makeJsonReader("testJsonData/riders/riderSetUpRide.json")
	if err != nil {
		t.Fatal("error making json reader")
	}
	request, err := http.NewRequest("POST", "/setupride/needed", reader)
	 if err != nil {
		 t.Fatal("error forming request")
	 }
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var ride data.Ride
	json.Unmarshal(writer.Body.Bytes(), &ride)
	if ride.Id == 0 {
		t.Errorf("Cannot retrieve JSON ride")
	}
	//fmt.Println(ride)
}
