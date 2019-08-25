package goal

func ExampleConsoleLevel() {
	l, _ := GetLogger("test")

	// anything above Info is discarded by console backend
	l.Log(Debug(), "start")
	l.Log(Info(), "working")
	l.Log(Debug(), "end")	// not logged
	// Output: INFO | test | console_test.go:8 | working
}
