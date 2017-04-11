package main

import "net/http"

func handleLandingPage(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {

	case "GET":
		err = handleLandingGet(w, r)
	case "POST":
		err = handleLandingPost(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handleLandingGet(w http.ResponseWriter, r *http.Request) (err error) {
	return

}
func handleLandingPost(w http.ResponseWriter, r *http.Request) (err error) {
	return
}
