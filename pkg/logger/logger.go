package logger

import (
	"os"
	"strings"

	"github.com/lbrictson/squint/pkg/conf"
	log "github.com/sirupsen/logrus"
)

func New(config conf.Config) *log.Logger {
	l := log.New()
	if config.LogJSON {
		l.SetFormatter(&log.JSONFormatter{})
	}
	// Handle file logging, if we can't log to a file proceed anyway
	if config.LogFile != "" {
		file, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			l.Out = file
		} else {
			l.Errorf("Unable to log to file %v %v, defaulting to stdout logger",
				config.LogFile, err)
		}
	}
	switch strings.ToUpper(config.LogLevel) {
	case "INFO":
		l.SetLevel(log.InfoLevel)
	case "WARN":
		l.SetLevel(log.WarnLevel)
	case "ERROR":
		l.SetLevel(log.ErrorLevel)
	case "DEBUG":
		l.SetLevel(log.DebugLevel)
	default:
		l.SetLevel(log.InfoLevel)
	}
	return l
}
