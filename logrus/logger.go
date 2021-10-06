package logrus

import (
	"github.com/axpira/gop/log"
	gla "github.com/axpira/goplogadapter"
	"github.com/sirupsen/logrus"
)

var LevelGopToLogrus = map[log.Level]logrus.Level{
	log.TraceLevel: logrus.TraceLevel,
	log.DebugLevel: logrus.DebugLevel,
	log.InfoLevel:  logrus.InfoLevel,
	log.WarnLevel:  logrus.WarnLevel,
	log.ErrorLevel: logrus.ErrorLevel,
	log.FatalLevel: logrus.FatalLevel,
	log.PanicLevel: logrus.PanicLevel,
}

var LogrusFormatter = func(level log.Level, msg string, err error, m map[string]interface{}) {
	l := logrus.WithFields(m)
	if err != nil {
		l = l.WithError(err)
	}
	l.Log(LevelGopToLogrus[level], msg)
}

func init() {
	log.DefaultLogger = gla.New(gla.WithFormatter(LogrusFormatter))
}
