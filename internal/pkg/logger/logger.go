package logger

import (
	"encoding/json"
	"io"

	gommonLog "github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Logger struct {
	prefix string
	lv     gommonLog.Lvl
	out    io.Writer
}

func Prepare(out io.Writer, lv gommonLog.Lvl) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &Logger{out: out, lv: lv}
}

func (l *Logger) DebugErr(err error) {
	log.Debug().Err(err).Send()
}
func (l *Logger) InfoErr(err error) {
	log.Info().Err(err).Send()
}
func (l *Logger) WarnErr(err error) {
	log.Warn().Err(err).Send()
}
func (l *Logger) ErrorErr(err error) {
	log.Error().Err(err).Send()
}
func (l *Logger) FatalErr(err error) {
	log.Fatal().Err(err).Send()
}
func (l *Logger) PanicErr(err error) {
	log.Panic().Err(err).Send()
}

func (l *Logger) DebugMsg(str string) {
	log.Debug().Msg(str)
}
func (l *Logger) InfoMsg(str string) {
	log.Info().Msg(str)
}
func (l *Logger) WarnMsg(str string) {
	log.Warn().Msg(str)
}
func (l *Logger) ErrorMsg(str string) {
	log.Error().Msg(str)
}
func (l *Logger) FatalMsg(str string) {
	log.Fatal().Msg(str)
}
func (l *Logger) PanicMsg(str string) {
	log.Panic().Msg(str)
}

func (l *Logger) DebugMap(m map[string]interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		l.ErrorErr(err)
		return
	}
	log.Panic().RawJSON("data", (b))
}
func (l *Logger) InfoMap(m map[string]interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		l.ErrorErr(err)
		return
	}
	log.Panic().RawJSON("data", (b))
}
func (l *Logger) WarnMap(m map[string]interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		l.ErrorErr(err)
		return
	}
	log.Panic().RawJSON("data", (b))
}
func (l *Logger) ErrorMap(m map[string]interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		l.ErrorErr(err)
		return
	}
	log.Panic().RawJSON("data", (b))
}
func (l *Logger) FatalMap(m map[string]interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		l.ErrorErr(err)
		return
	}
	log.Panic().RawJSON("data", (b))
}
func (l *Logger) PanicMap(m map[string]interface{}) {
	b, err := json.Marshal(m)
	if err != nil {
		l.ErrorErr(err)
		return
	}
	log.Panic().RawJSON("data", b)
}
