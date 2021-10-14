package ez_log

import "os"

var DefaultLogger Logger = NewNopLogger()

func ReInitLoggerToStdio(level Level, formatter LogFormatter) {
	DefaultLogger = NewStdLogger(level, formatter, os.Stdout, os.Stderr)
}

func ReInitLoggerToWriter(level Level, formatter LogFormatter, output, errorOutput WriteSyncer) {
	DefaultLogger = NewStdLogger(level, formatter, output, errorOutput)
}

func ReInitLoggerToCombineWriter(level Level, formatter LogFormatter, combineOutput WriteSyncer) {
	DefaultLogger = NewStdCombineLogger(level, formatter, combineOutput)
}
