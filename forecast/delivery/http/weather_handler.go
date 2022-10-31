package http

import (
	"net/http"

	"github.com/adibaulia/bmkg-weather/domain"
	"github.com/gin-gonic/gin"
)

type weatherHandler struct {
	weatherUseCase domain.WeatherUseCase
}

type response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewWeatherHandler(r *gin.Engine, us domain.WeatherUseCase) {
	handler := &weatherHandler{us}
	r.SetTrustedProxies(nil)
	r.GET("/weather/:province", handler.GetWeatherByProvince)
	r.GET("/weather/:province/:city", handler.GetWeatherByCity)

}

func (w *weatherHandler) GetWeatherByProvince(c *gin.Context) {
	province := c.Param("province")
	weatherForecasts, err := w.weatherUseCase.GetWeatherByProvince(province)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       weatherForecasts,
	})
}

func (w *weatherHandler) GetWeatherByCity(c *gin.Context) {
	province := c.Param("province")
	city := c.Param("city")
	weatherForecasts, err := w.weatherUseCase.GetWeatherByCity(province, city)
	if err != nil {
		c.JSON(http.StatusBadRequest, response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       weatherForecasts,
	})
}
