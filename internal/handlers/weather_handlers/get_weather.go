package weather_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zob456/weather_api/internal/envs"
	"github.com/zob456/weather_api/internal/models"
	"github.com/zob456/weather_api/internal/utils"
	"net/http"
	"slices"
	"time"
)

const weatherURL = "https://api.openweathermap.org/data/2.5/weather?"

func GetWeather(c echo.Context) error {
	var req models.GetWeatherRequest
	err := (&echo.DefaultBinder{}).BindBody(c, &req)
	if err != nil {
		utils.LocalErrorLogger(err)
		return utils.BadRequest(c)
	}

	message := "success"

	unit := models.ImperialUnit
	if req.Unit != nil && slices.Contains(models.AcceptableUnits, *req.Unit) {
		unit = *req.Unit
	}

	if req.Unit != nil && !slices.Contains(models.AcceptableUnits, *req.Unit) {
		message += " - the unit of measure passed in is not in the acceptable lists so defaulting to 'imperial'"
	}

	reqURL := fmt.Sprintf("%vlat=%v&lon=%v&appid=%v&units=%v", weatherURL, req.Lat, req.Lon, envs.APIKey, unit)

	client := http.Client{Timeout: time.Duration(3) * time.Second}
	resp, err := client.Get(reqURL)
	if err != nil {
		utils.LocalErrorLogger(err)
		return utils.InternalServerErr(c)
	}
	defer resp.Body.Close()

	weather := models.Weather{}
	err = json.NewDecoder(resp.Body).Decode(&weather)
	if err != nil {
		utils.LocalErrorLogger(err)
		return utils.InternalServerErr(c)
	}

	weatherResp := models.WeatherResp{
		Temp:        weather.Main.Temp,
		FeelsLike:   weather.Main.FeelsLike,
		Clouds:      weather.Clouds.All,
		Humidity:    weather.Main.Humidity,
		WindSpeed:   weather.Wind.Speed,
		Description: weather.Weather[0].Description,
	}

	return utils.RespondOK(c, message, weatherResp)
}
