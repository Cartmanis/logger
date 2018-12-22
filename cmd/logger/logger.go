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
	OutToFile    bool
	log          *log.Logger
	file         *os.File
	outToConsole bool
}

func NewLogger(path string, outToConsole bool) (*Logger, error) {
	l := Logger{OutToFile: true}
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	l.file = logFile
	l.outToConsole = outToConsole

	if err != nil {
		return nil, err
	}
	return &l, nil
}

func NewMainLogger(path string, outToConsole bool) error {
	logMain = &Logger{OutToFile: true}
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	logMain.file = logFile
	logMain.outToConsole = outToConsole
	if err != nil {
		return err
	}
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

func (l *Logger) LogInfo(format string, v ...interface{}) {
	logInfo := l.returnLog(l.outToConsole, l.OutToFile, info)
	err := logInfo.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}
func (l *Logger) LogWarn(format string, v ...interface{}) {
	logWarn := l.returnLog(l.outToConsole, l.OutToFile, warn)
	err := logWarn.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}

func (l *Logger) LogError(format string, v ...interface{}) {
	logError := l.returnLogError(l.outToConsole, l.OutToFile)
	err := logError.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}

func LogInfo(format string, v ...interface{}) {
	logInfo := logMain.returnLog(logMain.outToConsole, logMain.OutToFile, info)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the MainLogger function")
		return
	}
	err := logInfo.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogInfo")
	}
}

func LogWarn(format string, v ...interface{}) {
	logWarn := logMain.returnLog(logMain.outToConsole, logMain.OutToFile, warn)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the MainLogger function")
		return
	}
	err := logWarn.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogWarn основого лога")
	}
}

func LogError(format string, v ...interface{}) {
	logError := logMain.returnLogError(logMain.outToConsole, logMain.OutToFile)
	if logMain == nil {
		fmt.Println("ERROR : The main stream logger has not been initialized. Please call the MainLogger function")
		return
	}
	err := logError.Output(2, fmt.Sprintf(format, v...))
	if err != nil {
		fmt.Printf("ERROR: при записи LogWarn основого лога")
	}
}

func (l *Logger) returnLog(outToConsole, outToFile bool, level string) *log.Logger {
	if l.outToConsole && l.OutToFile {
		l.log = log.New(io.MultiWriter(l.file, os.Stdout), level, flags)
		return l.log
	}
	if !l.OutToFile {
		l.log = log.New(io.MultiWriter(os.Stdout), level, flags)
		return l.log
	}
	if l.OutToFile && !l.outToConsole {
		l.log = log.New(io.MultiWriter(l.file), level, flags)
		return l.log
	}
	return l.log
}

func (l *Logger) returnLogError(outToConsole, outToFile bool) *log.Logger {
	if l.OutToFile {
		l.log = log.New(io.MultiWriter(l.file, os.Stderr), errLog, flags)
	}
	if !l.OutToFile {
		l.log = log.New(io.MultiWriter(os.Stdout), errLog, flags)
	}
	return l.log
}
