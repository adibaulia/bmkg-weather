package domain

import (
	"github.com/adibaulia/bmkg-weather/domain/location"
	"github.com/adibaulia/bmkg-weather/forecast/repository/bmkg"
)

type ForecastRepository interface {
	GetByProvince(location.Province) (*bmkg.BMKGForecast, error)
}

type (
	Forecasts struct {
		Source           string   `json:"source"`
		Productioncenter string   `json:"production_center"`
		Forecast         Forecast `json:"forecast"`
	}

	Forecast struct {
		Domain   string     `json:"domain"`
		Issue    Issue      `json:"issue"`
		Location []Location `json:"location"`
	}

	Issue struct {
		Timestamp string `json:"timestamp"`
		Year      string `json:"year"`
		Month     string `json:"month"`
		Day       string `json:"day"`
		Hour      string `json:"hour"`
		Minute    string `json:"minute"`
		Second    string `json:"second"`
	}

	Location struct {
		ID          string            `json:"id"`
		Latitude    string            `json:"latitude"`
		Longitude   string            `json:"longitude"`
		Coordinate  string            `json:"coordinate"`
		Type        string            `json:"type"`
		Region      string            `json:"region"`
		Level       string            `json:"level"`
		Description string            `json:"description"`
		Domain      string            `json:"domain"`
		Tags        string            `json:"tags"`
		NameLang    []NameLang        `json:"name_lang"`
		Province    location.Province `json:"province"`
	}

	NameLang struct {
		Name string `json:"name"`
		Lang string `json:"lang"`
	}

	Parameter struct {
		ID          string `json:"id"`
		Description string `json:"description"`
		Type        string `json:"type"`
		Timerange   []struct {
			Type     string `json:"type"`
			H        string `json:"h"`
			Datetime string `json:"datetime"`
			Day      string `json:"day"`
			Value    []struct {
				Value string `json:"value"`
				Unit  string `json:"unit"`
			} `json:"value"`
		} `json:"timerange"`
	}
)
