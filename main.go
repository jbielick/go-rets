package main

import (
	"compress/gzip"
	"encoding/xml"
	"flag"
	"fmt"
	r "gopkg.in/dancannon/gorethink.v2"
	"log"
	"os"
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

var inputFile = flag.String("infile", "adwerx.xml.gz", "Input file path")

type customTime struct {
	time.Time
}

func (c *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parsed, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return err
	}
	*c = customTime{parsed}
	return nil
}

func main() {
	flag.Parse()

	r.SetTags("gorethink", "json")

	var session *r.Session

	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "realestate",
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	archive, err := os.Open(*inputFile)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer archive.Close()

	xmlFile, gerr := gzip.NewReader(archive)

	if gerr != nil {
		log.Fatalln(gerr.Error())
	}

	decoder := xml.NewDecoder(xmlFile)

	total := 0
	runner := 0

	logCount := func() { fmt.Println("per second: ", runner); runner = 0 }

	stop := schedule(logCount, time.Second)

	concurrency := 10
	sem := make(chan bool, concurrency)

	var inElement string

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
			inElement = se.Name.Local
			// ...and its name is "Listing"
			if inElement == "Listing" {
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
