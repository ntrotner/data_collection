package logger

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var (
	Logger *LogComponent
)

type LogComponent struct {
	loggerError *log.Logger
	loggerDebug *log.Logger
	loggerInfo  *log.Logger
	buffer      *bytes.Buffer
}

func (l LogComponent) LogDebug(msg string) {
	l.loggerDebug.Println(msg)
	if l.buffer != nil {
		l.bufferRead()
	}
}

func (l LogComponent) LogError(msg string) {
	l.loggerError.Println(msg)
	if l.buffer != nil {
		l.bufferRead()
	}
}

func (l LogComponent) LogInfo(msg string) {
	l.loggerInfo.Println(msg)
	if l.buffer != nil {
		l.bufferRead()
	}
}

func (l LogComponent) bufferRead() {
	fmt.Printf(l.buffer.String())
	l.buffer.Reset()
}

func (l LogComponent) IsLoggingToFile() bool {
	return l.buffer == nil
}

func (l LogComponent) IsLoggingToFileDesc() string {
	if l.IsLoggingToFile() {
		return "File"
	}
	return "Terminal"
}

func SetupLogger() *LogComponent {
	var buf bytes.Buffer

	loggerComponent := LogComponent{
		loggerError: log.New(&buf, "[ERROR] ", log.Ldate|log.Ltime),
		loggerDebug: log.New(&buf, "[DEBUG] ", log.Ldate|log.Ltime),
		loggerInfo:  log.New(&buf, "[INFO] ", log.Ldate|log.Ltime),
		buffer:      &buf,
	}

	if os.Getenv("LOGFILE") != "" {
		file, err := os.OpenFile(os.Getenv("LOGFILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			loggerComponent.LogError(err.Error())
		} else {
			loggerComponent.loggerDebug.SetOutput(file)
			loggerComponent.loggerInfo.SetOutput(file)
			loggerComponent.loggerError.SetOutput(file)
			loggerComponent.buffer = nil
		}
	}

	return &loggerComponent
}

func init() {
	Logger = SetupLogger()
}
