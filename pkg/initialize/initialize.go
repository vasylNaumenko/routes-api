/*
 * telegram: @VasylNaumenko
 */

package initialize

import (
	"os"

	"routes-api/pkg/config"
	"routes-api/pkg/log"
)

// Logger inits log.Logger
func Logger(cfg *config.Logger, serviceName string) (log.Logger, error) {
	logOut := log.Output(os.Stdout)

	if cfg.OutputFilePath != "" {
		f, err := os.OpenFile(cfg.OutputFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		logOut = log.Output(f)
	}

	logOptions := []log.Option{
		logOut,
		log.Tags(map[string]interface{}{"service": serviceName}),
		log.Level(cfg.DebugLevel),
		log.Formatter(log.Format(cfg.LogFormat), false, "2006-01-02 15:04:05"),
	}
	if cfg.IncludeCallerMethod {
		logOptions = append(logOptions, log.WithCallerReporting())
	}

	return log.New(logOptions...), nil
}
