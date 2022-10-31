package bmkg

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/adibaulia/bmkg-weather/domain/errormessage"
	"github.com/adibaulia/bmkg-weather/domain/location"
)

const (
	BMKGAPI   = "https://data.bmkg.go.id/DataMKG/MEWS/DigitalForecast/DigitalForecast-%v.xml"
	userAgent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36"
)

func NewBMKGClient() *BMKGForecast {
	return new(BMKGForecast)
}

func (b *BMKGForecast) GetByProvince(province location.Province) (*BMKGForecast, error) {
	b = &BMKGForecast{}
	url := fmt.Sprintf(BMKGAPI, province)
	log.Printf("Get BMKG Data from URL: %s, with province: %v", url, province)

	c := http.Client{Timeout: 1 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error when creating client HTTP BMKG API: %v", err)
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Accept", "*/*")

	resp, err := c.Do(req)
	if err != nil {
		log.Printf("error when getting BMKG API: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errormessage.ErrNotFound
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
