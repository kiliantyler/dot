package cli

import (
	"os"

	"github.com/rs/zerolog/log"
)

func isTerminal() bool {
	funcLogger := log.With().Str("func", "isTerminal").Logger()
	funcLogger.Trace().Msg("Getting terminal status")
	fileInfo, _ := os.Stdout.Stat()
	funcLogger.Trace().Msgf("File mode: %s, ModeCharDevice: %s", fileInfo.Mode(), os.ModeCharDevice)
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}
