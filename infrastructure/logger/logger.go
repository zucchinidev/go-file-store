package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type logEvent struct{ msg string }

type Standard struct {
	*logrus.Logger
}

func New() *Standard {
	l := &Standard{logrus.New()}
	fm := logrus.FieldMap{
		logrus.FieldKeyTime:  "time",
		logrus.FieldKeyLevel: "level_name",
		logrus.FieldKeyMsg:   "message",
		logrus.FieldKeyFunc:  "caller",
	}
	l.Formatter = &logrus.JSONFormatter{FieldMap: fm}
	l.Out = os.Stdout
	l.Level = logrus.InfoLevel
	return l
}

const serviceName = "file-store service"

// Declare variables to store log messages as new log Events
var (
	httpServerMsg              = logEvent{"%s - Initializing http status server: %s"}
	fatalMsg                   = logEvent{"%s - Fatal error: %s"}
	brokerMsg                  = logEvent{"%s - Broker error: %s"}
	fatalBrokerMsg             = logEvent{"%s - Fatal Broker error: %s"}
	shutdownServiceMsg         = logEvent{"%s - Gracefully shutting down"}
	httpServerShutdownMsg      = logEvent{"%s - HTTP server gracefully shutting down"}
	httpServerErrorMsg         = logEvent{"%s - HTTP server error %s"}
	brokerShutdownMsg          = logEvent{"%s - Broker gracefully shutting down"}
	receivedInterruptSignalMsg = logEvent{"%s - Received interrupt signal, stopping gracefully"}
	httpServerShutdownErrorMsg = logEvent{"%s - Error closing listeners, or context timeout: %s"}
)

// HTTPServerInitialization is a standard error message occurred when the
// http server has a problem in the initialization
func (l *Standard) HTTPServerInitialization(err error) {
	l.Errorf(httpServerMsg.msg, serviceName, err)
}

// Fatal error, it exists process after log info
func (l *Standard) FatalError(err error) {
	l.Fatalf(fatalMsg.msg, serviceName, err)
}

// ShowBrokerError shows the errors occurred in the message broker
func (l *Standard) ShowBrokerError(err error) {
	l.Errorf(brokerMsg.msg, serviceName, err)
}

// ShowFatalBrokerError shows the fatal errors occurred in the message broker
func (l *Standard) ShowFatalBrokerError(err error) {
	l.Errorf(fatalBrokerMsg.msg, serviceName, err)
}

// ShutdownService shows the info about the shutdown of the service
func (l *Standard) ShutdownService() {
	l.Infof(shutdownServiceMsg.msg, serviceName)
}

// HTTPServerShutdown shows the info about the shutdown of the http server
func (l *Standard) HTTPServerShutdown() {
	l.Infof(httpServerShutdownMsg.msg, serviceName)
}

// HTTPServerShutdownError shows the error from closing listeners, or context timeout
func (l *Standard) HTTPServerShutdownError(err error) {
	l.Infof(httpServerShutdownErrorMsg.msg, serviceName, err)
}

// HTTPServerError shows the errors occurred in the http server initialization
func (l *Standard) HTTPServerError(err error) {
	l.Errorf(httpServerErrorMsg.msg, serviceName, err)
}

// BrokerShutdown shows the info about the shutdown of the message broker
func (l *Standard) BrokerShutdown() {
	l.Infof(brokerShutdownMsg.msg, serviceName)
}

// ReceivedInterruptSignal shows the info about the received interrupt signal
func (l *Standard) ReceivedInterruptSignal() {
	l.Infof(receivedInterruptSignalMsg.msg, serviceName)
}
