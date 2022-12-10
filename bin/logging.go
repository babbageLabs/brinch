package bin

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Logger.SetOutput(os.Stdout)
	//Logger.SetReportCaller(true)
	Logger.WithFields(logrus.Fields{
		"app": "brinch",
	})
}
