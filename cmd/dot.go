package cmd

import (
	"strconv"

	_ "github.com/kiliantyler/dot/cmd/brew"
	_ "github.com/kiliantyler/dot/cmd/stow"
	"github.com/kiliantyler/dot/internal/config"
	"github.com/kiliantyler/dot/internal/registry"
	"github.com/kiliantyler/dot/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dot",
	Short: "dot is used to manage your dotfiles and some other things",
	Long: `dot is a CLI tool to manage your dotfiles and some other things.
It is designed to be simple and easy to use.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		funcLogger := log.With().Str("func", "PersistentPreRun").Logger()
		funcLogger.Trace().Msg("Running PersistentPreRun on Root Command")
		config.SetupConfig(cmd)
	},
	SilenceErrors: true,
}

func Execute() {
	funcLogger := log.With().Str("func", "Execute").Logger()
	funcLogger.Trace().Msg("Executing root command")
	config.InitConfig(rootCmd)
	registry.SetRootCommand(rootCmd)
	registry.AttachCommands()
	funcLogger.Trace().Msgf("%s", strconv.FormatBool(utils.FileNotOwned("/Users/kilian/dotfiles")))

	if err := rootCmd.Execute(); err != nil {
		funcLogger.Fatal().Msg("Error: " + err.Error())
	}
}

func AddCommand(cmd *cobra.Command) {
	funcLogger := log.With().Str("func", "AddCommand").Logger()
	funcLogger.Trace().Msgf("Adding command: %s", cmd.Name())
	rootCmd.AddCommand(cmd)
}
