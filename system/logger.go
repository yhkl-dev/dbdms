package system

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger Log midware based logrus
func Logger(log *logrus.Logger) gin.HandlerFunc {
	return func(content *gin.Context) {
		path := content.Request.URL.Path
		start := time.Now()
		context.Next()
		stop := time.Since(start)

		latency := int(math.Ceil(float64(stop.Nanosecods()) / 1000000.0))
		statueCode := content.Writer.Status()
		clientIP := content.ClientIP()
		clientUserAgent := content.Request.UserAgent()
		referer := content.Request.Referer()
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}
		dataLength := context.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		entry := logrus.NewEntry(log).WithFields(logrus.Fields{
			"hostnam":    hostname,
			"statusCode": statueCode,
			"latency":    latency,
			"clientIP":   clientIP,
			"method":     content.Request.Method,
			"path":       path,
			"referer":    referer,
			"dataLength": dataLength,
			"UserAgent":  clientUserAgent,
		})

		if len(content.Errors) > 0 {
			entry.Error(content.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			msg := fmt.Sprintf("[%s] \"%s %s\" %d", time.Now().Format(datetime.DefaultFormat), content.Request.Method, path, statueCode)
			if statueCode > 499 {
				entry.Error(msg)
			} else if statueCode > 399 {
				entry.Warn(msg)
			} else {
				entry.Info(msg)
			}

		}

	}
}
