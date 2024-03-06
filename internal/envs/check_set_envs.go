package envs

import (
	"log"
	"os"
)

const (
	apiKey = "WEATHER_API_KEY"
	env    = "ENV"
)

var (
	missingENVs []string
	foundENV    bool
	APIKey      string
	ENV         string
)

func CheckSetEnvs() {
	ENV, foundENV = os.LookupEnv(env)
	if !foundENV {
		missingENVs = append(missingENVs, env)
	}

	APIKey, foundENV = os.LookupEnv(apiKey)
	if !foundENV {
		missingENVs = append(missingENVs, apiKey)
	}

	if len(missingENVs) > 0 {
		log.Fatalf("Failed to start due to missing ENVVar(s): %v", missingENVs)
	}
}
