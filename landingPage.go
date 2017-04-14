package main

import "net/http"

func handleLandingPage(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {

	default:
		err = handleLandingDefault(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handleLandingDefault(w http.ResponseWriter, r *http.Request) (err error) {
	http.NotFound(w, r)
	return err

}
