package cities

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	city, err := api.service.GetCityByID(c, cityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         "success",
		"city_id":        cityID,
		"name":           city.Name,
		"country":        city.Country,
		"state":          city.State,
		"lat":            city.Lat,
		"lon":            city.Lon,
		"short_forecast": forecast,
	})
}

func (api *API) GetForecast(c *gin.Context) {
	dateParam := c.Query("date")
	timeParam := c.Query("time")
	cityID, err := strconv.Atoi(c.Param("cityID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err,
		})
		return
	}
	city, err := api.service.GetCityByID(c, cityID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    err.Error(),
		})
		return
	}
	if dateParam != "" {
		forecast, err := api.service.GetForecastByDate(c, dateParam, cityID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"msg":    err,
			})
			return
		}
		if timeParam != "" {
			for _, forecastAtTime := range forecast.DetailInfo {
				forecastTime := strings.Split(forecastAtTime.DtTxt, " ")[1]
				if forecastTime == timeParam {
					c.JSON(http.StatusOK, gin.H{
						"status":  "success",
						"city_id": cityID,
						"name":    city.Name,
						"country": city.Country,
						"state":   city.State,
						"lat":     city.Lat,
						"lon":     city.Lon,
						fmt.Sprintf("%s %s", dateParam, timeParam): forecastAtTime,
					})
					return
				}
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"msg":    "there is no forecast at this time",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"city_id": cityID,
			"name":    city.Name,
			"country": city.Country,
			"state":   city.State,
			"lat":     city.Lat,
			"lon":     city.Lon,
			dateParam: forecast,
		})
	} else {
		forecast5days, err := api.service.GetForecastByCityID(c, cityID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"msg":    err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"city_id":   cityID,
			"name":      city.Name,
			"country":   city.Country,
			"state":     city.State,
			"lat":       city.Lat,
			"lon":       city.Lon,
			"forecasts": forecast5days,
		})
	}

}
