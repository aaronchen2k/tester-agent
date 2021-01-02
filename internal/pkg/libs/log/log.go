package _logUtils

import (
	"fmt"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"runtime"
	"strings"
)

var logger *logrus.Logger

func InitLogger() *logrus.Logger {
	if logger != nil {
		return logger
	}

	logger = logrus.New()

	logPath, _ := os.Getwd()
	logName := logPath + "/%Y%m%d.log"
	r, _ := rotatelogs.New(logName)
	mw := io.MultiWriter(os.Stdout, r)
	logger.SetOutput(mw)

	logger.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05",
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logger.SetFormatter(formatter)

	return logger
}

func formatFilePath(path string) string {
	index := strings.Index(path, "src")
	return path[index:]
}
