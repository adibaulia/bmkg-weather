package domain

import "github.com/adibaulia/bmkg-weather/domain/location"

type (
	TemperatureForecast struct {
		Location                     location.Location              `json:"location"`
		LastUpdatedApi               LastUpdatedTimeFormatter       `json:"last_updated_api"`
		LastUpdatedAccess            LastUpdatedTimeFormatter       `json:"last_updated_access"`
		TemperatureForecastTimeRange []TemperatureForecastTimeRange `json:"temperature_forecast_time_range"`
	}
	TemperatureForecastTimeRange struct {
		Temperature    Temperature   `json:"temperature"`
		Date           DateFormatter `json:"date"`
		TimeStartRange TimeFormatter `json:"time_start_range"`
		TimeEndRange   TimeFormatter `json:"time_end_range"`
	}

	Temperature struct {
		MinTemperature TemperatureUnit `json:"min_temperature"`
		MaxTemperature TemperatureUnit `json:"max_temperature"`
		Celcius        float64         `json:"celcius,omitempty"`
		Fahrenheit     float64         `json:"fahrenheit,omitempty"`
	}

	TemperatureUnit struct {
		Celcius    float64 `json:"celcius,omitempty"`
		Fahrenheit float64 `json:"fahrenheit,omitempty"`
	}

	TemperatureUseCase interface {
		GetTemperatureByProvince(string) ([]TemperatureForecast, error)
		GetTemperatureByCity(string, string) (TemperatureForecast, error)
	}
)
