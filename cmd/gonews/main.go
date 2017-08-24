package main

import (
	"net/http"

	"github.com/saspallow/gonews/pkg/app"
)

const port = ":8900"

func main() {
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}
