package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// API url endpoint for Hair Jordan's appointments
var url string = "https://www.genbook.com/bookings/api/serviceproviders/30230662/services/989056738/resources/989056742"

type body struct {
	Timezone     map[string]interface{} `json:"timezone"`
	Dates        []string               `json:"dates"`
	Bookingdates []string               `json:"bookingdates`
}

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	resp, err := http.Get(url)
	if err != nil {
		logger.Fatal(err)
	}

	var bd body
	err = json.NewDecoder(resp.Body).Decode(&bd)
	if err != nil {
		logger.Fatal(err)
	}

	var lstminapts []time.Time
	for _, apt := range bd.Bookingdates {

		year := apt[:4]
		month := apt[4:6]
		day := apt[6:8]

		dts := fmt.Sprintf("%s-%s-%sT00:00:00.000Z", year, month, day)

		t, err := time.Parse(time.RFC3339, dts)
		if err != nil {
			logger.Fatal(err)
		}

		now := time.Now()
		dif := t.Sub(now)

		if dif < 30*24*time.Hour {
			lstminapts = append(lstminapts, t)
		}

	}

	srv := &http.Server{
		Addr:         ":4000",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	type Data struct {
		Barber       string
		Appointments []time.Time
	}

	d := Data{"Jordan", lstminapts}
	tmpl := template.Must(template.ParseFiles("templates/index.html.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, d)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	err = srv.ListenAndServe()
	logger.Fatal(err)
}
