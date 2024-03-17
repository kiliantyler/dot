package registry

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command
var (
	// Initialize the command registry. The map's key is the parent command's use string.
	// If a command is intended to be a root command, use an empty string as its key.
	commandRegistry = make(map[string][]*cobra.Command)
)

func SetRootCommand(cmd *cobra.Command) {
	funcLogger := log.With().Str("func", "SetRootCommand").Logger()
	funcLogger.Trace().Msg("Setting root command")
	rootCmd = cmd
}

func RegisterCommand(parentName string, cmd *cobra.Command) {
	funcLogger := log.With().Str("func", "RegisterCommand").Logger()
	if parentName == "" {
		funcLogger.Trace().Msgf("Registering root command: %s", cmd.Use)
	} else {
		funcLogger.Trace().Msgf("Registering command: %s under parent: %s", cmd.Use, parentName)
	}
	commandRegistry[parentName] = append(commandRegistry[parentName], cmd)
}

func AttachCommands() {
	funcLogger := log.With().Str("func", "AttachCommands").Logger()
	funcLogger.Trace().Msg("Attaching commands")
	for parentName, cmds := range commandRegistry {
		if parentName == "" { // Direct children of the root command.
			for _, cmd := range cmds {
				funcLogger.Trace().Msgf("Adding command: %s to root command", cmd.Use)
				rootCmd.AddCommand(cmd)
			}
		} else { // Subcommands.
			for _, parentCmd := range rootCmd.Commands() {
				if parentCmd.Use == parentName {
					for _, cmd := range cmds {
						funcLogger.Trace().Msgf("Adding command: %s to parent: %s", cmd.Use, parentName)
						parentCmd.AddCommand(cmd)
					}
					break
				}
			}
		}
	}
}
