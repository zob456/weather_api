package utils

import (
	"fmt"
	"github.com/zob456/weather_api/internal/envs"
	"log"
)

func ErrorLogger(err error, calldepth ...any) {
	defaultCalldepth := 2

	if len(calldepth) > 0 {
		calldepthInt, ok := calldepth[0].(int)
		if ok {
			defaultCalldepth = calldepthInt
		}
	}
	_ = log.Output(defaultCalldepth, fmt.Sprintf("[ERROR] %v", err))
}

func InfoLogger(message any, calldepth ...any) {
	defaultCalldepth := 2

	if len(calldepth) > 0 {
		calldepthInt, ok := calldepth[0].(int)
		if ok {
			defaultCalldepth = calldepthInt
		}
	}

	_ = log.Output(defaultCalldepth, fmt.Sprintf("[INFO] %v", message))
}

func FatalLogger(err error, calldepth ...any) {
	defaultCalldepth := 2

	if len(calldepth) > 0 {
		calldepthInt, ok := calldepth[0].(int)
		if ok {
			defaultCalldepth = calldepthInt
		}
	}

	_ = log.Output(defaultCalldepth, fmt.Sprintf("[FATAL] %v", err))
	log.Fatal(err)
}

// LocalErrorLogger added for local debugging but not for prod
func LocalErrorLogger(err error, calldepth ...any) {
	defaultCalldepth := 2
	if envs.ENV == "local" {

		if len(calldepth) > 0 {
			calldepthInt, ok := calldepth[0].(int)
			if ok {
				defaultCalldepth = calldepthInt
			}
		}
		_ = log.Output(defaultCalldepth, fmt.Sprintf("[ERROR] %v", err))
	}
}
