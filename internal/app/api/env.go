package api

import (
	"os"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm/logger"
)

type (
	APIEnv string
)

const (
	production APIEnv = "production"
	staging    APIEnv = "staging"
	local      APIEnv = "local"
)

var apiEnv APIEnv

func Env() APIEnv {
	return apiEnv
}

func prepare() APIEnv {
	apiEnvTmp := APIEnv(os.Getenv("API_ENV"))
	switch apiEnvTmp {
	case local, staging, production:
		apiEnv = apiEnvTmp
	default:
		panic("Unregistered API_ENV")
	}

	return apiEnvTmp
}

func (e APIEnv) IsLocal() bool {
	return e == local
}

func (e APIEnv) LogLevel() log.Lvl {
	switch e {
	case local:
		return log.DEBUG
	case staging:
		return log.WARN
	case production:
		return log.WARN
	default:
		panic("No env")
	}
}

func (e APIEnv) GormLogger() logger.Interface {
	switch e {
	case local:
		return logger.Default
	case staging:
		return logger.Discard
	case production:
		return logger.Discard
	default:
		panic("No env")
	}
}
