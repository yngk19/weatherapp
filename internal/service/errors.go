package service

import "errors"

var (
	ErrCitiesListIsEmpty     = errors.New("cities list is empty")
	ErrNoForecasts           = errors.New("there is no forecasts")
	ErrNoForecastForThisDate = errors.New("no forecast for this day")
	ErrNoSuchCity            = errors.New("no such city")
)
