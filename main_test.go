package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
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
	mux.HandleFunc("/login", handleLogin)
	mux.HandleFunc("/vehicle", handleVehicle)
}

func tearDown() {

}
