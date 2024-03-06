package weather_handlers

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/zob456/weather_api/internal/envs"
	"github.com/zob456/weather_api/internal/models"
	"github.com/zob456/weather_api/internal/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	e            = echo.New()
	imperialUnit = models.ImperialUnit
	badUnit      = "bad unit"

	happyPathImperialUnitReq = models.GetWeatherRequest{
		Lat:  28.35611111111111,
		Lon:  -82.69333333333334,
		Unit: &imperialUnit,
	}

	happyPathNilUnitReq = models.GetWeatherRequest{
		Lat:  28.35611111111111,
		Lon:  -82.69333333333334,
		Unit: nil,
	}

	happyPathBadUnitReq = models.GetWeatherRequest{
		Lat:  28.35611111111111,
		Lon:  -82.69333333333334,
		Unit: &badUnit,
	}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	envs.CheckSetEnvs()
}

func Test_GetWeather_Happy_Path(t *testing.T) {
	reqBytes, err := json.Marshal(happyPathImperialUnitReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = GetWeather(c)
	assert.NoError(t, err)

	weatherResp := utils.EchoMap{}
	err = json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&weatherResp)
	assert.NoError(t, err)

	assert.Equal(t, "success", weatherResp.Message)
}

func Test_GetWeather_Happy_Path_Nil_Unit(t *testing.T) {
	reqBytes, err := json.Marshal(happyPathNilUnitReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = GetWeather(c)
	assert.NoError(t, err)

	weatherResp := utils.EchoMap{}
	err = json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&weatherResp)
	assert.NoError(t, err)

	assert.Equal(t, "success", weatherResp.Message)
}

func Test_GetWeather_Happy_Path_Bad_Unit(t *testing.T) {
	reqBytes, err := json.Marshal(happyPathBadUnitReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = GetWeather(c)
	assert.NoError(t, err)

	weatherResp := utils.EchoMap{}
	err = json.NewDecoder(strings.NewReader(rec.Body.String())).Decode(&weatherResp)
	assert.NoError(t, err)

	assert.Equal(t, "success - the unit of measure passed in is not in the acceptable lists so defaulting to 'imperial'", weatherResp.Message)
}

func Test_GetWeather_Sad_Path_Bad_Req(t *testing.T) {
	var badReq interface{}
	reqBytes, err := json.Marshal(badReq)
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err = GetWeather(c)
	assert.NoError(t, err)
}
