package usecase

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/adibaulia/bmkg-weather/domain"
	"github.com/adibaulia/bmkg-weather/domain/errormessage"
	"github.com/adibaulia/bmkg-weather/domain/location"
	"github.com/adibaulia/bmkg-weather/forecast/repository/bmkg"
)

type temperatureUseCase struct {
	forecastRepo domain.ForecastRepository
}

func NewTemperatureUseCase(forecastRepo domain.ForecastRepository) domain.TemperatureUseCase {
	return &temperatureUseCase{forecastRepo: forecastRepo}
}

func (t *temperatureUseCase) GetTemperatureByProvince(provinceStr string) ([]domain.TemperatureForecast, error) {

	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	province := location.Province(provinceStr)
	BMKGForecast, err := t.forecastRepo.GetByProvince(province)
	if err != nil {
		log.Printf("error getting weather forecast: %v", err)
		return nil, err
	}

	lastUpdate, err := time.ParseInLocation("20060102150405", BMKGForecast.Forecast.Issue.Timestamp, time.UTC)
	if err != nil {
		log.Printf("error parsing time lastUpdate from BMKGForecast: %v", err)
		return nil, err
	}

	var result []domain.TemperatureForecast

	for _, l := range BMKGForecast.Forecast.Area {
		temperaturePerLocation := domain.TemperatureForecast{
			Location: location.Location{
				ID:          l.ID,
				Latitude:    l.Latitude,
				Longitude:   l.Longitude,
				Coordinate:  l.Coordinate,
				Type:        l.Type,
				Region:      l.Region,
				Level:       l.Level,
				Province:    province,
				Tags:        l.Tags,
				Domain:      l.Domain,
				Description: l.Description,
				NameLang:    nameLangCompose(l.Name),
			},
			LastUpdatedApi:    domain.LastUpdatedTimeFormatter(lastUpdate.In(loc)),
			LastUpdatedAccess: domain.LastUpdatedTimeFormatter(now.In(loc)),
		}

		var temperatureForecasts []domain.TemperatureForecastTimeRange
		var maxTemperature bmkg.Parameter
		var minTemperature bmkg.Parameter
		var temperature bmkg.Parameter
		for _, t := range l.Parameter {
			t := t
			switch t.ID {
			case "t":
				temperature = t
				continue
			case "tmin":
				minTemperature = t
				continue
			case "tmax":
				maxTemperature = t
				continue
			}
		}

		for _, temperatureTR := range temperature.Timerange {

			datetime, err := time.Parse("200601021504", temperatureTR.Datetime)
			if err != nil {
				return nil, err
			}

			maxTempC, maxTempF := ComposeTemperatureValue(maxTemperature.Timerange, datetime.Format("20060102"))
			minTempC, minTempF := ComposeTemperatureValue(minTemperature.Timerange, datetime.Format("20060102"))
			tempC, tempF := ComposeTemperatureValue(minTemperature.Timerange, "")

			temperatureForecast := domain.TemperatureForecastTimeRange{
				Temperature: domain.Temperature{
					MaxTemperature: domain.TemperatureUnit{
						Celcius:    maxTempC,
						Fahrenheit: maxTempF,
					},
					MinTemperature: domain.TemperatureUnit{
						Celcius:    minTempC,
						Fahrenheit: minTempF,
					},
					Celcius:    tempC,
					Fahrenheit: tempF,
				},
				Date:           domain.DateFormatter(datetime),
				TimeStartRange: domain.TimeFormatter(datetime),
				TimeEndRange:   domain.TimeFormatter(datetime.Add(6 * time.Hour)),
			}

			temperatureForecasts = append(temperatureForecasts, temperatureForecast)
		}

		temperaturePerLocation.TemperatureForecastTimeRange = temperatureForecasts
		result = append(result, temperaturePerLocation)
	}
	return result, nil

}

func (t *temperatureUseCase) GetTemperatureByCity(provinceStr, cityStr string) (domain.TemperatureForecast, error) {

	cities, err := t.GetTemperatureByProvince(provinceStr)
	if err != nil {
		return domain.TemperatureForecast{}, err
	}

	for _, c := range cities {
		if strings.EqualFold(c.Location.Description, cityStr) {
			return c, nil
		}
	}
	return domain.TemperatureForecast{}, errormessage.ErrNotFound
}

func ComposeTemperatureValue(temperatureValue []bmkg.Timerange, datetime string) (float64, float64) {
	var temperatureTimerange bmkg.Timerange
	for _, temperatureTR := range temperatureValue {
		if datetime == "" {
			temperatureTimerange = temperatureTR
		} else if temperatureTR.Day == datetime {
			temperatureTimerange = temperatureTR
		}
	}

	var celcius float64
	var fahrenheit float64

	for _, v := range temperatureTimerange.Value {
		if v.Unit == "C" {
			mc, err := strconv.ParseFloat(v.Text, 32)
			if err != nil {
				panic(err)
			}
			celcius = mc
		}
		if v.Unit == "F" {
			mc, err := strconv.ParseFloat(v.Text, 32)
			if err != nil {
				panic(err)
			}
			fahrenheit = mc
		}
	}
	return celcius, fahrenheit
}
