package logger

import (
	"errors"
	"fmt"
	"os"
)

var logMain *Logger

func NewMainLogger(path string, outToConsole bool, outToFile bool) error {
	logMain = &Logger{}
	if outToFile && path == "" {
		return errors.New("To record the logger in the file, You must specify the path to the file")
	}
	if outToFile {
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		logMain.file = logFile
		if err != nil {
			return err
		}
	}
	logMain.outToConsole = outToConsole
	logMain.outToFile = outToFile

	return nil
}

func Close() error {
	if logMain != nil && logMain.file != nil {
		return logMain.file.Close()
	}
	return errors.New("The main logger is not initialized. Please call NewMainLogger()")
}

func InfoDepth(depth int, i ...interface{}) {
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
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
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
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
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
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
	logInfo := logMain.returnLog(logMain.outToConsole, logMain.outToFile, info)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
		return
	}
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
	logWarn := logMain.returnLog(logMain.outToConsole, logMain.outToFile, warn)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
		return
	}
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
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
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
