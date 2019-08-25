package goal

type Message struct {
	logger string
	level  Level
	text   string
	source string
	line   int
}

type Backend interface {
	Log(msg Message) (err error)
	GetLevel() Level
	Shutdown() (err error)
	Flush() (err error)
}