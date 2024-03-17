package brew

import (
	"context"

	"github.com/kiliantyler/dot/internal/registry"
	"github.com/kiliantyler/dot/internal/task"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	funcLogger := log.With().Str("func", "init").Logger()
	funcLogger.Trace().Msg("Initializing brew command")
	// Example subcommand setup
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List something",
		RunE: func(cmd *cobra.Command, args []string) error {
			funcLogger := log.With().Str("func", "RunE").Logger()
			// Extract the executor from the command's context
			exec, ok := cmd.Context().Value(ExecutorKey).(*task.GenericExecutor)
			if !ok {
				funcLogger.Fatal().Msg("Failed to get executor from context")
				// Handle error: Executor not found in context
				return nil
			}

			funcLogger.Trace().Msgf("Executor: %s", exec.CommandName)
			// Dynamically use the command's Use field as the subcommand
			return exec.Execute(cmd.Use, args)
		},
	}
	funcLogger.Trace().Msgf("Command: %s", cmd.Name())
	ctx := context.WithValue(context.Background(), ExecutorKey, exec)
	funcLogger.Trace().Msgf("Context: %v", ctx)
	cmd.SetContext(ctx)
	funcLogger.Trace().Msgf("Registering command: %s", CommandName)
	registry.RegisterCommand(CommandName, cmd)
}
