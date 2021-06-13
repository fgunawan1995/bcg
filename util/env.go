package util

import (
	"os"

	"github.com/fgunawan1995/bcg/model"
)

func GetEnv() string {
	env := os.Getenv(model.EnvKey)
	if env == "" {
		env = "local"
	}
	return env
}
