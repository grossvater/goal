package goal

type nullLogger struct {
}

func NewNullLogger() Logger {
	return nullLogger{}
}

func (l nullLogger) Log(level Level, msg string, args...interface{}) {
}

func (l nullLogger) GetLevel() Level {
	return None()
}

func (l nullLogger) SetLevel(level Level) {
}