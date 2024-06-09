package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	cnst "github.com/LostArrows27/snippetbox/pkg/const"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = newColoredLogger(cnst.GreenBackground, "INFO", log.LstdFlags)
	errorLogger = newColoredLogger(cnst.RedBackground, "ERROR", log.LstdFlags)
}

func newColoredLogger(color, prefix string, flag int) *log.Logger {
	paddedPrefix := color + " " + prefix + " " + cnst.Reset
	return log.New(os.Stdout, paddedPrefix, flag)
}

func Info(format string, v ...interface{}) {
	logMessage(infoLogger, format, v...)
}

func Error(err error) {
	if err != nil {
		logMessage(errorLogger, err.Error())
		os.Exit(1)
	}
}

func logMessage(logger *log.Logger, format string, v ...interface{}) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	message := fmt.Sprintf(format, v...)
	logType := strings.TrimSpace(logger.Prefix())

	formattedMessage := fmt.Sprintf("%s %s %s", timestamp, logType, strings.TrimSpace(message))
	fmt.Println(formattedMessage)
}
