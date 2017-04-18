package main


import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"fmt"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
)

func TestHandleLoginPostDriver(t *testing.T) {
	writer = httptest.NewRecorder()
	reader, err := makeJsonReader("testJsonData/drivers/driver1.json")
	if err != nil {
		t.Fatal("error making json reader")

	}
	request, err := http.NewRequest("POST", "/login", reader)
	if err != nil {
		t.Fatal("Error forming the request")
	}
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var session data.Session
	json.Unmarshal(writer.Body.Bytes(), &session)
	if session.Id == 0 {
		t.Errorf("Cannot retrive JSON user")
	}
	fmt.Println(session)

}
func TestHandleLoginPostRider(t *testing.T) {
	writer = httptest.NewRecorder()
	reader, err := makeJsonReader("testJsonData/riders/rider1.json")
	if err != nil {
		t.Fatal("error making json reader")

	}
	request, err := http.NewRequest("POST", "/login", reader)
	if err != nil {
		t.Fatal("Error forming the request")
	}
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var session data.Session
	json.Unmarshal(writer.Body.Bytes(), &session)
	if session.Id == 0 {
		t.Errorf("Cannot retrive JSON user")
	}
	fmt.Println(session)

}
