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

	weatherRepo := bmkg.NewBMKGClient()
	weatherUseCase := usecase.NewWeatherUseCase(weatherRepo)

	http.NewWeatherHandler(r, weatherUseCase)

	r.Run(fmt.Sprintf(":%v", 8899))
}
