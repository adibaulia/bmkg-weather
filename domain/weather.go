package domain

import (
	"time"
)

type (
	WeatherForecast struct {
		Location                 Location                   `json:"location"`
		LastUpdated              time.Time                  `json:"last_updated"`
		WeatherForecastTimeRange []WeatherForecastTimeRange `json:"weather_forecast_time_range"`
	}

	WeatherForecastTimeRange struct {
		Weather        Weather       `json:"weather"`
		Date           DateFormatter `json:"date"`
		TimeStartRange TimeFormatter `json:"time_start_range"`
		TimeEndRange   TimeFormatter `json:"time_end_range"`
	}

	Weather struct {
		NameID string `json:"name_id"`
		NameEN string `json:"name_en"`
		Code   int    `json:"code"`
	}

	WeatherCode int

	WeatherUseCase interface {
		GetWeather(Location) (Weather, error)
	}
)

const (
	ClearSkies         WeatherCode = 0
	PartlyCloudy       WeatherCode = 1
	PartlyCloudy2      WeatherCode = 2
	MostlyCloudy       WeatherCode = 3
	Overcast           WeatherCode = 4
	Haze               WeatherCode = 5
	Smoke              WeatherCode = 10
	Fog                WeatherCode = 45
	LightRain          WeatherCode = 60
	Rain               WeatherCode = 61
	HeavyRain          WeatherCode = 63
	IsolatedShower     WeatherCode = 80
	Thunderstorm       WeatherCode = 95
	SevereThunderstorm WeatherCode = 97
)

var (
	weatherCodeNameMap = map[WeatherCode][2]string{
		ClearSkies:         {"Cerah", "Clear Skies"},
		PartlyCloudy:       {"Cerah Berawan", "Partly Cloudy"},
		PartlyCloudy2:      {"Cerah Berawan", "Partly Cloudy"},
		MostlyCloudy:       {"Berawan", "Mostly Cloudy"},
		Overcast:           {"Berawan Tebal", "Overcast"},
		Haze:               {"Udara Kabur", "Haze"},
		Smoke:              {"Asap", "Smoke"},
		Fog:                {"Kabut", "Fog"},
		LightRain:          {"Hujan Ringan", "Light Rain"},
		Rain:               {"Hujan Sedang", "Rain"},
		HeavyRain:          {"Hujan Lebat", "Heavy Rain"},
		IsolatedShower:     {"Hujan Lokal", "Isolated Shower"},
		Thunderstorm:       {"Hujan Petir", "Severe Thunderstorm"},
		SevereThunderstorm: {"Hujan Petir", "Severe Thunderstorm"},
	}
)

func getWeatherName(wheaterCode WeatherCode, lang string) string {
	langIndex := 0
	if lang == "en" {
		langIndex = 1
	}
	return weatherCodeNameMap[wheaterCode][langIndex]
}
