package _logUtils

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"strings"
)

var ()

func Info(str string) {
	logger.Infoln(str)
}
func Infof(str string, args ...interface{}) {
	logger.Infof(str, args)
}

func Warn(str string) {
	logger.Warnln(str)
}
func Warnf(str string, args ...interface{}) {
	logger.Warnf(str, args)
}

func Error(str string) {
	logger.Errorln(str)
}
func Errorf(str string, args ...interface{}) {
	logger.Errorf(str, args)
}

func Print(str string) {
	logger.Println(str)
}
func Printf(str string, args ...interface{}) {
	logger.Printf(str, args)
}

func PrintUnicode(str []byte) {
	var a interface{}

	temp := strings.Replace(string(str), "\\\\", "\\", -1)

	err := json.Unmarshal([]byte(temp), &a)

	var msg string
	if err == nil {
		msg = fmt.Sprint(a)
	} else {
		msg = temp
	}

	logger.Println(msg)
}

func PrintToWithColor(msg string, attr color.Attribute) {
	output := color.Output

	if attr == -1 {
		fmt.Fprint(output, msg+"\n")
	} else {
		color.New(attr).Fprintf(output, msg+"\n")
	}
}
