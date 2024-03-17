package utils

import (
	"os"

	"github.com/rs/zerolog/log"
)

func CreateDirIfNotExist(dirPath string) error {
	funcLogger := log.With().Str("func", "createDirIfNotExist").Logger()
	funcLogger.Trace().Msgf("Checking if directory exists: %s", dirPath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		funcLogger.Trace().Msgf("Creating directory: %s", dirPath)
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			funcLogger.Error().Msgf("Error creating directory: %s", err)
			return err
		}
		funcLogger.Trace().Msgf("Directory created: %s", dirPath)
	}
	return nil
}
