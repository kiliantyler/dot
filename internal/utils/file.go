package utils

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

func GetPermission(path string) (string, error) {
	funcLogger := log.With().Str("func", "GetPermission").Logger()
	funcLogger.Trace().Msg(path)
	cmd := exec.Command("command", "stat", "-u", "%A", path)
	out, err := cmd.Output()
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return "", err
	}
	funcLogger.Trace().Msgf("%s", strings.TrimSpace(string(out)))
	return strings.TrimSpace(string(out)), nil
}

func IsUserOnlyChmod(path string) bool {
	funcLogger := log.With().Str("func", "IsUserOnlyChmod").Logger()
	funcLogger.Trace().Msg(path)
	info, err := os.Stat(path)
	if err != nil || !info.IsDir() {
		funcLogger.Info().Msg(err.Error())
		return false
	}
	permission, err := GetPermission(path)
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return false
	}
	funcLogger.Trace().Msg(permission)
	return !strings.HasPrefix(permission, "75") || permission == "750" || permission == "751" || permission == "754" || permission == "755"
}

func ExistsButNotWritable(path string) bool {
	funcLogger := log.With().Str("func", "ExistsButNotWritable").Logger()
	info, err := os.Stat(path)
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return false
	}
	funcLogger.Trace().Msgf("%s", info.Mode().Perm())
	return !info.Mode().IsRegular() || string(info.Mode().Perm()) == "-rwx"
}

func GetOwner(path string) (string, error) {
	funcLogger := log.With().Str("func", "GetOwner").Logger()
	funcLogger.Trace().Msg(path)
	cmd := exec.Command("command", "stat", "-c", "%u", path)
	funcLogger.Trace().Msg(cmd.String())
	out, err := cmd.Output()
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return "", err
	}
	funcLogger.Trace().Msgf("%s", strings.TrimSpace(string(out)))
	return strings.TrimSpace(string(out)), nil
}

func FileNotOwned(path string) bool {
	funcLogger := log.With().Str("func", "FileNotOwned").Logger()
	funcLogger.Trace().Msg("FileNotOwned: " + path)
	owner, err := GetOwner(path)
	if err != nil {
		funcLogger.Info().Msg("FileNotOwned: " + err.Error())
		return false
	}
	funcLogger.Trace().Msg("FileNotOwned: " + owner + " and " + strconv.Itoa(os.Geteuid()))
	return owner != strconv.Itoa(os.Geteuid())
}

func GetGroup(path string) (string, error) {
	funcLogger := log.With().Str("func", "GetGroup").Logger()
	cmd := exec.Command("stat", "-u", "%g", path)
	out, err := cmd.Output()
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return "", err
	}
	funcLogger.Trace().Msgf("%s", strings.TrimSpace(string(out)))
	return strings.TrimSpace(string(out)), nil
}

func FileNotGrpowned(path string) bool {
	funcLogger := log.With().Str("func", "FileNotGrpowned").Logger()
	funcLogger.Trace().Msg(path)
	group, err := GetGroup(path)
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return false
	}
	userGroupsCmd := exec.Command("id", "-G", os.Getenv("USER"))
	userGroupsOut, err := userGroupsCmd.Output()
	if err != nil {
		funcLogger.Info().Msg(err.Error())
		return false
	}
	userGroups := strings.Fields(string(userGroupsOut))
	funcLogger.Trace().Msgf("%s", userGroups)
	return !strings.Contains(" "+strings.Join(userGroups, " ")+" ", " "+group+" ")
}
