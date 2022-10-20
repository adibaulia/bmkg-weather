package usecase

import "github.com/adibaulia/bmkg-weather/domain"

type weatherUseCase struct {
	forecastRepo domain.ForecastRepository
}

func NewWeatherUseCase(forecastRepo domain.ForecastRepository) domain.WeatherUseCase {

	return &weatherUseCase{
		forecastRepo: forecastRepo,
	}
}

func (u *weatherUseCase) GetWeather(location domain.Location) (domain.Weather, error) {
	return domain.Weather{}, nil
}
