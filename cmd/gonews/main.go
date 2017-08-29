package main

import (
	"log"
	"net/http"

	"github.com/saspallow/gonews/pkg/app"
	"github.com/saspallow/gonews/pkg/model"
)

const (
	port     = ":8900"
	mongoURL = "mongodb://127.0.0.1:27017"
)

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	err := model.Init(mongoURL)
	if err != nil {
		log.Fatalf("Can not init model: %v", err)
	}
	http.ListenAndServe(port, mux)
}
