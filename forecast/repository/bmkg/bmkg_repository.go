package bmkg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adibaulia/bmkg-weather/domain"
	"github.com/adibaulia/bmkg-weather/domain/location"
)

func NewBMKGClient() *BMKGForecast {
	return new(BMKGForecast)
}

func (b *BMKGForecast) GetByProvince(province location.Province) (*domain.Forecast, error) {
	url := fmt.Sprintf("https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-%v.xml", province)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, domain.ErrBadParamInput
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return nil, domain.ErrInternalServerError
	}

	byte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
	}

	xml.Unmarshal(byte, b)

	return nil, nil
}
