package bmkg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adibaulia/bmkg-weather/domain/errormessage"
	"github.com/adibaulia/bmkg-weather/domain/location"
)

func NewBMKGClient() *BMKGForecast {
	return new(BMKGForecast)
}

func (b *BMKGForecast) GetByProvince(province location.Province) (*BMKGForecast, error) {
	url := fmt.Sprintf("https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-%v.xml", province)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error when getting BMKG API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errormessage.ErrBadParamInput
	}
	if resp.StatusCode >= http.StatusInternalServerError {
		return nil, errormessage.ErrInternalServerError
	}

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error reading response body: %v", err)
		return nil, err
	}

	err = xml.Unmarshal(respByte, b)
	if err != nil {
		log.Printf("error unmarshalling response xml body: %v", err)
		return nil, err
	}
	return b, nil
}
