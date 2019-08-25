package goal

import (
	"fmt"
)

const (
	ConsoleBackendName = "console"
)

func NewConsoleBackend(name string, level Level) Backend {
	return &consoleBackend{name: name, level: level}
}

func GetDefaultConsoleLevel() Level {
	return Info()
}

type consoleBackend struct {
	name string
	level Level
}

func (l *consoleBackend) Log(msg Message) error {
	_, err := fmt.Printf("%s | %s | %s:%d | %s", msg.level.Name(), msg.logger, msg.source, msg.line, msg.text)

	return err
}

func (l *consoleBackend) Shutdown() error {
	return nil
}

func (l *consoleBackend) Flush() error {
	return nil
}

func (l *consoleBackend) GetLevel() Level {
	return l.level
}