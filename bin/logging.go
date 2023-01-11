package bin

import (
	"github.com/babbageLabs/brinch/bin/core/methods"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Logger = logrus.New()

func init() {
	//Logger.SetFormatter(&logrus.JSONFormatter{
	//	PrettyPrint: true,
	//})

	Logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             false,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
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
		return logrus.DebugLevel
	}
	return lvl
}

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := methods.GetDurationInMillseconds(start)

		entry := Logger.WithFields(logrus.Fields{
			"client_ip":  methods.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
			// "api_version": util.ApiVersion,
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
