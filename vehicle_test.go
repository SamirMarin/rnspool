package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"encoding/json"
	"fmt"
)

func TestHandleVehicleLoginPost(t *testing.T) {
	writer = httptest.NewRecorder()
	reader, err := makeJsonReader("testJsonData/vehicles/vehicle1.json")

	if err != nil {
		t.Fatal("error making json reader")
	}
	request , err := http.NewRequest("POST", "/vehicle", reader)
	if err != nil {
		t.Fatal("Error forming the request")
	}
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Resoponse code is %v", writer.Code)
	}
	var vehicle data.Vehicle
	json.Unmarshal(writer.Body.Bytes(), &vehicle)
	if vehicle.Id == 0 {
		t.Errorf("Cannot retrieve JSON vehicle")
	}
	fmt.Println(vehicle)
}
