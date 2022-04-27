package main

import (
	"assigment3/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.GetStatus)

	http.ListenAndServe(":8080", nil)
}
