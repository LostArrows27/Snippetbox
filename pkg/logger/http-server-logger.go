package logger

import (
	"fmt"
	"log"
	"os"
)

type CustomLogger struct {
	*log.Logger
}

func (c *CustomLogger) getLogger(prefix string) *log.Logger {
	switch prefix {
	case "INFO":
		return infoPrinter
	case "ERROR":

		return errorPrinter
	default:

		return c.Logger
	}
}

func (c *CustomLogger) Print(v ...interface{}) {
	logger := c.getLogger(c.Logger.Prefix())

	logMessage(logger, fmt.Sprint(v...))
}

func (c *CustomLogger) Printf(format string, v ...interface{}) {
	logger := c.getLogger(c.Logger.Prefix())

	logMessage(logger, format, v...)

}

func (c *CustomLogger) Println(v ...interface{}) {
	logger := c.getLogger(c.Logger.Prefix())

	logMessage(logger, fmt.Sprint(v...))
}

func ErrorLogger() *CustomLogger {
	logger := log.New(os.Stderr, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger}
}

func InfoLogger() *CustomLogger {
	logger := log.New(os.Stderr, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger}
}
