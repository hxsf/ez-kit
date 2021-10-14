package ez_log

type nopLogger struct {
	filter CategoryFilter
}

func (n *nopLogger) GetFilter() CategoryFilter {
	return n.filter
}

func (n *nopLogger) SetFilter(filter CategoryFilter) {
	n.filter = filter
}

func NewNopLogger() Logger {
	return &nopLogger{}
}

func (n *nopLogger) SetLevel(_ Level) {}

func (n *nopLogger) Enable(_ Level) bool {
	return false
}

func (n *nopLogger) GetLevel() Level {
	return _maxLevel
}

func (n nopLogger) TraceSkip(_ int, _ *Category, _ ...Field) {}

func (n nopLogger) Trace(_ *Category, _ ...Field) {}

func (n nopLogger) Debug(_ *Category, _ ...Field) {}

func (n nopLogger) Info(_ *Category, _ ...Field) {}

func (n nopLogger) Warn(_ *Category, _ ...Field) {}

func (n nopLogger) Error(_ *Category, _ ...Field) {}

func (n nopLogger) Panic(_ *Category, _ ...Field) {}

func (n nopLogger) Fatal(_ *Category, _ ...Field) {}

func (n nopLogger) Sync() error {
	return nil
}
