package brew

import (
	"context"

	"github.com/kiliantyler/dot/internal/registry"
	"github.com/kiliantyler/dot/internal/task"
	"github.com/rs/zerolog/log"

	"github.com/spf13/cobra"
)

type contextKey string

const CommandName = "brew"
const ExecutorKey contextKey = CommandName

var exec = task.NewGenericExecutor(string(ExecutorKey))

func brewCommand() *cobra.Command {
	funcLogger := log.With().Str("func", "brewCommand").Logger()
	funcLogger.Trace().Msg("Creating brew command")

	cmd := &cobra.Command{
		Use:   "brew",
		Short: "A front for brew which also updates your installs.yaml file",
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
	funcLogger.Trace().Msg("Initializing brew command")
	funcLogger.Trace().Msgf("Registering command: %s", CommandName)
	registry.RegisterCommand("", brewCommand())
}
