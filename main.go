package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
//	"sync/atomic"
)


// only care about the bookingdates json key
// silently ignore other data on decode
type response struct {
	Bookingdates []string               `json:"bookingdates`
}

type barber struct {
	Name string
	Img  string
	Apts []time.Time
}

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	genbookapi := "https://www.genbook.com/bookings/api/serviceproviders/"

	barberendpoints := map[string]string{
		"jordan":  fmt.Sprintf("%s/30230662/services/989056738/resources/989056742", genbookapi),
		"pete":    fmt.Sprintf("%s/31191440/services/10282291592/resources/10282190294", genbookapi),
		"brandon": fmt.Sprintf("%s/30377943/services/2394050193/resources/2394025610", genbookapi),
		"luis":    fmt.Sprintf("%s/30250062/services/1173749692/resources/1173749696", genbookapi),
		"zach":    fmt.Sprintf("%s/30302725/services/1547629284/resources/1547629288", genbookapi),
		"paul":    fmt.Sprintf("%s/30309745/services/1603733980/resources/1603733984", genbookapi),
		"kegan":   fmt.Sprintf("%s/30352805/services/2098565278/resources/2098565282", genbookapi),
	}

	var barbers []barber
	for k, v := range  barberendpoints {
		var b barber
		b.Name = k
		b.Img = fmt.Sprintf("/static/img/%s.jpeg", b.Name)

		resp, err := http.Get(v)
		if err != nil {
			logger.Fatal(err)
		}

		var r response
		err = json.NewDecoder(resp.Body).Decode(&r)
		if err != nil {
			logger.Fatal(err)
		}

		for _, apt := range r.Bookingdates {

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

			if dif < 7*24*time.Hour {
				b.Apts = append(b.Apts, t)
			}

		}
		barbers = append(barbers, b)
	}

	srv := &http.Server{
		Addr:         ":4000",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}


	tmpl := template.Must(template.ParseFiles("templates/index.html.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, barbers)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	err := srv.ListenAndServe()
	logger.Fatal(err)
}
