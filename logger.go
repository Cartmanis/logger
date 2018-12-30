package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

const flags = log.LstdFlags | log.Lshortfile

const (
	info   = "INFO : "
	warn   = "WARN : "
	errLog = "ERROR : "
)

const noInit = "The stream logger has not been initialized. Please call the NewLogger() method"

type Logger struct {
	outToFile    bool
	disableWarn  bool
	disableInfo  bool
	log          *log.Logger
	file         *os.File
	outToConsole bool
}

func NewLogger(path string, outToConsole bool, outToFile bool) (*Logger, error) {
	if outToFile && path == "" {
		return nil, errors.New("To record the logger in the file, You must specify the path to the file")
	}
	l := Logger{}
	if outToFile {
		logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
		l.file = logFile
		if err != nil {
			return nil, err
		}
	}
	l.outToConsole = outToConsole
	l.outToFile = outToFile
	return &l, nil
}

func (l *Logger) Close() error {
	if l != nil && l.file != nil {
		return l.file.Close()
	}
	return errors.New(noInit)
}

func (l *Logger) DisableWarn() {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	l.disableWarn = true
}

func (l *Logger) DisableInfo() {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	l.disableInfo = true
}

func (l *Logger) InfoDepth(depth int, i ...interface{}) {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	if l.disableInfo {
		return
	}
	logInfo := l.returnLog(l.outToConsole, l.outToFile, info)
	if logInfo == nil {
		return
	}

	err := logInfo.Output(depth, fmt.Sprintln(i...))
	if err != nil {
		fmt.Println("ERROR: while writing in InfoDepth. Error:", err)
	}
}

func (l *Logger) Info(i ...interface{}) {
	l.InfoDepth(3, i...)
}

func (l *Logger) WarnDepth(depth int, i ...interface{}) {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	if l.disableWarn {
		return
	}
	warnInfo := l.returnLog(l.outToConsole, l.outToFile, warn)
	if warnInfo == nil {
		return
	}

	err := warnInfo.Output(depth, fmt.Sprintln(i...))
	if err != nil {
		fmt.Println("ERROR: while writing in WarnDepth. Error:", err)
	}
}

func (l *Logger) Warn(i ...interface{}) {
	l.WarnDepth(3, i...)
}

func (l *Logger) ErrorDepth(depth int, i ...interface{}) {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	logErr := l.returnLogError(l.outToConsole, l.outToFile)
	if logErr == nil {
		return
	}

	err := logErr.Output(depth, fmt.Sprintln(i...))
	if err != nil {
		fmt.Println("ERROR: while writing in ErrorDepth. Error:", err)
	}
}

func (l *Logger) Error(i ...interface{}) {
	l.ErrorDepth(3, i...)
}

func (l *Logger) InfoDepthf(depth int, format string, i ...interface{}) {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	if l.disableInfo {
		return
	}
	logInfo := l.returnLog(l.outToConsole, l.outToFile, info)
	if logInfo == nil {
		return
	}
	err := logInfo.Output(depth, fmt.Sprintf(format, i...))
	if err != nil {
		fmt.Println("ERROR: while writing in InfoDepthf. Error:", err)
	}
}

func (l *Logger) Infof(format string, i ...interface{}) {
	l.InfoDepthf(3, format, i...)
}

func (l *Logger) WarnDepthf(depth int, format string, i ...interface{}) {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	if l.disableWarn {
		return
	}
	logWarn := l.returnLog(l.outToConsole, l.outToFile, warn)
	if logWarn == nil {
		return
	}
	err := logWarn.Output(depth, fmt.Sprintf(format, i...))
	if err != nil {
		fmt.Println("ERROR: while writing in WarnDepthf. Error:", err)
	}
}

func (l *Logger) Warnf(format string, i ...interface{}) {
	l.WarnDepthf(3, format, i...)
}

func (l *Logger) ErrorDepthf(depth int, format string, i ...interface{}) {
	if l == nil {
		fmt.Println(noInit)
		return
	}
	logError := l.returnLogError(l.outToConsole, l.outToFile)
	if logError == nil {
		return
	}
	err := logError.Output(depth, fmt.Sprintf(format, i...))
	if err != nil {
		fmt.Println("ERROR: while writing in ErrorDepthf. Error:", err)
	}
}

func (l *Logger) Errorf(format string, i ...interface{}) {
	l.ErrorDepthf(3, format, i...)
}

func (l *Logger) returnLog(outToConsole, outToFile bool, level string) *log.Logger {
	if l.outToConsole && l.outToFile && l.file != nil {
		l.log = log.New(io.MultiWriter(l.file, os.Stdout), level, flags)
		return l.log
	}
	if !l.outToFile && l.outToConsole {
		l.log = log.New(io.MultiWriter(os.Stdout), level, flags)
		return l.log
	}
	if l.outToFile && !l.outToConsole && l.file != nil {
		l.log = log.New(io.MultiWriter(l.file), level, flags)
		return l.log
	}
	return l.log
}

func (l *Logger) returnLogError(outToConsole, outToFile bool) *log.Logger {
	if l.outToFile && l.file != nil {
		l.log = log.New(io.MultiWriter(l.file, os.Stderr), errLog, flags)
	}
	if !l.outToFile {
		l.log = log.New(io.MultiWriter(os.Stderr), errLog, flags)
	}
	return l.log
}
