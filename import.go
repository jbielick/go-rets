// Total Listings: 1751590
// go run types.go main.go  5992.19s user 517.07s system 226% cpu 47:50.42 total

package main

import (
	"compress/gzip"
	"encoding/xml"
	"flag"
	"fmt"
	r "gopkg.in/dancannon/gorethink.v2"
	"log"
	"net/http"
	"time"
)

func schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func main() {
	r.SetTags("gorethink", "json")

	pickupUrl := flag.String("url", "", "the pickup url for the gzipped xml file")
	user := flag.String("user", "", "the basic auth user")
	pass := flag.String("pass", "", "the basic auth password")

	flag.Parse()

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "realestate",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", *pickupUrl, nil)

	req.SetBasicAuth(*user, *pass)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	xmlStream, gerr := gzip.NewReader(resp.Body)

	if gerr != nil {
		log.Fatalln(gerr.Error())
		return
	}

	decoder := xml.NewDecoder(xmlStream)

	total := 0
	runner := 0

	logCount := func() { fmt.Printf("per second: %v, total: %v \n", runner, total); runner = 0 }

	stop := schedule(logCount, time.Second)

	concurrency := 100
	sem := make(chan bool, concurrency)

	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			// ...and its name is "Listing"
			if se.Name.Local == "Listing" {
				var l Listing

				err := decoder.DecodeElement(&l, &se)

				if err != nil {
					fmt.Println(err)
				}

				sem <- true
				go func(session *r.Session, listing *Listing) {
					defer func() { <-sem }()
					rerr := r.DB("realestate").Table("listings").Insert(listing).Exec(session)

					if rerr != nil {
						fmt.Println(rerr)
					}
				}(session, &l)

				runner++
				total++
			}
		default:
		}

	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}

	stop <- true
	fmt.Printf("Total Listings: %d \n", total)
}
