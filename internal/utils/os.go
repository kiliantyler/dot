package utils

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

type OS int

const (
	linux OS = iota
	macOS
)

// String method to get the string representation of the enum
func (d OS) String() string {
	return [...]string{"linux", "macOS"}[d]
}

func getOS() OS {
	funcLogger := log.With().Str("func", "getOS").Logger()
	funcLogger.Trace().Msg("Getting OS")
	switch runtime.GOOS {
	case "linux":
		funcLogger.Trace().Msg("Homebrew on Linux")
		return linux
		// Set a variable or perform Linux-specific initialization here
	case "darwin":
		funcLogger.Trace().Msg("Homebrew on macOS")
		return macOS
	}
	funcLogger.Fatal().Msg("Homebrew is only supported on macOS and Linux.")
	return -1
}
