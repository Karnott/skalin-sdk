package skalinsdk

import (
	"os"

	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

type CustomLog struct {
	logrus.FieldLogger
}

var Log = NewLogger()

// init log for skalin
func initLogger() *logrus.Logger {
	newLogger := logrus.New()
	newLogger.SetReportCaller(true)
	newLogger.SetOutput(os.Stdout)
	// set format
	switch os.Getenv("LOG_FORMAT") {
	case "json":
		newLogger.SetFormatter(joonix.NewFormatter())
	default:
		newLogger.SetFormatter(&logrus.TextFormatter{
			QuoteEmptyFields:       true,
			FullTimestamp:          true,
			ForceColors:            true,
			DisableLevelTruncation: true,
		})
	}

	logLevelString := "info"
	if os.Getenv("LOG_LEVEL") != "" {
		logLevelString = os.Getenv("LOG_LEVEL")
	}
	logLevel, err := logrus.ParseLevel(logLevelString)
	if err != nil {
		newLogger.Errorf("log level is not ok, setting to info by default : %v", err.Error())
		newLogger.SetLevel(logrus.InfoLevel)
	} else {
		newLogger.SetLevel(logLevel)
	}
	return newLogger
}

func NewLogger() *CustomLog {
	l := initLogger()
	return &CustomLog{
		l,
	}
}
func (logger *CustomLog) AddSkalinApplicationField() *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"application": "skalin",
	})
}

func (logger *CustomLog) Infof(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Infof(format, args...)
}

func (logger *CustomLog) Printf(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Printf(format, args...)
}

func (logger *CustomLog) Warnf(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Warnf(format, args...)
}

func (logger *CustomLog) Warningf(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Warningf(format, args...)
}

func (logger *CustomLog) Errorf(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Errorf(format, args...)
}

func (logger *CustomLog) Fatalf(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Fatalf(format, args...)
}

func (logger *CustomLog) Panicf(format string, args ...interface{}) {
	logger.AddSkalinApplicationField().Panicf(format, args...)
}
