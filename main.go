package main

import (
	"net/http"
	"log"
	"os"
	"encoding/json"
	"fmt"
	"time"
)

// API url endpoint for Hair Jordan's appointments 
var url string = "https://www.genbook.com/bookings/api/serviceproviders/30230662/services/989056738/resources/989056742"


func main() {
	var data map[string]interface{}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatal(err)
	}

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		logger.Fatal(err)
	}

	apts := data["bookingdates"]
	fmt.Println(apts)

	// Parse times into time.Time go objects? 

	// display this data in beautiful tailwind cards
}
