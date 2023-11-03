package logger

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger struct containing logging library
type Logger struct {
	Log zerolog.Logger
}

// NewLogger returns new instance of Logger struct
func NewLogger(consoleLog, fileLog bool, name string) *Logger {
	var writers []io.Writer
	if consoleLog {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stdout})
	}
	if fileLog {
		writers = append(writers, newRollingFile(name))
	}
	mw := io.MultiWriter(writers...)
	zlog := zerolog.New(mw).With().Timestamp().Logger()
	logger := Logger{
		Log: zlog,
	}
	return &logger
}

func newRollingFile(fileName string) io.Writer {
	// if err := os.MkdirAll(config.Directory, 0744); err != nil {
	// 	log.Error().Err(err).Str("path", config.Directory).Msg("can't create log directory")
	// 	return nil
	// }

	return &lumberjack.Logger{
		Filename:   path.Join(".", fmt.Sprintf("%s.log", fileName)),
		MaxBackups: 10,  // files
		MaxSize:    100, // megabytes
		MaxAge:     200, // days
	}
}
