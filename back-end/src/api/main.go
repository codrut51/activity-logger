package main

import (
	"fmt"
	"net/http"
	"api/components/models/config"
)

var PORT string

func main() {
	port := initServer()
	http.HandleFunc("/", helloServer)
	http.HandleFunc("/api/login", login)
	http.ListenAndServe(":" + port, nil)
	
}

func initServer() string {
	cfg := config.GetConfig()
	return cfg.Server.Port
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf("%+v", config.getConfig())
	cfg := config.GetConfig()
	// fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	fmt.Fprintf(w, "Your config is: %+v! \n", cfg)
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}