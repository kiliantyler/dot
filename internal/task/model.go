package task

import "github.com/rs/zerolog/log"

// GenericExecutor executes a specified command with given arguments.
type GenericExecutor struct {
	CommandName string
	Args        []string // Preset arguments, can be empty
}

// NewGenericExecutor creates a new instance of GenericExecutor for a given command.
func NewGenericExecutor(commandName string, args ...string) *GenericExecutor {
	funcLogger := log.With().Str("func", "NewGenericExecutor").Logger()
	funcLogger.Trace().Msgf("Creating new GenericExecutor for command: %s", commandName)
	return &GenericExecutor{
		CommandName: commandName,
		Args:        args,
	}
}
