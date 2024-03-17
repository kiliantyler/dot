package config

import (
	"github.com/kiliantyler/dot/internal/logger"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbosity string

func InitConfig(rootCmd *cobra.Command) {
	funcLogger := log.With().Str("func", "InitConfig").Logger()
	funcLogger.Trace().Msg("Initializing config")
	rootCmd.PersistentFlags().StringP("verbosity", "v", "error", "Set the logging verbosity (trace, debug, info, warn, error, fatal)")
	funcLogger.Trace().Msgf("Got verbosity flag: %s", rootCmd.PersistentFlags().Lookup("verbosity").Value.String())
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))
	funcLogger.Trace().Msgf("Bound verbosity to viper: %s", viper.GetString("verbosity"))
	viper.AutomaticEnv()
	funcLogger.Trace().Msg("Automatically bound environment variables to viper")
}

func SetupConfig(rootCmd *cobra.Command) {
	funcLogger := log.With().Str("func", "SetupConfig").Logger()
	funcLogger.Trace().Msg("Setting up config with verbosity: " + viper.GetString("verbosity"))
	logger.SetupLogger(viper.GetString("verbosity"))
}
