package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rafaeljusto/cctldcentral/config"
	"github.com/rafaeljusto/cctldcentral/db"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Error initializing configuration. Details: %s", err)
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Error initializing the database connection. Details: %s", err)
	}

	http.Handle("/", http.HandlerFunc(cctldCentral))
	http.Handle("/domains/registered", http.HandlerFunc(registeredDomains))
	log.Fatal(http.ListenAndServe(":80", nil))
}
