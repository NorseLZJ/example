package log

import (
	"github.com/astaxie/beego/logs"
)

const (
	Console = iota + 1 // os.stdout
	File
)

type Log struct {
	log *logs.BeeLogger
}

var log *Log

func init() {
	log = &Log{}
	log.log = logs.NewLogger(10000)
	log.log.SetLogger("console", "")
}

func LLog() *Log {
	return log
}

func (l *Log) InitLog(output int, logName string) *Log {
	switch output {
	case Console:
		log.log.SetLogger("console", "")
	case File:
		log.log.SetLogger("file", logName)
	default:
		log.log.SetLogger("console", "")
	}
	return log
}

func (l *Log) Close(param interface{}) {
	log.log.Close()
}

func (l *Log) Trace(format string, a ...interface{}) {
	log.log.Trace(format, a...)
}
func (l *Log) Warn(format string, a ...interface{}) {
	log.log.Warn(format, a...)
}
func (l *Log) Debug(format string, a ...interface{}) {
	log.log.Debug(format, a...)
}
func (l *Log) Info(format string, a ...interface{}) {
	log.log.Info(format, a...)
}
func (l *Log) Critical(format string, a ...interface{}) {
	log.log.Critical(format, a...)
}
func (l *Log) Error(format string, a ...interface{}) {
	log.log.Error(format, a...)
}
