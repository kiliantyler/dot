package task

import (
	"context"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/kiliantyler/dot/internal/utils"
	"github.com/rs/zerolog/log"
)

// Execute runs the specified command with the provided arguments, combining them with preset args.
func (e *GenericExecutor) Execute(subcommand string, runtimeArgs []string) error {
	funcLogger := log.With().Str("func", "Execute").Logger()
	if !utils.IsCommandAvailable("brew") {
		funcLogger.Error().Msgf("The required command '%s' is not installed or not in PATH.", e.CommandName)
		if e.CommandName != "brew" {
			funcLogger.Fatal().Msgf("You can install it with 'brew install %s'", e.CommandName)
			// TODO: Handle brew install here
		}
		return nil
	}
	var args []string
	if subcommand != "" {
		args = append(args, subcommand)
	}
	args = append(args, runtimeArgs...)

	funcLogger.Trace().Msgf("Executing '%s' with args: %v\n", e.CommandName, args)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cmd := exec.CommandContext(ctx, e.CommandName, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Setting up signal capturing
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Start command execution in a goroutine
	if err := cmd.Start(); err != nil {
		funcLogger.Error().Msgf("Error starting command: %s", err)
		return err
	}

	// // Listen for the interrupt signal in a separate goroutine
	go func() {
		<-sigChan
		funcLogger.Fatal().Msgf("Received an interrupt, stopping tasks...")
		cancel() // cancels the context
	}()

	// Wait for the command to finish
	return cmd.Wait()
}
