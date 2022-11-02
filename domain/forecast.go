package domain

import (
	"github.com/adibaulia/bmkg-weather/domain/location"
	"github.com/adibaulia/bmkg-weather/forecast/repository/bmkg"
)

type ForecastRepository interface {
	GetByProvince(province location.Province) (*bmkg.BMKGForecast, error)
}
