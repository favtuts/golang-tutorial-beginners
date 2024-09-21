package main

import "net/http"

func Signin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
