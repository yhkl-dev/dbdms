package helper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var ErrorLogger, AccessLogger, WorkLogger, SQLLogger *logrus.Logger

func Logger(outpath string) *logrus.Logger {
	logger := logrus.New()
	_, err := os.Stat(outpath)
	if os.IsNotExist(err) {
		os.Create(outpath)
	}

	file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		logger.Out = file
	} else {
		logger.Info("Open" + outpath + "'s log file failed, use default log format'")
	}
	return logger
}

func init() {
	WorkLogger = Logger("logs/work.log")
	AccessLogger = Logger("logs/access.log")
	ErrorLogger = Logger("logs/error.log")
	SQLLogger = Logger("logs/sql.log")
}
