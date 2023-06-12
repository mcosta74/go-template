package main

import (
	"os"
	"path/filepath"

	"golang.org/x/exp/slog"
)

func newLogger(logLevel, logFormat string) *slog.Logger {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(logLevel)); err != nil {
		lvl = slog.LevelInfo
	}
	opts := &slog.HandlerOptions{
		Level:     lvl,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				// use UTC time
				a.Value = slog.TimeValue(a.Value.Time().UTC())

			case slog.SourceKey:
				// remove directories from File
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}
			return a
		},
	}

	var h slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if logFormat == "json" {
		h = slog.NewJSONHandler(os.Stdout, opts)
	}
	return slog.New(h)
}
