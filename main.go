package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	bu := "https://us.booksy.com/api/us/2/customer_api/me/bookings/time_slots/"
	start := time.Now()
	s := start.Format("2006-01-02")
	end := start.Add(14 * 24 * time.Hour)
	e := end.Format("2006-01-02")

	// Service Variant Id determines which barbor! - This one is Hair Jordan
	// TODO: Update to map of map[barber]svi
	barbers := map[string]string{
		"jordan": "8941492",
		"kegan":  "9080419",
		"pete":   "9098534",
	}

	for _, id := range barbers {
		url := fmt.Sprintf("%s?start_date=%s&end_date=%s&service_variant_id=%s", bu, s, e, id)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("error creating new request")
		}
		req.Header.Add("x-api-key", "web-e3d812bf-d7a2-445d-ab38-55589ae6a121")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error completing request")
		}
		defer resp.Body.Close()

		var data struct {
			TimeSlots interface{} `json:"time_slots"`
			Resources interface{} `json:"resources"`
		}

		err = ReadJSON(resp, &data)
		if err != nil {
			fmt.Println("error decoding")
			fmt.Println(err)
		}

		fmt.Println(data.TimeSlots)
	}

}
