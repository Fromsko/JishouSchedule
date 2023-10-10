package utils

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// CustomTextFormatter 自定义的日志格式
type CustomTextFormatter struct{}

func (f *CustomTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	level := entry.Level.String()
	callerInfo := entry.Message

	// 手动设置颜色
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel:
		levelColor = 36 // Cyan
	case logrus.InfoLevel:
		levelColor = 32 // Green
	case logrus.WarnLevel:
		levelColor = 33 // Yellow
	case logrus.ErrorLevel:
		levelColor = 31 // Red
	default:
		levelColor = 0 // Default color
	}

	return []byte(fmt.Sprintf("\x1b[1;%dm%s | %s | %s\x1b[0m\n", levelColor, timestamp, level, callerInfo)), nil
}

func InitLogger() *logrus.Logger {
	logger := logrus.New()

	// 设置日志输出格式为自定义格式
	logger.SetFormatter(&CustomTextFormatter{})

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
