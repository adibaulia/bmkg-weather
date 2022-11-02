package usecase

import (
	"log"
	"strings"
	"time"

	"github.com/adibaulia/bmkg-weather/domain"
	"github.com/adibaulia/bmkg-weather/domain/errormessage"
	"github.com/adibaulia/bmkg-weather/domain/location"
	"github.com/adibaulia/bmkg-weather/forecast/repository/bmkg"
)

type weatherUseCase struct {
	forecastRepo domain.ForecastRepository
}

func NewWeatherUseCase(forecastRepo domain.ForecastRepository) domain.WeatherUseCase {
	return &weatherUseCase{
		forecastRepo: forecastRepo,
	}
}

func (u *weatherUseCase) GetWeatherByCity(provinceStr string, cityStr string) (domain.WeatherForecast, error) {

	provinceWeatherForecast, err := u.GetWeatherByProvince(provinceStr)
	if err != nil {
		log.Printf("error when getting province weather forecast: %v", err)
		return domain.WeatherForecast{}, err
	}

	for _, p := range provinceWeatherForecast {
		if strings.EqualFold(p.Location.Description, cityStr) {
			return p, nil
		}
	}

	return domain.WeatherForecast{}, errormessage.ErrNotFound
}

func (u *weatherUseCase) GetWeatherByProvince(provinceStr string) ([]domain.WeatherForecast, error) {
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")
	province := location.Province(provinceStr)
	BMKGForecast, err := u.forecastRepo.GetByProvince(province)
	if err != nil {
		log.Printf("error getting weather forecast: %v", err)
		return nil, err
	}

	lastUpdate, err := time.ParseInLocation("20060102150405", BMKGForecast.Forecast.Issue.Timestamp, time.UTC)
	if err != nil {
		log.Printf("error parsing time lastUpdate from BMKGForecast: %v", err)
		return nil, err
	}

	var result []domain.WeatherForecast

	for _, l := range BMKGForecast.Forecast.Area {
		weatherPerLocation := domain.WeatherForecast{
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

		var weatherForecasts []domain.WeatherForecastTimeRange
		for _, p := range l.Parameter {
			if p.ID == "weather" {
				for _, tr := range p.Timerange {
					datetime, err := time.Parse("200601021504", tr.Datetime)
					if err != nil {
						return nil, err
					}

					weatherCode := domain.GetWeatherCode(tr.Value[0].Text)

					weatherForecast := domain.WeatherForecastTimeRange{
						Weather: domain.Weather{
							WeatherCode: weatherCode,
							NameID:      domain.GetWeatherName(weatherCode, "id"),
							NameEN:      domain.GetWeatherName(weatherCode, "en"),
						},
						Date:           domain.DateFormatter(datetime),
						TimeStartRange: domain.TimeFormatter(datetime),
						TimeEndRange:   domain.TimeFormatter(datetime.Add(6 * time.Hour)),
					}
					weatherForecasts = append(weatherForecasts, weatherForecast)

				}
			}
		}
		weatherPerLocation.WeatherForecastTimeRange = weatherForecasts
		result = append(result, weatherPerLocation)
	}

	return result, nil
}

func nameLangCompose(nameLang []bmkg.Name) []location.NameLang {
	var result []location.NameLang
	for _, v := range nameLang {
		result = append(result, location.NameLang{
			Name: v.Text,
			Lang: v.Lang,
		})
	}

	return result
}
