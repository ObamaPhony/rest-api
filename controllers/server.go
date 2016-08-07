package controllers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome.\n")
}

func Server() {
	router := httprouter.New()
	router.GET("/", Index)

	http.ListenAndServe(":8080", router)
}
