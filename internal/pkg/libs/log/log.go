package _logUtils

import (
	_vari "github.com/aaronchen2k/tester/internal/pkg/vari"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/log"
	"io/ioutil"
	"os"
	"os/user"
	"path"
	"strings"
)

var logger *logrus.Logger

func Init(app string) {
	if logger != nil {
		return
	}

	usr, _ := user.Current()
	log.Info("Run as user " + usr.Username)

	_vari.WorkDir = addPathSepIfNeeded(path.Join(usr.HomeDir, "tester"))
	logDir := addPathSepIfNeeded(path.Join(_vari.WorkDir, "log"))
	MkDirIfNeeded(logDir)
	log.Info("Log dir is " + logDir)

	logger = logrus.New()
	logger.Out = ioutil.Discard

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  logDir + app + "-log.txt",
		logrus.WarnLevel:  logDir + app + "-log.txt",
		logrus.ErrorLevel: logDir + app + "-error.txt",
	}

	logger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&MyFormatter{},
	))

	logger.SetFormatter(&MyFormatter{})

	return
}

type MyFormatter struct {
	logrus.TextFormatter
}

func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message + "\n"), nil
}

func addPathSepIfNeeded(pth string) string {
	sepa := string(os.PathSeparator)

	if strings.LastIndex(pth, sepa) < len(pth)-1 {
		pth += sepa
	}
	return pth
}

func MkDirIfNeeded(dir string) error {
	if !FileExist(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		return err
	}

	return nil
}
func FileExist(path string) bool {
	var exist = true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
