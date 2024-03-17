package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/kiliantyler/dot/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	SetupLogger()
}

func SetupLogger() {
	verbosity := os.Getenv("VERBOSITY")
	if verbosity == "" {
		verbosity = "error" // Default level
	}
	setGlobalLogLevel(verbosity)
}

func setGlobalLogLevel(verbosity string) {
	funcLog := log.With().Str("func", "setGlobalLogLevel").Logger()
	level, err := zerolog.ParseLevel(verbosity)
	if err != nil {
		funcLog.Warn().Msgf("Error parsing log level: %s", err)
		level = zerolog.ErrorLevel
	}

	zerolog.SetGlobalLevel(level)
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"}
	output.FormatFieldName = func(i interface{}) string {
		switch i {
		case "func":
			return ""
		}
		return fmt.Sprintf("%s: ", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		length := 7
		if i == nil {
			return "[" + utils.TrimStringToLengthUTF8("MISSING", length) + "]"
		}
		return "[" + utils.TrimStringToLengthUTF8(i.(string), length) + "]"
	}
	output.FieldsExclude = []string{"func"}
	if level == zerolog.TraceLevel {
		output.FormatCaller = func(i interface{}) string {
			length := 5
			caller := strings.Split(i.(string), "/")
			pkg := caller[len(caller)-2]
			file := caller[len(caller)-1]
			file = strings.Split(file, ".")[0]
			return "{" + utils.TrimStringToLengthUTF8(pkg, length) + "}" + "(" + utils.TrimStringToLengthUTF8(file, length) + ")"
		}
		output.PartsOrder = []string{
			zerolog.LevelFieldName,
			zerolog.TimestampFieldName,
			zerolog.CallerFieldName,
			"func",
			zerolog.MessageFieldName,
		}

		log.Logger = log.Output(output).With().Caller().Logger()
	} else {
		log.Logger = log.Output(output).With().Logger()
	}
}
