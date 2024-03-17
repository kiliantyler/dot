package utils

import (
	"os/exec"

	"github.com/rs/zerolog/log"
)

// IsCommandAvailable checks if the given command is available in the system's PATH.
func IsCommandAvailable(name string) bool {
	funcLogger := log.With().Str("func", "IsCommandAvailable").Logger()
	funcLogger.Trace().Msgf("Checking if command '%s' is available", name)
	_, err := exec.LookPath(name)
	if err != nil {
		funcLogger.Trace().Msgf("Command '%s' is not available", name)
	}
	return err == nil
}
