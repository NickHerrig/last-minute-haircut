package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func ReadJSON(r *http.Response, dst interface{}) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(nil, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			return fmt.Errorf("body contains unknown key %s", fieldName)
		case err.Error() == "http: request body too large":
			return fmt.Errorf("body must not be larger than %d bytes", maxBytes)
		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func main() {

	bu := "https://us.booksy.com/api/us/2/customer_api/me/bookings/time_slots/"
	start := time.Now()
	s := start.Format("2006-01-02")
	end := start.Add(40 * 24 * time.Hour)
	e := end.Format("2006-01-02")

	// Service Variant Id determines which barbor! - This one is Hair Jordan
	// TODO: Update to map of map[barber]svi
	svi := "8941492"

	url := fmt.Sprintf("%s?start_date=%s&end_date=%s&service_variant_id=%s", bu, s, e, svi)
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

	fmt.Println(data.Resources)

}
