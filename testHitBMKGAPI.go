package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adibaulia/bmkg-weather/forecast/repository/bmkg"
)

func main() {
	resp, err := http.Get("https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-DKIJakarta.xml")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
	}

	var bmkgForecast bmkg.BMKGForecast

	xml.Unmarshal(byte, &bmkgForecast)

	fmt.Printf("forecast: %+v\n", bmkgForecast)

}
