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

func CreateFileIfNotExists(filePath string) error {
	funcLogger := log.With().Str("func", "CreateFileIfNotExists").Logger()
	funcLogger.Trace().Msg(filePath)
	if err := CreateDirIfNotExist(getFolderFromFile(filePath)); err != nil {
		funcLogger.Info().Msg(err.Error())
		return err
	}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		funcLogger.Trace().Msgf("Creating file: %s", filePath)
		file, err := os.Create(filePath)
		if err != nil {
			funcLogger.Error().Msgf("Error creating file: %s", err)
			return err
		}
		file.Close()
		funcLogger.Trace().Msgf("File created: %s", filePath)
	}
	return nil
}

func getFolderFromFile(filePath string) string {
	funcLogger := log.With().Str("func", "getFolderFromFile").Logger()
	funcLogger.Trace().Msg(filePath)
	if strings.HasSuffix(filePath, "/") {
		funcLogger.Info().Msgf("Filepath is a folder: %s", filePath)
		return filePath
	}
	funcLogger.Trace().Msgf("Filepath Found: %s", filePath[:strings.LastIndex(filePath, "/")+1])
	return filePath[:strings.LastIndex(filePath, "/")+1]
}

func FileExists(filePath string) bool {
	funcLogger := log.With().Str("func", "FileExists").Logger()
	funcLogger.Trace().Msg(filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		funcLogger.Info().Msgf("File does not exist: %s", filePath)
		return false
	}
	funcLogger.Trace().Msgf("File exists: %s", filePath)
	return true
}
