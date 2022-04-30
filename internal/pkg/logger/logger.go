package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func NewLogger() *logger {
	return &logger{
		out: os.Stdout,
	}
}

func Prepare(out io.Writer) *logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return &logger{out: out}
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