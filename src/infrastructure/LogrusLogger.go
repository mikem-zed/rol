package infrastructure

import (
	"os"
	"rol/app/errors"
	"rol/app/interfaces"
	"rol/domain"

	"github.com/sirupsen/logrus"
)

//NewLogrusLogger creates new instance of the logrus logger with two hooks for end-to-end logging
//Params
//	config - yaml configuration struct *domain.AppConfig
//Return
//	*logrus.Logger - logrus logger instance
//	error - if error occurs return error, otherwise nil
func NewLogrusLogger(config *domain.AppConfig) (*logrus.Logger, error) {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&Formatter{})

	var err error
	logger.Level, err = logrus.ParseLevel(config.Logger.Level)
	if err != nil {
		return nil, errors.Internal.Wrap(err, "failed to parse logrus level")
	}
	return logger, nil
}

//RegisterLogHooks registers logrus hooks which will duplicate all logs to database
func RegisterLogHooks(logger *logrus.Logger, httpLogRepo interfaces.IGenericRepository[domain.HTTPLog], logRepo interfaces.IGenericRepository[domain.AppLog], config *domain.AppConfig) {
	if config.Logger.LogsToDatabase {
		httpHook := NewHTTPHook(httpLogRepo)
		appHook := NewAppHook(logRepo)
		logger.AddHook(httpHook)
		logger.AddHook(appHook)
	}
}
