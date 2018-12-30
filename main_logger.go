package logger

import (
	"errors"
	"fmt"
)

var logMain *Logger

const noMainInit = "The stream main logger has not been initialized. Please call NewMainLogger() function"

func NewMainLogger(path string, outToConsole bool, outToFile bool) error {
	l, err := NewLogger(path, outToConsole, outToFile)
	if err != nil {
		return err
	}
	logMain = l
	return nil
}

func Close() error {
	if logMain != nil && logMain.file != nil {
		return logMain.file.Close()
	}
	return errors.New(noMainInit)
}

func DisableWarn() {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.disableWarn = true
}

func DisableInfo() {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.disableInfo = true
}

func InfoDepth(depth int, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.InfoDepth(depth, i...)
}

func Info(i ...interface{}) {
	InfoDepth(4, i...)
}

func WarnDepth(depth int, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.WarnDepth(depth, i...)
}

func Warn(i ...interface{}) {
	WarnDepth(4, i...)
}

func ErrorDepth(depth int, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.ErrorDepth(depth, i...)
}

func Error(i ...interface{}) {
	ErrorDepth(4, i...)
}

func InfoDepthf(depth int, format string, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.InfoDepthf(depth, format, i...)
}

func Infof(format string, i ...interface{}) {
	InfoDepthf(4, format, i...)
}

func WarnDepthf(depth int, format string, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.WarnDepthf(depth, format, i...)
}

func Warnf(format string, i ...interface{}) {
	WarnDepthf(4, format, i...)
}

func ErrorDepthf(depth int, format string, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logMain.ErrorDepthf(depth, format, i...)
}

func Errorf(format string, i ...interface{}) {
	ErrorDepthf(4, format, i...)
}
