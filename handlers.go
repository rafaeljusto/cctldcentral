package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/rafaeljusto/cctldcentral/db"
)

func cctldCentral(w http.ResponseWriter, r *http.Request) {
	query := "SELECT DISTINCT cctld FROM registered_domains ORDER BY cctld"

	rows, err := db.Connection.Query(query)
	if err != nil {
		log.Printf("Error executing a database query. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var cctlds []string
	for rows.Next() {
		var cctld string
		if err := rows.Scan(&cctld); err != nil {
			log.Printf("Error reading data from the database. Details: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		cctlds = append(cctlds, cctld)
	}

	t, err := template.New("cctld-central").Parse(ccTLDTemplate)
	if err != nil {
		log.Printf("Error parsing the template. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data := struct {
		CCTLDs []string
	}{
		CCTLDs: cctlds,
	}

	if err := t.Execute(w, data); err != nil {
		log.Printf("Error executing the template. Details: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func registeredDomains(w http.ResponseWriter, r *http.Request) {
	var query string
	var rows *sql.Rows
	var err error

	cctld := r.URL.Query().Get("cctld")

	var period time.Duration
	if periodStr := r.URL.Query().Get("period"); periodStr != "" {
		if period, err = time.ParseDuration(periodStr); err != nil {
			log.Printf("Invalid duration '%s'. Details: %s", periodStr, err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if cctld != "" && period > 0 {
		from := time.Now().UTC().Add(-period)
		query = "SELECT CAST(date AS DATE), SUM(number) FROM registered_domains WHERE cctld = $1 AND date > $2 GROUP BY CAST(date AS DATE) ORDER BY CAST(date AS DATE)"
		rows, err = db.Connection.Query(query, cctld, from)

	} else if cctld != "" {
		query = "SELECT CAST(date AS DATE), SUM(number) FROM registered_domains WHERE cctld = $1 GROUP BY CAST(date AS DATE) ORDER BY CAST(date AS DATE)"
		rows, err = db.Connection.Query(query, cctld)

	} else if period > 0 {
		from := time.Now().UTC().Add(-period)
		query = "SELECT CAST(date AS DATE), SUM(number) FROM registered_domains WHERE date > $1 GROUP BY CAST(date AS DATE) ORDER BY CAST(date AS DATE)"
		rows, err = db.Connection.Query(query, from)

	} else {
		query = "SELECT CAST(date AS DATE), SUM(number) FROM registered_domains GROUP BY CAST(date AS DATE) ORDER BY CAST(date AS DATE)"
		rows, err = db.Connection.Query(query)
	}

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
