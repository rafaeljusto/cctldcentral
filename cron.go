package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"

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

	checkWebsite("br", "https://registro.br/estatisticas.html", regexp.MustCompile(`<div class="info total strong">(([0-9]+\.)*[0-9]+)`))
	checkWebsite("gt", "https://www.gt/", regexp.MustCompile(`<div class="span count-domains">[\s\n]*<span>([0-9]+)`))
	checkWebsite("uy", "https://www.nic.uy/Registrar/estadist/index.htm", regexp.MustCompile(`<td align="right" id="Tot-UY">([0-9]+)`))
	checkWebsite("pe", "https://punto.pe", regexp.MustCompile(`<div class="total">[\s\n]*<p>[\s\n]*<span class="num">([0-9]+)`))
	checkWebsite("mx", "https://www.registry.mx/jsf/domain_statistics/instant/iinfo_nb.jsf", regexp.MustCompile(`<td class="InteriorCursosCalendarioListaFila"><strong>TOTAL</strong></td>[\s\n]*<td class="InteriorCursosCalendarioListaFila" style="text-align: right"><strong>(([0-9]+\,)*[0-9]+)`))
	checkWebsite("hn", "http://nic.hn/", regexp.MustCompile(`<strong>Total de Dominios Registrados</strong></p>[\s\n]*<div align="center"><span style="color: red; font-family: arial; font-size: 300%;"><strong><em>(([0-9]+\,)*[0-9]+)`))
	checkWebsite("ar", "https://nic.ar/nic-argentina/en-cifras", regexp.MustCompile(`(([0-9]+\.)+[0-9]+)</td>[\n\s]*</tr>`))
	checkWebsite("do", "https://www.nic.do/en/domain-names-registered-under-do-instantly/", regexp.MustCompile(`<th align="left">Total</th><th align="left">([0-9]+)`))
	checkWebsite("cl", "https://www.nic.cl/estadisticas/numDominios.json", regexp.MustCompile(`([0-9]+),"f":null}\]}[\n\s]*]`))
}

func checkCCTLD(cctld string, server net.IP) {
	log.Printf("Checking server %s of ccTLD %s", server, cctld)

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

	query := `INSERT INTO registered_domains (cctld, date, number) VALUES ($1, CURRENT_TIMESTAMP AT TIME ZONE 'UTC', $2)`
	if _, err := db.Connection.Exec(query, cctld, registeredDomainsResponse.Number); err != nil {
		log.Printf("Error saving data from server %s of ccTLD %s. Details: %s", server, cctld, err)
		return
	}
}

func checkWebsite(cctld, url string, totalStatisticRX *regexp.Regexp) {
	log.Printf("Checking website %s of ccTLD %s", url, cctld)

	client := http.Client{
		Timeout: config.CCTLDCentral.Scheduler.Timeout,
	}

	var err error
	var response *http.Response

	for i := 0; i < config.CCTLDCentral.Scheduler.Retries; i++ {
		if response, err = client.Get(url); err == nil {
			break
		}
		log.Printf("Error retrieving information from website %s of ccTLD %s (attempt %d/%d). Details: %s", url, cctld, i+1, config.CCTLDCentral.Scheduler.Retries, err)
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

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading the response body from website %s of ccTLD %s. Details: %s", url, cctld, err)
		return
	}

	findResults := totalStatisticRX.FindSubmatch(content)
	if len(findResults) <= 1 {
		log.Printf("Total not found in response body from website %s of ccTLD %s", url, cctld)
		return
	}

	result := string(findResults[1])
	result = strings.Replace(result, ".", "", -1)
	result = strings.Replace(result, ",", "", -1)

	total, err := strconv.Atoi(result)
	if err != nil {
		log.Printf("Error converting total '%s' from website %s of ccTLD %s. Details: %s", string(findResults[1]), url, cctld, err)
	}

	query := `INSERT INTO registered_domains (cctld, date, number) VALUES ($1, CURRENT_TIMESTAMP AT TIME ZONE 'UTC', $2)`
	if _, err := db.Connection.Exec(query, cctld, total); err != nil {
		log.Printf("Error saving data from website %s of ccTLD %s. Details: %s", url, cctld, err)
		return
	}
}
