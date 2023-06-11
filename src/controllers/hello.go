package controllers

import (
	"net/http"
	"fmt"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Hello Golang")
}
