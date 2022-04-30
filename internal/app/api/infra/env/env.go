package env

import "gorm.io/gorm/logger"

type (
	Env interface {
		GormLogger() logger.Interface
		IsLocal() bool
	}
)
