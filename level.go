package goal

type Level interface {
	Value() uint
	Name() string
}

type level struct {
	value uint
	name string
}

const maxUint = ^uint(0)

var (
	levelNone Level  = level{0, "NONE"}
	levelError Level = level{100, "ERROR"}
	levelWarn Level  = level{200, "WARN"}
	levelInfo Level  = level{300, "INFO"}
	levelDebug Level = level{400, "DEBUG"}
	levelTrace Level = level{500, "TRACE"}
	levelAll Level   = level{maxUint, "ALL"}
)

func None() Level {
	return levelNone
}

func Error() Level {
	return levelError
}

func Warn() Level {
	return levelWarn
}

func Info() Level {
	return levelInfo
}

func Debug() Level {
	return levelDebug
}

func Trace() Level {
	return levelTrace
}

func All() Level {
	return levelAll
}

func (l level) Value() uint {
	return l.value
}

func (l level) Name() string {
	return l.name
}
