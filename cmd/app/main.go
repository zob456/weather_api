package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/zob456/weather_api/internal/envs"
	"github.com/zob456/weather_api/internal/handlers"
	"github.com/zob456/weather_api/internal/utils"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	envs.CheckSetEnvs()
}

func main() {
	utils.InfoLogger(fmt.Sprintf("ENV set to: %v", envs.ENV))
	e := echo.New()
	handlers.Routes(e)

	utils.FatalLogger(e.Start(fmt.Sprintf(":%v", "8080")))
}
