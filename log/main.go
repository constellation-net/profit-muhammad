package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"

	"github.com/constellation-net/profit-muhammad/config"
)

var (
	Log *logrus.Logger
)

func LogLevel() logrus.Level {
	env := strings.ToLower(config.Config.Environment)
	switch env {
	case "production":
		return logrus.InfoLevel
	default:
		return logrus.DebugLevel
	}
}

func init() {
	Log = &logrus.Logger{
		Out: os.Stderr,
		Formatter: &easy.Formatter{
			TimestampFormat: "02/01/2006 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
		Hooks: make(logrus.LevelHooks),
		Level: LogLevel(),
	}
}

func Error(err error, trace string, fatal ...bool) {
	fatalErr := false
	if len(fatal) > 0 {
		fatalErr = fatal[0]
	}

	if err != nil {
		m := fmt.Sprintf("[trace: %s] %s", trace, err.Error())

		if fatalErr {
			Log.Fatal(m)
		} else {
			Log.Error(m)
		}
	}
}
