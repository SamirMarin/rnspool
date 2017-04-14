package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SamirMarin/rnspool/backend_webservice/data"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)

}

func setUp() {
	mux = http.NewServeMux()
	mux.HandleFunc("/signup/", handleSignUp)
	writer = httptest.NewRecorder()
}

func tearDown() {

}
func makeJsonReader(path string) (reader *bytes.Reader, err error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	reader = bytes.NewReader(jsonData)
	return
}

func TestHandleSignUpPostDriver(t *testing.T) {
	reader, err := makeJsonReader("testJsonData/drivers/driver1.json")
	if err != nil {
		t.Fatal("error making json reader")

	}
	request, err := http.NewRequest("POST", "/signup/driver", reader)
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
func TestHandleSignUpPostRider(t *testing.T) {
	writer = httptest.NewRecorder()
	jsonReader, err := makeJsonReader("testJsonData/riders/rider1.json")
	if err != nil {
		t.Fatal("error making the json reader")
	}
	request, err := http.NewRequest("POST", "/signup/driver", jsonReader)
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