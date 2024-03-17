package cli

import (
	"os/exec"

	"github.com/rs/zerolog/log"
)

func cloneRepo(repoURL, destDir string) error {
	funcLogger := log.With().Str("func", "cloneRepo").Logger()
	// TODO: Move this to an executor
	funcLogger.Debug().Msgf("Cloning %s to %s", repoURL, destDir)
	cmd := exec.Command("git", "clone", repoURL, destDir)
	return cmd.Run()
}
