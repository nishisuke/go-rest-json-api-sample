package logger

import (
	"io"

	"github.com/labstack/gommon/log"
)

type logger struct {
	prefix string
	lv     log.Lvl
	out    io.Writer
}

func (l *logger) Output() io.Writer {
	return l.out

}
func (l *logger) SetOutput(w io.Writer) {
	l.out = w
}
func (l *logger) Prefix() string {
	return l.prefix

}
func (l *logger) SetPrefix(p string) {
	l.prefix = p

}
func (l *logger) Level() log.Lvl {
	return l.lv

}
func (l *logger) SetLevel(v log.Lvl) {
	l.lv = v
}
func (l *logger) SetHeader(h string)                        {}
func (l *logger) Print(i ...interface{})                    {}
func (l *logger) Printf(format string, args ...interface{}) {}
func (l *logger) Printj(j log.JSON)                         {}
func (l *logger) Debug(i ...interface{})                    {}
func (l *logger) Debugf(format string, args ...interface{}) {}
func (l *logger) Debugj(j log.JSON)                         {}
func (l *logger) Info(i ...interface{})                     {}
func (l *logger) Infof(format string, args ...interface{})  {}
func (l *logger) Infoj(j log.JSON)                          {}
func (l *logger) Warn(i ...interface{})                     {}
func (l *logger) Warnf(format string, args ...interface{})  {}
func (l *logger) Warnj(j log.JSON)                          {}
func (l *logger) Error(i ...interface{})                    {}
func (l *logger) Errorf(format string, args ...interface{}) {}
func (l *logger) Errorj(j log.JSON)                         {}
func (l *logger) Fatal(i ...interface{})                    {}
func (l *logger) Fatalj(j log.JSON)                         {}
func (l *logger) Fatalf(format string, args ...interface{}) {}
func (l *logger) Panic(i ...interface{})                    {}
func (l *logger) Panicj(j log.JSON)                         {}
func (l *logger) Panicf(format string, args ...interface{}) {}
