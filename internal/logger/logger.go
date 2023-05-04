package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Logger struct {
	logFile     *os.File
	logFilePath string
	logFileName string
}

func NewLogger() *Logger {
	logger := &Logger{
		logFilePath: "./log",
		logFileName: "log",
	}
	_ = logger.initLogFile()
	return logger
}

func (l *Logger) initLogFile() error {
	if err := os.MkdirAll(l.logFilePath, 0755); err != nil {
		return err
	}
	now := time.Now()
	logFilePath := filepath.Join(l.logFilePath, fmt.Sprintf("%s_%d-%02d-%02d.log",
		l.logFileName, now.Year(), now.Month(), now.Day()))
	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	l.logFile = logFile
	return nil
}

func (l *Logger) Log(format string, v ...interface{}) {
	now := time.Now().Format("2006-01-02 15:04:05")
	logStr := fmt.Sprintf("[%s] %s\n\r\n", now, fmt.Sprintf(format, v...))
	l.logFile.Write([]byte(logStr))
	l.checkAndSplitLogFile()
}

func (l *Logger) checkAndSplitLogFile() {
	logFilePath := filepath.Join(l.logFilePath, l.logFileName)
	info, err := l.logFile.Stat()
	if err != nil {
		return
	}
	if info.Size() > 1024*1024*10 {
		l.logFile.Close()
		now := time.Now()
		newLogFilePath := filepath.Join(l.logFilePath, fmt.Sprintf("%s_%d-%02d-%02d.log",
			l.logFileName, now.Year(), now.Month(), now.Day()))
		os.Rename(logFilePath, newLogFilePath)
		l.initLogFile()
	}
}
