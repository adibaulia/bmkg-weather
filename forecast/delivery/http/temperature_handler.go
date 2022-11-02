package http

import (
	"net/http"

	"github.com/adibaulia/bmkg-weather/domain"

	"github.com/gin-gonic/gin"
)

type temperatureHandler struct {
	temperatureUseCase domain.TemperatureUseCase
}

func NewTemperatureHandler(r *gin.Engine, us domain.TemperatureUseCase) {
	handler := &temperatureHandler{us}
	r.SetTrustedProxies(nil)
	r.GET("/temperature/:province", handler.GettemperatureByProvince)
	r.GET("/temperature/:province/:city", handler.GettemperatureByCity)

}

func (w *temperatureHandler) GettemperatureByProvince(c *gin.Context) {
	province := c.Param("province")
	temperatureForecasts, err := w.temperatureUseCase.GetTemperatureByProvince(province)
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
		Data:       temperatureForecasts,
	})
}

func (w *temperatureHandler) GettemperatureByCity(c *gin.Context) {
	province := c.Param("province")
	city := c.Param("city")
	temperatureForecasts, err := w.temperatureUseCase.GetTemperatureByCity(province, city)
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
		Data:       temperatureForecasts,
	})
}
