package module_one

import "github.com/joaosoft/logger"

func One() string {
	return logger.Info("Hello, i'm the module one!").ToError().Error()
}
