package bin

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(os.Stdout)
	Logger.Info("logging with level:", getLogLevel().String())
	Logger.SetLevel(getLogLevel())
	// Logger.SetReportCaller(true)
	Logger.WithFields(logrus.Fields{
		"app": "brinch",
	})
}

func getLogLevel() logrus.Level {
	var lvl logrus.Level
	err := lvl.UnmarshalText([]byte(os.Getenv("config.Logging.Level")))
	if err != nil {
		return logrus.WarnLevel
	}
	return lvl
}
