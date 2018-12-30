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
	if logMain.disableInfo {
		return
	}
	logInfo := logMain.returnLog(logMain.outToConsole, logMain.outToFile, info)
	if logInfo == nil {
		return
	}

	err := logInfo.Output(depth, fmt.Sprintln(i...))
	if err != nil {
		fmt.Println("ERROR: while writing main logger in InfoDepth. Error:", err)
	}
}

func Info(i ...interface{}) {
	InfoDepth(3, i...)
}

func WarnDepth(depth int, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	if logMain.disableWarn {
		return
	}
	warnInfo := logMain.returnLog(logMain.outToConsole, logMain.outToFile, warn)
	if warnInfo == nil {
		return
	}

	err := warnInfo.Output(depth, fmt.Sprintln(i...))
	if err != nil {
		fmt.Println("ERROR: while writing main logger in WarnDepth. Error:", err)
	}
}

func Warn(i ...interface{}) {
	WarnDepth(3, i...)
}

func ErrorDepth(depth int, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	logErr := logMain.returnLogError(logMain.outToConsole, logMain.outToFile)
	if logErr == nil {
		return
	}

	err := logErr.Output(depth, fmt.Sprintln(i...))
	if err != nil {
		fmt.Println("ERROR: while writing main logger in ErrorDepth. Error:", err)
	}
}

func Error(i ...interface{}) {
	ErrorDepth(3, i...)
}

func InfoDepthf(depth int, format string, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	if logMain.disableInfo {
		return
	}
	logInfo := logMain.returnLog(logMain.outToConsole, logMain.outToFile, info)
	if logInfo == nil {
		return
	}
	err := logInfo.Output(depth, fmt.Sprintf(format, i...))
	if err != nil {
		fmt.Println("ERROR: while writing main logger in InfoDepthf", err)
	}
}

func Infof(format string, i ...interface{}) {
	InfoDepthf(3, format, i...)
}

func WarnDepthf(depth int, format string, i ...interface{}) {
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	if logMain.disableWarn {
		return
	}
	logWarn := logMain.returnLog(logMain.outToConsole, logMain.outToFile, warn)
	if logWarn == nil {
		return
	}
	err := logWarn.Output(depth, fmt.Sprintf(format, i...))
	if err != nil {
		fmt.Println("ERROR: while writing main logger in WarnDepthf", err)
	}
}

func Warnf(format string, i ...interface{}) {
	WarnDepthf(3, format, i...)
}

func ErrorDepthf(depth int, format string, i ...interface{}) {
	logError := logMain.returnLogError(logMain.outToConsole, logMain.outToFile)
	if logMain == nil {
		fmt.Println(noMainInit)
		return
	}
	if logError == nil {
		return
	}
	err := logError.Output(depth, fmt.Sprintf(format, i...))
	if err != nil {
		fmt.Println("ERROR: while writing main logger in ErrorDepthf", err)
	}
}

func Errorf(format string, i ...interface{}) {
	ErrorDepthf(3, format, i...)
}
