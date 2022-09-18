package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main12() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./swagger/"))
	router.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", fs))

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:" + goDotEnvVariable("AUCTIONEER_PORT"),
	}

	fmt.Println("Starting Auctioneer Service At Port " + goDotEnvVariable("AUCTIONEER_PORT"))
	log.Fatal(srv.ListenAndServe())
}
