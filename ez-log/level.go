package ez_log

import "sync/atomic"

type Level int32

const (
	// TraceLevel logs anything with one caller frame
	TraceLevel Level = iota - 2
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// PanicLevel logs a message, then does not panic.
	PanicLevel
	// FatalLevel logs a message, then does not call os.Exit(1).
	FatalLevel

	// OffLevel used to turn off log. nothing gets logged at all.
	OffLevel

	_minLevel = TraceLevel
	_maxLevel = OffLevel
	LevelMaxLength = 5
)

func (i Level) String() string {
	if i < _minLevel || i > _maxLevel {
		panic("level out of range")
	}
	switch i {
	case TraceLevel:
		return "Trace"
	case DebugLevel:
		return "Debug"
	case InfoLevel:
		return "Info"
	case WarnLevel:
		return "Warn"
	case ErrorLevel:
		return "Error"
	case PanicLevel:
		return "Panic"
	case FatalLevel:
		return "Fatal"
	case OffLevel:
		return "Off"
	default:
		panic("not implement")
	}
}

func NewAtomicLevel(i Level) *AtomicLevel {
	if i < _minLevel || i > _maxLevel {
		panic("level out of range")
	}
	return &AtomicLevel{int32(i)}
}

type AtomicLevel struct{ v int32 }

func (i *AtomicLevel) Load() Level {
	return Level(atomic.LoadInt32(&i.v))
}

func (i *AtomicLevel) Store(n Level) {
	if n < _minLevel || n > _maxLevel {
		panic("level out of range")
	}
	atomic.StoreInt32(&i.v, int32(n))
}

func (i *AtomicLevel) Enable(n Level) bool {
	return i.Load() <= n
}
