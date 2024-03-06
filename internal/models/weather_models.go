package models

// WeatherAPIUnit type using `=` to avoid requiring type assertion
type WeatherAPIUnit = string

const (
	StandardUnit WeatherAPIUnit = "standard"
	ImperialUnit WeatherAPIUnit = "imperial"
	MetricUnit   WeatherAPIUnit = "metric"
)

var AcceptableUnits = []WeatherAPIUnit{StandardUnit, ImperialUnit, MetricUnit}

type Weather struct {
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Main struct {
		FeelsLike   float64 `json:"feels_like"`
		GroundLevel int     `json:"grnd_level"`
		Humidity    int     `json:"humidity"`
		Pressure    int     `json:"pressure"`
		SeaLevel    int     `json:"sea_level"`
		Temp        float64 `json:"temp"`
		TempMax     float64 `json:"temp_max"`
		TempMin     float64 `json:"temp_min"`
		Name        string  `json:"name"`
	} `json:"main"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Weather []struct {
		Description string `json:"description"`
		Icon        string `json:"icon"`
		ID          int    `json:"id"`
		Main        string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

type GetWeatherRequest struct {
	Lat  float64         `json:"lat"`
	Lon  float64         `json:"lon"`
	Unit *WeatherAPIUnit `json:"unit"`
}

type WeatherResp struct {
	Temp        float64 `json:"temp"`
	FeelsLike   float64 `json:"feels_like"`
	Clouds      int     `json:"clouds"`
	Humidity    int     `json:"humidity"`
	WindSpeed   float64 `json:"wind_speed"`
	Description string  `json:"description"`
}
