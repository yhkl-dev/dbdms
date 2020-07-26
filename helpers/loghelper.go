package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var ErrorLogger, AccessLogger, WorkLogger, SQLLogger *logrus.Logger

func Logger(outPath string) *logrus.Logger {
	logger := logrus.New()
	_, err := os.Stat(outPath)
	if os.IsNotExist(err) {
		_, _ = os.Create(outPath)
	}

	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Open" + outPath + "'s log file failed, use default log format'")
	}
	return logger
}

func init() {
	WorkLogger = Logger("logs/work.log")
	AccessLogger = Logger("logs/access.log")
	ErrorLogger = Logger("logs/error.log")
	SQLLogger = Logger("logs/sql.log")
}
