package config

import (
	"github.com/adrg/xdg"
	"github.com/kiliantyler/dot/internal/logger"
	"github.com/kiliantyler/dot/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Verbosity string

func InitConfig(rootCmd *cobra.Command) {
	funcLogger := log.With().Str("func", "InitConfig").Logger()
	// Set the file name of the configurations file
	rootCmd.PersistentFlags().StringVarP(&Verbosity, "verbosity", "v", "error", "Set the logging verbosity (trace, debug, info, warn, error, fatal)")
	funcLogger.Error().Msgf("Setting verbosity to: %s", rootCmd.PersistentFlags().Lookup("verbosity").Value.String())
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))

	// Use viper to handle configuration file and environment variables
	viper.SetDefault("verbosity", "error")

	confDir := xdg.ConfigHome + "/.dot"
	viper.AddConfigPath(confDir)
	viper.SetConfigName("dot")
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	utils.CreateDirIfNotExist(confDir)
	// Create a default config file if it doesn't exist
	if !utils.FileExists(confDir + "/dot.yaml") {
		funcLogger.Trace().Msg("Creating default config file")
		if err := viper.SafeWriteConfig(); err != nil {
			funcLogger.Fatal().Msgf("Can't write config: %s", err)
		}
	}

	// Read in environment variables that match
	viper.AutomaticEnv()
	// Attempt to read the config file
	if err := viper.ReadInConfig(); err == nil {
		funcLogger.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	} else {
		funcLogger.Warn().Msgf("Can't read config: %s", err)
	}

	// Setup logger as needed
	logger.SetupLogger()
}
