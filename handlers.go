package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"time"

	"github.com/rafaeljusto/cctldcentral/db"
)

func cctldCentral(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("cctld-central").Parse(ccTLDTemplate)
	if err != nil {
		log.Printf("Error parsing the template. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		log.Printf("Error executing the template. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func registeredDomains(w http.ResponseWriter, r *http.Request) {
	query := "SELECT date, SUM(number) FROM registered_domains GROUP BY date ORDER BY date"

	rows, err := db.Connection.Query(query)
	if err != nil {
		log.Printf("Error executing a database query. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	dataset := struct {
		Labels []string `json:"labels"`
		Data   []int    `json:"data"`
	}{
		Labels: make([]string, 0),
		Data:   make([]int, 0),
	}

	for rows.Next() {
		var label time.Time
		var data int

		if err := rows.Scan(&label, &data); err != nil {
			log.Printf("Error reading data from the database. Details: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dataset.Labels = append(dataset.Labels, label.Format("2006-01-02"))
		dataset.Data = append(dataset.Data, data)
	}

	output, err := json.Marshal(dataset)
	if err != nil {
		log.Printf("Error encoding the response into JSON. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}
