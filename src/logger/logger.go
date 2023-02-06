package logger

import (
	"app/src/config"
	"io"
	"log"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	ERROR int = 0
	INFO  int = 1
	DEBUG int = 2
)

var (
	logConf     config.Config
	appLogger   *log.Logger
	errorLogger *log.Logger
	level       int
)

func Init(conf *config.Config) error {
	logConf = *conf
	switch conf.LogConf.LogLevel {
	case "ERROR":
		level = ERROR
	case "INFO":
		level = INFO
	case "DEBUG":
		level = DEBUG
	default:
		log.Fatal("Wrong log level")
	}
	appLogWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{
		Filename: os.Getenv("LOG_PATH") + "/app.log",
		// Filename:   logConf.LogConf.LogFilePath + "/app.log",
		MaxSize:    logConf.LogConf.MaxSize,
		MaxBackups: logConf.LogConf.MaxBackupTerms,
		MaxAge:     logConf.LogConf.MaxAge,
	})
	appLogger = log.New(appLogWriter, "", log.LstdFlags)
	errorLogWriter := io.MultiWriter(os.Stderr, &lumberjack.Logger{
		Filename: os.Getenv("LOG_PATH") + "/error.log",
		// Filename:   logConf.LogConf.LogFilePath + "/error.log",
		MaxSize:    logConf.LogConf.MaxSize,
		MaxBackups: logConf.LogConf.MaxBackupTerms,
		MaxAge:     logConf.LogConf.MaxAge,
	})
	errorLogger = log.New(errorLogWriter, "", log.LstdFlags)
	return nil
}

type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Error(format string, v ...interface{})
}

func Debug(format string, v ...interface{}) {
	if level >= DEBUG {
		format = "[DEBUG] " + format
		appLogger.Printf(format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if level >= INFO {
		format = "[INFO] " + format
		appLogger.Printf(format, v...)
	}
}

func Error(format string, v ...interface{}) {
	if level >= ERROR {
		format = "[ERROR] " + format
		appLogger.Printf(format, v...)
		errorLogger.Printf(format, v...)
	}
}
