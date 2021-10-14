package ez_log

import (
	"io"
	"time"
)

type Logger interface {
	SetLevel(level Level)
	GetLevel() Level
	Enable(level Level) bool
	SetFilter(CategoryFilter)
	GetFilter() CategoryFilter
	TraceSkip(skip int, category *Category, fields ...Field)
	Trace(category *Category, fields ...Field)
	Debug(category *Category, fields ...Field)
	Info(category *Category, fields ...Field)
	Warn(category *Category, fields ...Field)
	Error(category *Category, fields ...Field)
	Panic(category *Category, fields ...Field)
	Fatal(category *Category, fields ...Field)
	Sync() error
}

type WriteSyncer interface {
	io.Writer
	Sync() error
}

type LogFormatter interface {
	FormatWrite(w io.Writer, time time.Time, level Level, category *Category, fields []Field, ext ...Field)
}
