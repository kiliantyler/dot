package stow

import (
	"context"

	"github.com/kiliantyler/dot/internal/registry"
	"github.com/kiliantyler/dot/internal/task"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

type contextKey string

const CommandName = "stow"
const ExecutorKey contextKey = CommandName

var exec = task.NewGenericExecutor(string(ExecutorKey))

func stowCommand() *cobra.Command {
	funcLogger := log.With().Str("func", "stowCommand").Logger()
	funcLogger.Trace().Msg("Creating stow command")

	cmd := &cobra.Command{
		Use:   "stow",
		Short: "Automatically stow and unstow dotfiles",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			funcLogger := log.With().Str("func", "RunE").Logger()
			// Extract the executor from the command's context
			exec, ok := cmd.Context().Value(ExecutorKey).(*task.GenericExecutor)
			if !ok {
				funcLogger.Fatal().Msg("Failed to get executor from context")
				return nil
			}

			funcLogger.Trace().Msgf("Executor: %s", exec.CommandName)
			// Dynamically use the command's Use field as the subcommand
			return exec.Execute("", args)
		},
	}
	funcLogger.Trace().Msgf("Command: %s", cmd.Name())
	ctx := context.WithValue(context.Background(), ExecutorKey, exec)
	funcLogger.Trace().Msgf("Context: %v", ctx)
	cmd.SetContext(ctx)
	return cmd
}

func init() {
	funcLogger := log.With().Str("func", "init").Logger()
	funcLogger.Trace().Msg("Initializing stow command")
	funcLogger.Trace().Msgf("Registering command: %s", CommandName)
	registry.RegisterCommand("", stowCommand())
}
