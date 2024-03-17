package utils

import (
	"runtime"

	"github.com/rs/zerolog/log"
)

type Arch int

const (
	amd64 Arch = iota
	arm64
)

// String method to get the string representation of the enum
func (a Arch) String() string {
	return [...]string{"amd64", "arm64"}[a]
}

func getArch() Arch {
	funcLogger := log.With().Str("func", "getArch").Logger()
	funcLogger.Trace().Msg("Getting architecture")
	switch runtime.GOARCH {
	case "amd64":
		funcLogger.Trace().Msg("Architecture: amd64")
		return amd64
	case "arm64":
		funcLogger.Trace().Msg("Architecture: arm64")
		return arm64
	}
	funcLogger.Fatal().Msg("Unsupported architecture.")
	return -1
}
