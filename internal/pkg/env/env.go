package env

import (
	"os"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm/logger"
)

type (
	Env string
)

const (
	production Env = "production"
	staging    Env = "staging"
	local      Env = "local"

	noEnv = "No env"
)

func Prepare() Env {
	e := Env(os.Getenv("API_ENV"))
	switch e {
	case local, staging, production:
		return e
	default:
		panic(noEnv)
	}
}

func (e Env) IsLocal() bool {
	return e == local
}

func (e Env) LogLevel() log.Lvl {
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

func (e Env) GormLogger() logger.Interface {
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
