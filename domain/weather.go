package domain

import (
	"log"
	"strconv"
)

type (
	WeatherForecast struct {
		Location                 Location                   `json:"location"`
		LastUpdatedApi           LastUpdatedTimeFormatter   `json:"last_updated_api"`
		LastUpdatedAccess        LastUpdatedTimeFormatter   `json:"last_updated_access"`
		WeatherForecastTimeRange []WeatherForecastTimeRange `json:"weather_forecast_time_range"`
	}

	WeatherForecastTimeRange struct {
		Weather        Weather       `json:"weather"`
		Date           DateFormatter `json:"date"`
		TimeStartRange TimeFormatter `json:"time_start_range"`
		TimeEndRange   TimeFormatter `json:"time_end_range"`
	}

	Weather struct {
		NameID      string      `json:"name_id"`
		NameEN      string      `json:"name_en"`
		WeatherCode WeatherCode `json:"weather_code"`
	}

	WeatherCode int

	WeatherUseCase interface {
		GetWeatherByProvince(string) ([]WeatherForecast, error)
		GetWeatherByCity(string, string) (WeatherForecast, error)
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

func GetWeatherCode(weatherCodeStr string) WeatherCode {
	weatherCodeInt, err := strconv.Atoi(weatherCodeStr)
	if err != nil {
		log.Printf("error converting weather code int: %v", err)
		return 500
	}
	for weatherCode := range weatherCodeNameMap {
		if weatherCodeInt == int(weatherCode) {
			return weatherCode
		}
	}
	return 500
}

func GetWeatherName(weatherCode WeatherCode, lang string) string {
	langIndex := 0
	if lang == "en" {
		langIndex = 1
	}
	return weatherCodeNameMap[weatherCode][langIndex]
}
