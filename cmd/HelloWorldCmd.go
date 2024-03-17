package cmd

import (
	"context"

	"github.com/kiliantyler/dot/internal/registry"
	"github.com/kiliantyler/dot/internal/task"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	funcLogger := log.With().Str("func", "init").Logger()
	funcLogger.Trace().Msg("Initializing HelloWorldCmd")
	type contextKey string

	const executorKey contextKey = "echo"
	exec := task.NewGenericExecutor("echo")
	funcLogger.Trace().Msgf("Executor: %s", exec.CommandName)
	cmd := &cobra.Command{
		Use:   "TestLogging",
		Short: "Print 'Hello, World!' to the console and log it at different levels",
		RunE: func(cmd *cobra.Command, args []string) error {
			funcLogger := log.With().Str("func", "RunE").Logger()
			funcLogger.Trace().Msg("Running HelloWorldCmd")
			return run()
		},
	}
	funcLogger.Trace().Msgf("Command: %s", cmd.Name())
	ctx := context.WithValue(context.Background(), executorKey, exec)
	funcLogger.Trace().Msgf("Context: %v", ctx)
	cmd.SetContext(ctx)
	funcLogger.Trace().Msg("Registering command")
	registry.RegisterCommand("", cmd)
}

func run() error {
	log.Trace().Msg("Hello, World!")
	log.Debug().Msg("Hello, World!")
	log.Info().Msg("Hello, World!")
	log.Warn().Msg("Hello, World!")
	log.Error().Msg("Hello, World!")
	return nil
}
