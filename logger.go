package goal

import (
	"log"
	"path"
	"runtime"
	"sync"
)

const RootLogger = "root"

const LoggerNameSep = "."

// Logger is the interface providing the logging contract.
type Logger interface {
	// Gets the logging level of the logger.
	GetLevel() Level

	// Sets the logging level. All messages
	SetLevel(level Level)

	// Logs passed message at the specified level.
	Log(level Level, msg string, args...interface{})
}

type logger struct {
	name string
	parent *logger
	level Level
	backends map[string]Backend

	m *sync.Mutex
}

func makeLogger(name string, parent *logger, level Level, backends map[string]Backend) *logger {
	return &logger{name: name, parent: parent, level: level, backends: backends, m: &sync.Mutex{}}
}

func (l *logger) GetLevel() Level {
	l.m.Lock()
	level := l.level
	l.m.Unlock()

	return level
}

func (l *logger) SetLevel(level Level) {
	l.m.Lock()
	l.level = level
	l.m.Unlock()
}

func (l *logger) Log(newLevel Level, msg string, args...interface{}) {
	l.m.Lock()
	level := l.level
	backends := l.backends
	l.m.Unlock()

	if newLevel.Value() > level.Value() {
		return;
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		// TODO: configure what to put here
		file = ""
		line = -1
	} else {
		file = path.Base(file)
	}

	for _, b := range backends {
		if newLevel.Value() > b.GetLevel().Value() {
			continue
		}

		err := b.Log(Message{logger: l.name, text: msg, source: file, line: line, level: newLevel})

		// TODO: report errors back to the caller?
		if err != nil {
			log.Println(err)
		}
	}
}

func (l *logger) getBackends() map[string]Backend {
	return l.backends
}
