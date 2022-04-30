package env

import (
	"os"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm/logger"
)

type (
	env string
)

const (
	production env = "production"
	staging    env = "staging"
	local      env = "local"

	noEnv = "No env"
)

func Prepare() env {
	e := env(os.Getenv("API_ENV"))
	switch e {
	case local, staging, production:
		return e
	default:
		panic(noEnv)
	}
}

func (e env) IsLocal() bool {
	return e == local
}

func (e env) LogLevel() log.Lvl {
	switch e {
	case local:
		return log.DEBUG
	case staging:
		return log.WARN
	case production:
		return log.WARN
	default:
		panic(noEnv)
	}
}

func (e env) GormLogger() logger.Interface {
	switch e {
	case local:
		return logger.Default
	case staging:
		return logger.Discard
	case production:
		return logger.Discard
	default:
		panic(noEnv)
	}
}
