package ez_log

import (
	"runtime"
	"time"
)

type stdLogger struct {
	level     *AtomicLevel
	formatter LogFormatter
	filter    CategoryFilter

	out WriteSyncer
	err WriteSyncer
}

func NewStdLogger(level Level, formatter LogFormatter, w, e WriteSyncer) Logger {
	return &stdLogger{
		level:     NewAtomicLevel(level),
		formatter: formatter,
		filter:    filter,
		out:       w,
		err:       e,
	}
}

func NewStdCombineLogger(level Level, formatter LogFormatter, w WriteSyncer) Logger {
	return &stdLogger{
		level:     NewAtomicLevel(level),
		formatter: formatter,
		filter:    filter,
		out:       w,
		err:       w,
	}
}

func (s *stdLogger) SetFilter(filter CategoryFilter) {
	s.filter = filter
}

func (s *stdLogger) GetFilter() CategoryFilter {
	return s.filter
}

func (s *stdLogger) SetLevel(level Level) {
	s.level.Store(level)
}

func (s *stdLogger) Enable(level Level) bool {
	return s.level.Enable(level)
}
func (s *stdLogger) enable(level Level, category *Category) bool {
	return s.level.Enable(level) && s.filter.Enable(category)
}

func (s *stdLogger) GetLevel() Level {
	return s.level.Load()
}

func getCallerFrame(skip int) (frame runtime.Frame) {
	const skipOffset = 2 // skip getCallerFrame and Callers

	pc := make([]uintptr, 1)
	numFrames := runtime.Callers(skip+skipOffset, pc[:])
	if numFrames < 1 {
		return
	}

	frame, _ = runtime.CallersFrames(pc).Next()
	return frame
}

func (s *stdLogger) TraceSkip(skip int, category *Category, fields ...Field) {
	if s.enable(TraceLevel, category) {
		s.formatter.FormatWrite(s.out, time.Now(), TraceLevel, category, fields, Caller(getCallerFrame(skip)))
	}
}

func (s *stdLogger) Trace(category *Category, fields ...Field) {
	s.TraceSkip(0, category, fields...)
}

func (s *stdLogger) Debug(category *Category, fields ...Field) {
	if s.enable(DebugLevel, category) {
		s.formatter.FormatWrite(s.out, time.Now(), DebugLevel, category, fields)
	}
}

func (s *stdLogger) Info(category *Category, fields ...Field) {
	if s.enable(InfoLevel, category) {
		s.formatter.FormatWrite(s.out, time.Now(), InfoLevel, category, fields)
	}
}

func (s *stdLogger) Warn(category *Category, fields ...Field) {
	if s.enable(WarnLevel, category) {
		s.formatter.FormatWrite(s.out, time.Now(), WarnLevel, category, fields)
	}
}

func (s *stdLogger) Error(category *Category, fields ...Field) {
	if s.enable(ErrorLevel, category) {
		s.formatter.FormatWrite(s.err, time.Now(), ErrorLevel, category, fields)
	}
}

func (s *stdLogger) Panic(category *Category, fields ...Field) {
	if s.enable(PanicLevel, category) {
		s.formatter.FormatWrite(s.err, time.Now(), PanicLevel, category, fields)
	}
}

func (s *stdLogger) Fatal(category *Category, fields ...Field) {
	if s.enable(FatalLevel, category) {
		s.formatter.FormatWrite(s.err, time.Now(), FatalLevel, category, fields)
	}
}

func (s *stdLogger) Sync() error {
	err := s.out.Sync()
	if err != nil {
		return err
	}
	if s.out != s.err {
		err = s.err.Sync()
	}
	return err
}
