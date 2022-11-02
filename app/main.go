package main

import (
	"fmt"

	"github.com/adibaulia/bmkg-weather/forecast/delivery/http"
	"github.com/adibaulia/bmkg-weather/forecast/repository/bmkg"
	"github.com/adibaulia/bmkg-weather/forecast/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	forecastRepo := bmkg.NewBMKGClient()
	weatherUseCase := usecase.NewWeatherUseCase(forecastRepo)

	http.NewWeatherHandler(r, weatherUseCase)

	temperatureUseCase := usecase.NewTemperatureUseCase(forecastRepo)
	http.NewTemperatureHandler(r, temperatureUseCase)

	r.Run(fmt.Sprintf(":%v", 8899))
}
