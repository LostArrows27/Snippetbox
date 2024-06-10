package logger

import (
	"fmt"
	"log"
	"os"
)

type CustomLogger struct {
	*log.Logger
}

func (c *CustomLogger) Print(v ...interface{}) {
	logMessage(c.Logger, fmt.Sprint(v...))
}

func (c *CustomLogger) Printf(format string, v ...interface{}) {
	logMessage(c.Logger, format, v...)
}

func NewErrorLogger() *CustomLogger {
	logger := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger}
}
