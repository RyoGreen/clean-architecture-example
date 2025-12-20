package logger

import (
	"os"
	"path/filepath"

	"log/slog"
)

var L *slog.Logger

func init() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		return a
	}
	L = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, ReplaceAttr: replace}))
}
