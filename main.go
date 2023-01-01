package main

import (
	"fmt"
	"log"
)

// func homeHandler(wr http.ResponseWriter, rq *http.Request) {
// 	fmt.Fprintln(wr, "Therapy Session for Docker compose container build issue")
// }

func main() {
	fmt.Println("Staring a simple web application")
	log.Println("debugging Dcoker compose issue")
	
	// mux := http.NewServeMux()
	
	// mux.HandleFunc("/", homeHandler)

	// srv := http.Server{
	// 	Addr:    ":4040",
	// 	Handler: mux,
	// }
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	log.Println("error starting go-app server")
	// }
}
