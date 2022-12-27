package main

import (
	service "TaskFinal/service"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:8080", service.BuildRoutes())
	if err != nil {
		panic(err)
	}
}
