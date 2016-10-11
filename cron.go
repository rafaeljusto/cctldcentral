package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/rafaeljusto/cctldcentral/config"
	"github.com/rafaeljusto/cctldcentral/db"
	"github.com/rafaeljusto/cctldstats/protocol"
)

func startScheduler() chan bool {
	scheduler := gocron.NewScheduler()
	scheduler.Every(1).Day().At("00:00").Do(runSchedulerTask, nil)
	return scheduler.Start()
}

func runSchedulerTask() {
	rows, err := db.Connection.Query(`SELECT cctld, server FROM cctld_server`)
	if err != nil {
		log.Printf("Error executing a database query. Details: %s", err)
		return
	}

	type cctldServer struct {
		cctld  string
		server net.IP
	}

	var cctldServers []cctldServer

	for rows.Next() {
		var data cctldServer
		if err := rows.Scan(&data.cctld, &data.server); err != nil {
			log.Printf("Error reading data from the database. Details: %s", err)
			return
		}

		cctldServers = append(cctldServers, data)
	}

	for _, data := range cctldServers {
		checkCCTLD(data.cctld, data.server)
	}
}

func checkCCTLD(cctld string, server net.IP) {
	client := http.Client{
		Timeout: config.CCTLDCentral.Scheduler.Timeout,
	}

	var err error
	var response *http.Response
	url := fmt.Sprintf("http://%s/domains/registered", server)

	for i := 0; i < config.CCTLDCentral.Scheduler.Retries; i++ {
		if response, err = client.Get(url); err == nil {
			break
		}
		log.Printf("Error retrieving information from server %s of ccTLD %s (attempt %d/%d). Details: %s", server, cctld, i+1, config.CCTLDCentral.Scheduler.Retries, err)
	}

	if response == nil {
		// all attempts caused an error, so we are moving to the next ccTLD

		// TODO(rafaeljusto): insert the same value from the last successful
		// response? This is important to don't mess up the graph
		return
	}

	defer func() {
		response.Body.Close()
	}()

	decoder := json.NewDecoder(response.Body)

	var registeredDomainsResponse protocol.RegisteredDomainsResponse
	if err := decoder.Decode(&registeredDomainsResponse); err != nil {
		log.Printf("Invalid data format from server %s of ccTLD %s. Details: %s", server, cctld, err)
		return
	}

	query := `INSERT INTO registered_domains (cctld, date, number) VALUES ($1, CAST(NOW() at time zone "utc" AS timestamp, $2))`
	if _, err := db.Connection.Exec(query, cctld, registeredDomainsResponse.Number); err != nil {
		log.Printf("Error saving data from server %s of ccTLD %s. Details: %s", server, cctld, err)
		return
	}
}
