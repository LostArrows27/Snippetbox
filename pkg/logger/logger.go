package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	cnst "github.com/LostArrows27/snippetbox/internal/const"
	"github.com/LostArrows27/snippetbox/pkg/env"
)

var (
	infoPrinter  *log.Logger
	errorPrinter *log.Logger
	infoLogger   *log.Logger
	errorLogger  *log.Logger
)

func init() {
	// 0. load .env file
	env.LoadEnv(".env")

	// 1. load info log file
	infoPath := env.GetEnv("INFO_LOG")
	infoLogFile, err := os.OpenFile(infoPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 2. load error log file
	errorPath := env.GetEnv("ERROR_LOG")
	errorLogFile, err := os.OpenFile(errorPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 3. init logger
	infoPrinter = newColoredLogger(os.Stdout, cnst.GreenBackground, "INFO", log.LstdFlags)
	errorPrinter = newColoredLogger(os.Stdout, cnst.RedBackground, "ERROR", log.LstdFlags)
	infoLogger = newColoredLogger(infoLogFile, "", "INFO", log.LstdFlags)
	errorLogger = newColoredLogger(errorLogFile, "", "ERROR", log.LstdFlags)
}

func newColoredLogger(f *os.File, color, prefix string, flag int) *log.Logger {
	paddedPrefix := ""
	if f == os.Stdout || f == os.Stderr {
		paddedPrefix = color + " " + prefix + " " + cnst.Reset
	}
	return log.New(f, paddedPrefix, flag)
}

func Info(format string, v ...interface{}) {
	logMessage(infoPrinter, format, v...)
}

func Error(err error) {
	if err != nil {
		logMessage(errorPrinter, err.Error())
	}
}

func logMessage(logger *log.Logger, format string, v ...interface{}) {
	timestamp := time.Now().Format("2006/01/02 15:04:05")
	message := fmt.Sprintf(format, v...)
	logType := strings.TrimSpace(logger.Prefix())

	formattedMessage := fmt.Sprintf("%s %s %s", timestamp, logType, strings.TrimSpace(message))
	fmt.Println(formattedMessage)

	plainLogType := strings.TrimSpace(strings.Trim(logType, removeDuplicates(cnst.GreenBackground+" "+cnst.RedBackground+" "+cnst.Reset)))
	plainFormattedMessage := fmt.Sprintf("%s %s", plainLogType, strings.TrimSpace(message))

	switch plainLogType {
	case "INFO":
		infoLogger.Println(plainFormattedMessage)
	case "ERROR":
		errorLogger.Println(plainFormattedMessage)
	default:
		break
	}
}

func removeDuplicates(s string) string {
	seen := make(map[rune]bool)
	var result []rune
	for _, r := range s {
		if !seen[r] {
			seen[r] = true
			result = append(result, r)
		}
	}
	return string(result)
}
