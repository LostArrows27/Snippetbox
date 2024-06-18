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

func (c *CustomLogger) Fatal(v ...interface{}) {
	logger := c.getLogger(c.Logger.Prefix())

	logMessage(logger, fmt.Sprint(v...))

	os.Exit(1)
}

func (c *CustomLogger) Fatalf(format string, v ...interface{}) {
	logger := c.getLogger(c.Logger.Prefix())

	logMessage(logger, fmt.Sprintf(format, v...))

	os.Exit(1)
}

func (c *CustomLogger) Fatalln(v ...interface{}) {
	logger := c.getLogger(c.Logger.Prefix())

	logMessage(logger, fmt.Sprint(v...))

	os.Exit(1)
}

func ErrorLogger() *CustomLogger {
	logger := log.New(os.Stderr, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger}
}

func InfoLogger() *CustomLogger {
	logger := log.New(os.Stderr, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	return &CustomLogger{logger}
}
