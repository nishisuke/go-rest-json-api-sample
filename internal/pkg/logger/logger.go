package logger

import (
	"encoding/json"
	"io"

	gommonLog "github.com/labstack/gommon/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Prepare(out io.Writer, lv gommonLog.Lvl) *logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &logger{out: out, lv: lv}
}

func (l *logger) DebugErr(err error) {
	log.Debug().Err(err).Send()
}
func (l *logger) InfoErr(err error) {
	log.Info().Err(err).Send()
}
func (l *logger) WarnErr(err error) {
	log.Warn().Err(err).Send()
}
func (l *logger) ErrorErr(err error) {
	log.Error().Err(err).Send()
}
func (l *logger) FatalErr(err error) {
	log.Fatal().Err(err).Send()
}
func (l *logger) PanicErr(err error) {
	log.Panic().Err(err).Send()
}

func (l *logger) DebugMsg(str string) {
	log.Debug().Msg(str)
}
func (l *logger) InfoMsg(str string) {
	log.Info().Msg(str)
}
func (l *logger) WarnMsg(str string) {
	log.Warn().Msg(str)
}
func (l *logger) ErrorMsg(str string) {
	log.Error().Msg(str)
}
func (l *logger) FatalMsg(str string) {
	log.Fatal().Msg(str)
}
func (l *logger) PanicMsg(str string) {
	log.Panic().Msg(str)
}

func (l *logger) DebugMap(b map[string]interface{}) {
	log.Panic().RawJSON("data", mapToBytes(b))
}
func (l *logger) InfoMap(b map[string]interface{}) {
	log.Panic().RawJSON("data", mapToBytes(b))
}
func (l *logger) WarnMap(b map[string]interface{}) {
	log.Panic().RawJSON("data", mapToBytes(b))
}
func (l *logger) ErrorMap(b map[string]interface{}) {
	log.Panic().RawJSON("data", mapToBytes(b))
}
func (l *logger) FatalMap(b map[string]interface{}) {
	log.Panic().RawJSON("data", mapToBytes(b))
}
func (l *logger) PanicMap(b map[string]interface{}) {
	log.Panic().RawJSON("data", mapToBytes(b))
}

func mapToBytes(m map[string]interface{}) []byte {
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}
