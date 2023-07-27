package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	Fatal LogLevel = iota
	Error
	Warn
	Info
	Debug
)

type Logger struct {
	maxLevel LogLevel
	writer   *os.File
}

func NewLogger() *Logger {
	return &Logger{maxLevel: Info}
}

func (l *Logger) SetMaxLogLevel(level LogLevel) {
	l.maxLevel = level
}

func (l *Logger) SetLogFile(filePath string) error {
	if l.writer != nil {
		_ = l.writer.Close()
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	l.writer = file
	return nil
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level <= l.maxLevel {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		message := fmt.Sprintf(format, args...)
		output := fmt.Sprintf("[%s] [%s] %s", timestamp, strings.ToUpper(levelToString(level)), message)

		if l.writer != nil {
			_, _ = l.writer.WriteString(output + "\n")
		} else {
			fmt.Println(output)
		}
	}
}

func levelToString(level LogLevel) string {
	switch level {
	case Fatal:
		return "FATAL"
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(Fatal, format, args...)
	os.Exit(1)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(Error, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(Warn, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(Info, format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(Debug, format, args...)
}
