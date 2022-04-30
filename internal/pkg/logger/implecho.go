package logger

import (
	"fmt"
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

func (l *logger) SetHeader(h string) {
	l.WarnMsg("SetHeader is not available.")
	return
}

func (l *logger) Printj(j log.JSON) { l.InfoMap(j) }
func (l *logger) Debugj(j log.JSON) { l.DebugMap(j) }
func (l *logger) Infoj(j log.JSON)  { l.InfoMap(j) }
func (l *logger) Warnj(j log.JSON)  { l.WarnMap(j) }
func (l *logger) Errorj(j log.JSON) { l.ErrorMap(j) }
func (l *logger) Fatalj(j log.JSON) { l.FatalMap(j) }
func (l *logger) Panicj(j log.JSON) { l.PanicMap(j) }

func (l *logger) Printf(format string, args ...interface{}) { l.InfoMsg(fmt.Sprintf(format, args...)) }
func (l *logger) Debugf(format string, args ...interface{}) { l.DebugMsg(fmt.Sprintf(format, args...)) }
func (l *logger) Infof(format string, args ...interface{})  { l.InfoMsg(fmt.Sprintf(format, args...)) }
func (l *logger) Warnf(format string, args ...interface{})  { l.WarnMsg(fmt.Sprintf(format, args...)) }
func (l *logger) Errorf(format string, args ...interface{}) { l.ErrorMsg(fmt.Sprintf(format, args...)) }
func (l *logger) Fatalf(format string, args ...interface{}) { l.FatalMsg(fmt.Sprintf(format, args...)) }
func (l *logger) Panicf(format string, args ...interface{}) { l.PanicMsg(fmt.Sprintf(format, args...)) }

func (l *logger) Print(i ...interface{}) { l.InfoMsg(fmt.Sprint(i...)) }
func (l *logger) Debug(i ...interface{}) { l.DebugMsg(fmt.Sprint(i...)) }
func (l *logger) Info(i ...interface{})  { l.InfoMsg(fmt.Sprint(i...)) }
func (l *logger) Warn(i ...interface{})  { l.WarnMsg(fmt.Sprint(i...)) }
func (l *logger) Error(i ...interface{}) { l.ErrorMsg(fmt.Sprint(i...)) }
func (l *logger) Fatal(i ...interface{}) { l.FatalMsg(fmt.Sprint(i...)) }
func (l *logger) Panic(i ...interface{}) { l.PanicMsg(fmt.Sprint(i...)) }
