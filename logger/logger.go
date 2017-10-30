package logger

import (
	"os"

	logging "github.com/op/go-logging"
)

var log = logging.MustGetLogger("lengine")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	//Set Color; print time, callpath, level, id; unset color; print message
	`%{color}%{time:15:04:05.000} %{callpath} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
var backend = logging.NewLogBackend(os.Stderr, "", 0)

func init() {
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

func Critical(message string) {
	log.Critical(message)
}

func Error(message string) {
	log.Error(message)
}

func Warning(message string) {
	log.Warning(message)
}

func Notice(message string) {
	log.Notice(message)
}

func Info(message string) {
	log.Info(message)
}

func Debug(message string) {
	log.Debug(message)
}
