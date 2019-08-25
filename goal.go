// Package goal provides a simple logging interface.
//
//
// Loggers are organized hierarchically, every logger inheriting and overriding settings of its parent. The parent of all
// allLoggers is the root logger.
//
//	root, err := goal.GetRootLogger()
//
//	inherits from root logger
//	test, err := goal.GetLogger("test")
//
//	inherits from test logger
//	test, err := goal.GetLogger("test.goal")
//
// Every message passed to the logger has an associated level. If the configured level of the logger is greater or equals
// to the provided level, the message is actually logged, otherwise it is discarded.
//
//	test.Log(Error, "I failed.")
package goal

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// GetRootLogger gets the root logger.
func GetRootLogger() (Logger, error) {
	return getRootLogger(), nil
}

// GetLogger gets the logger with the given name. If the logger doesn't exist, it is created and connected to all
// registered backends.
func GetLogger(name string) (Logger, error) {
	var logger *logger

	if logger = allLoggers[name]; logger == nil {
		if !validateLoggerName(name) {
			return nil, fmt.Errorf("invalid logger name: %v", name)
		}

		root := getRootLogger()
		chain := getQualifiedLoggerChain(name)

		parent := root
		for i := 0; i < len(chain); i++ {
			name := chain[i]
			logger = createLoggerOnDemand(name, parent)
			registerLogger(name, logger)

			if i == len(chain) {
				return logger, nil
			}

			parent = logger
		}
	}

	return logger, nil
}

func AddLogger(name string, level Level, backends []string) error {
	var logger *logger

	if logger = allLoggers[name]; logger != nil {
		// TODO: fix me, replace the logger
		return errors.New("logger already exists")
	}

	if !validateLoggerName(name) {
		return fmt.Errorf("invalid logger name: %v", name)
	}

	for _, b := range backends {
		if _, ok := allBackends[b]; !ok {
			return fmt.Errorf("no such backend: %s", b)
		}
	}

	return nil
}

var allLoggers map[string]*logger
var allBackends map[string]Backend

var loggerNamePattern = compileRegex("^[^.]+([.][^.]+)*$")

func init() {
	console := NewConsoleBackend(ConsoleBackendName, GetDefaultConsoleLevel())

	allBackends = make(map[string]Backend)
	allBackends[ConsoleBackendName] = console

	allLoggers = make(map[string]*logger)

	root := createLoggerOnDemand(RootLogger, nil)
	registerLogger(RootLogger, root)
}

func getRootLogger() *logger {
	return allLoggers[RootLogger]
}

func createLoggerOnDemand(name string, parent *logger) *logger {
	var parentBackends map[string]Backend
	var level Level

	if parent != nil {
		parentBackends = parent.getBackends()
		level = parent.GetLevel()
	} else {
		parentBackends = make(map[string]Backend)
		level = All()
	}

	backends := make(map[string]Backend, len(parentBackends))
	for k, b := range allBackends {
		backends[k] = b
	}

	return makeLogger(name, parent, level, backends)
}

func registerLogger(name string, logger *logger) {
	allLoggers[name] = logger
}

func validateLoggerName(name string) bool {
	return loggerNamePattern.Match([]byte(name))
}

func compileRegex(p string) *regexp.Regexp {
	re, err := regexp.Compile(p)
	if err != nil {
		panic(err)
	}

	return re
}

// Gets the chain of qualified logger names for the qualified logger name
// (e.g parent.child.nephew => []{ "parent", "parent.child", "parent.child.newphew" }
func getQualifiedLoggerChain(name string) []string {
	segs := strings.Split(name, LoggerNameSep)
	chain := make([]string, 0, len(segs) + 1)

	chain = append(chain, RootLogger)
	chain = append(chain, segs...)

	parentName := chain[0]
	for i := 1; i < len(chain); i++{
		qualifiedName := parentName + LoggerNameSep + chain[i]
		parentName = qualifiedName
	}

	return chain
}