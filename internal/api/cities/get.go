package cities

import (
	"net/http"

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
		"cities": cities,
	})
}
