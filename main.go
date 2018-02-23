package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	const (
		defaultPort      = "9090"
		defaultPortUsage = "default server port, '9090'"
	)

	// flags
	port := flag.String("port", defaultPort, defaultPortUsage)

	flag.Parse()

	fmt.Println("Server is listening on :", *port)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	airtableURL := os.Getenv("AIRTABLE_URL")
	tablePath := os.Getenv("AIRTABLE_PATH")

	rp := NewReverseProxy(airtableURL, tablePath)
	rp.Transport = &Transport{}

	http.HandleFunc("/", Handle(rp))

	log.Fatal(http.ListenAndServe(":"+*port, nil))

}
