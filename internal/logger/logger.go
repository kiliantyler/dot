package logger

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kiliantyler/dot/internal/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var packageLogger = log.With().Str("pkg", "logger").Logger()
var fileLogger = packageLogger.With().Str("file", "logger.go").Logger()

func init() {
	value, exists := os.LookupEnv("VERBOSITY")
	if exists {
		SetupLogger(value)
		return
	}
	SetupLogger("error")
}

// SetupLogger configures the global logging level based on the verbosity argument.
func SetupLogger(verbosity string) {
	funcLog := fileLogger.With().Str("func", "SetupLogger").Logger()
	var lvl zerolog.Level

	// Attempt to parse verbosity as an integer level.
	if numLvl, err := strconv.Atoi(verbosity); err == nil {
		// If successful, use the numeric value as the log level.
		lvl = zerolog.Level(numLvl)
		if lvl < zerolog.TraceLevel || lvl > zerolog.FatalLevel {
			funcLog.Warn().Msgf("Invalid numeric log level '%s'. Defaulting to 'info'.", verbosity)
			lvl = zerolog.InfoLevel
		}
	} else {
		switch verbosity {
		case "trace":
			lvl = zerolog.TraceLevel
		case "debug":
			lvl = zerolog.DebugLevel
		case "info":
			lvl = zerolog.InfoLevel
		case "warn":
			lvl = zerolog.WarnLevel
		case "error":
			lvl = zerolog.ErrorLevel
		case "fatal":
			lvl = zerolog.FatalLevel
		default:
			funcLog.Warn().Msgf("Invalid log level '%s'. Defaulting to 'info'.", verbosity)
			lvl = zerolog.InfoLevel
		}
	}
	zerolog.SetGlobalLevel(lvl)
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
	if lvl == zerolog.TraceLevel {
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
		output.FieldsExclude = []string{"func"}

		log.Logger = log.Output(output).With().Caller().Logger()
	} else {
		log.Logger = log.Output(output).With().Logger()
	}
}
