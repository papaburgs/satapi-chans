package main

import (
	"log"
	"strconv"
	"time"
)

type forman struct {
	id   int
	host string
}

func main() {
	log.Println("starting")
	start := time.Now()
	// taskChan := make(chan string)
	servergate := make(chan bool, 2)

	defer func() {
		log.Println("done! Duration: ", time.Since(start))
	}()

	formen := []forman{{1, "sls0599"}, {2, "sls0600"}}

	servergate <- true
	go getServers(formen[0], servergate)
	servergate <- true
	go getServers(formen[1], servergate)

	// these two will block until both servers pull off server gate, making sure we wait for both servers to complete
	servergate <- true
	servergate <- true

}

func getServers(host forman, sg chan bool) {
	gatekeeper := make(chan bool, 5)
	log.Println("size of gatekeeper? ", cap(gatekeeper))
	defer func() {
		log.Println(" releasing server lock")
		<-sg
	}()
	for i, j := range getThinList(host, gatekeeper) {
		log.Println("Processing record ", i, " from ", host.host)
		go getErrata(host, j, gatekeeper)
		go getHostDeets(host, j, gatekeeper)
		go getHostThirdOne(host, j, gatekeeper)
	}
	// when I can fill the channel, I know I'm done.

	for c := 0; c < cap(gatekeeper); c++ {
		gatekeeper <- true
	}
}

// getThinList reprevents calling a satellite server and getting
//  a list of Ids back, instead of a slice of strings, could get structs
//  with hostname and server id
func getThinList(host forman, gatekeeper chan bool) (list []string) {
	delay := time.Duration(300) * time.Millisecond
	log.Println("get thin list from ", host.host)
	// random loop to return slice of strings
	var limit int
	if host.id == 1 {
		limit = 120
	} else {
		limit = 412
	}

	for x := 0; x < limit; x++ {
		list = append(list, strconv.Itoa(host.id)+"-"+strconv.Itoa(x))
	}
	time.Sleep(delay)
	return list
}

// parseAndLoad represents the time it gets get a host, parse the result and send that to reddit
func parseAndLoad() {
	delay := time.Duration(300) * time.Millisecond
	time.Sleep(delay)
}

// each of thse getsXXXXX is a call to the server
func getErrata(host forman, j string, gatekeeper chan bool) {
	gatekeeper <- true
	defer func() {
		<-gatekeeper
		log.Println(" releasing gatekeeper")
	}()
	log.Println("getting errata for ", host.host, "-", j)
	parseAndLoad()
}

func getHostDeets(host forman, j string, gatekeeper chan bool) {
	gatekeeper <- true
	defer func() {
		<-gatekeeper
	}()
	log.Println("getting deets for ", host.host, "-", j)
	parseAndLoad()
}
func getHostThirdOne(host forman, j string, gatekeeper chan bool) {
	gatekeeper <- true
	defer func() {
		<-gatekeeper
	}()
	log.Println("getting third for ", host.host, "-", j)
	parseAndLoad()
}
