package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/rafaeljusto/cctldcentral/config"
	"github.com/rafaeljusto/cctldcentral/db"
	"github.com/robfig/cron"
)

var port int

func init() {
	flag.IntVar(&port, "port", 80, "server listen port")
}

func main() {
	flag.Parse()
	log.Println("Starting cctldcentral")

	if err := config.Load(); err != nil {
		log.Fatalf("Error initializing configuration. Details: %s", err)
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("Error initializing the database connection. Details: %s", err)
	}

	scheduler := cron.New()
	scheduler.AddFunc("0 0 1 * * *", runSchedulerTask)
	scheduler.Start()

	defer func() {
		scheduler.Stop()
	}()

	http.Handle("/", http.HandlerFunc(cctldCentral))
	http.Handle("/domains/registered", http.HandlerFunc(registeredDomains))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
