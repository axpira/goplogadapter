package json

import (
	"encoding/json"
	"io"
	"os"

	"github.com/axpira/gop/log"
	gla "github.com/axpira/goplogadapter"
)

var (
	Output     io.Writer = os.Stdout
	LevelFuncs           = map[log.Level]LevelConfig{
		log.TraceLevel: {"trace", nil},
		log.DebugLevel: {"debug", nil},
		log.InfoLevel:  {"info", nil},
		log.WarnLevel:  {"warn", nil},
		log.ErrorLevel: {"error", nil},
		log.FatalLevel: {"fatal", func(string) { os.Exit(1) }},
		log.PanicLevel: {"panic", func(msg string) { panic(msg) }},
	}

	JsonFormatter = func(level log.Level, msg string, err error, m map[string]interface{}) {
		levelCfg := getLevelConfig(level)
		if levelCfg.Hook != nil {
			defer levelCfg.Hook(msg)
		}
		m["msg"] = msg
		if err != nil {
			m["err"] = err.Error()
		}
		m["level"] = levelCfg.Name
		enc := json.NewEncoder(Output)
		err = enc.Encode(m)
		if err != nil {
			panic(err)
		}
	}
)

type LevelConfig struct {
	Name string
	Hook func(string)
}

func getLevelConfig(level log.Level) LevelConfig {
	levelCfg, ok := LevelFuncs[level]
	if !ok {
		levelCfg = LevelConfig{"info", nil}
	}
	return levelCfg
}

func init() {
	log.DefaultLogger = gla.New(gla.WithFormatter(JsonFormatter))
}
