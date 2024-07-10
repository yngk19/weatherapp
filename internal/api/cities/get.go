package cities

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type City struct {
	ID    int
	Name  string
	State string
}

func (api *API) GetCities(c *gin.Context) {
	citiesDomain, err := api.service.GetCities(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		api.logger.Err(err)
		return
	}
	var cities []City
	for _, city := range citiesDomain {
		cities = append(cities, City{
			ID:    city.ID,
			Name:  city.Name,
			State: city.State,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"cities": cities,
	})
}

func (api *API) GetShortForecast(c *gin.Context) {
	cityID, err := strconv.Atoi(c.Param("cityID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	forecast, err := api.service.GetShortByCityID(c, cityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stats":          "success",
		"short_forecast": forecast,
	})
}
