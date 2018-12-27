package logger

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
)

const flags = log.LstdFlags | log.Lshortfile

var logMain *Logger

const (
	info   = "INFO : "
	warn   = "WARN : "
	errLog = "ERROR : "
)

type Logger struct {
	outToFile    bool
	log          *log.Logger
	file         *os.File
	outToConsole bool
}

func NewLogger(path string, outToConsole bool, outToFile bool) (*Logger, error) {
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

func NewMainLogger(path string, outToConsole bool, outToFile bool) error {
	logMain = &Logger{}
	if outToFile && path == "" {
		return errors.New("To record the logger in the file, you must specify the path to the file")
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

func (l *Logger) Close() error {
	if l != nil && l.file != nil {
		return l.file.Close()
	}
	return errors.New("The logger is not initialized. Please call NewLogger()")
}

func Close() error {
	if logMain != nil && logMain.file != nil {
		return logMain.file.Close()
	}
	return errors.New("The main logger is not initialized. Please call MainLogger()")
}

func (l *Logger) Info(format string, v ...interface{}) {
	if l == nil {
		return
	}
	logInfo := l.returnLog(l.outToConsole, l.outToFile, info)
	if logInfo == nil {
		return
	}
	err := logInfo.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}
func (l *Logger) Warn(format string, v ...interface{}) {
	if l == nil {
		return
	}
	logWarn := l.returnLog(l.outToConsole, l.outToFile, warn)
	if logWarn == nil {
		return
	}
	err := logWarn.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}

func (l *Logger) Error(format string, v ...interface{}) {
	if l == nil {
		return
	}
	logError := l.returnLogError(l.outToConsole, l.outToFile)
	if logError == nil {
		return
	}
	err := logError.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}

func Info(format string, v ...interface{}) {
	InfoDepth(2, format, v...)
}

func Warn(format string, v ...interface{}) {
	WarnDepth(2, format, v...)
}

func Error(format string, v ...interface{}) {
	ErrorDepth(2, format, v...)
}

func InfoDepth(depth int, format string, v ...interface{}) {
	logInfo := logMain.returnLog(logMain.outToConsole, logMain.outToFile, info)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
		return
	}
	if logInfo == nil {
		return
	}
	err := logInfo.Output(depth, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo основного лога %v", err)
	}
}

func WarnDepth(depth int, format string, v ...interface{}) {
	logWarn := logMain.returnLog(logMain.outToConsole, logMain.outToFile, warn)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
		return
	}
	if logWarn == nil {
		return
	}
	err := logWarn.Output(depth, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogWarn основого лога", err)
	}
}

func ErrorDepth(depth int, format string, v ...interface{}) {
	logError := logMain.returnLogError(logMain.outToConsole, logMain.outToFile)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the NewMainLogger function")
		return
	}
	if logError == nil {
		return
	}
	err := logError.Output(depth, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogWarn основого лога")
	}
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
