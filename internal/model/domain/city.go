package domain

type Town struct {
	ID       int
	Name     string
	Country  string
	Lat      float64
	Lon      float64
	Forecast []WeatherForecast
}
