package logger

import (
	"fmt"
	"io"

	"github.com/labstack/gommon/log"
)

func (l *Logger) Output() io.Writer {
	return l.out
}

func (l *Logger) SetOutput(w io.Writer) {
	l.out = w
}

func (l *Logger) Prefix() string {
	return l.prefix
}

func (l *Logger) SetPrefix(p string) {
	l.prefix = p
}

func (l *Logger) Level() log.Lvl {
	return l.lv
}

func (l *Logger) SetLevel(v log.Lvl) {
	l.lv = v
}

func (l *Logger) SetHeader(h string) {
	l.WarnMsg("SetHeader is not available.")
	return
}

func (l *Logger) Printj(j log.JSON) { l.InfoMap(j) }
func (l *Logger) Debugj(j log.JSON) { l.DebugMap(j) }
func (l *Logger) Infoj(j log.JSON)  { l.InfoMap(j) }
func (l *Logger) Warnj(j log.JSON)  { l.WarnMap(j) }
func (l *Logger) Errorj(j log.JSON) { l.ErrorMap(j) }
func (l *Logger) Fatalj(j log.JSON) { l.FatalMap(j) }
func (l *Logger) Panicj(j log.JSON) { l.PanicMap(j) }

func (l *Logger) Printf(format string, args ...interface{}) {
	l.InfoMsg(fmt.Sprintf(format, args...))
}
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.DebugMsg(fmt.Sprintf(format, args...))
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.InfoMsg(fmt.Sprintf(format, args...))
}
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.WarnMsg(fmt.Sprintf(format, args...))
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.ErrorMsg(fmt.Sprintf(format, args...))
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.FatalMsg(fmt.Sprintf(format, args...))
}
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.PanicMsg(fmt.Sprintf(format, args...))
}

func (l *Logger) Print(i ...interface{}) { l.InfoMsg(fmt.Sprint(i...)) }
func (l *Logger) Debug(i ...interface{}) { l.DebugMsg(fmt.Sprint(i...)) }
func (l *Logger) Info(i ...interface{})  { l.InfoMsg(fmt.Sprint(i...)) }
func (l *Logger) Warn(i ...interface{})  { l.WarnMsg(fmt.Sprint(i...)) }
func (l *Logger) Error(i ...interface{}) { l.ErrorMsg(fmt.Sprint(i...)) }
func (l *Logger) Fatal(i ...interface{}) { l.FatalMsg(fmt.Sprint(i...)) }
func (l *Logger) Panic(i ...interface{}) { l.PanicMsg(fmt.Sprint(i...)) }
