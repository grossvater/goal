package goal

import (
	"testing"
)

func TestUnitializedtLogger(t *testing.T) {
	t.Run("testGetRoot", func(t *testing.T) { testGetRoot(&XT{t}) })
	t.Run("testGetRootByName", func(t *testing.T) { testGetRootByName(&XT{t}) })
	t.Run("testGetCreateLogger", func(t *testing.T) { testGetCreateLogger(&XT{t}) })
	t.Run("testGetCreateLogger", func(t *testing.T) { testGetCreateLogger(&XT{t}) })
}

func ExampleConsoleLogger() {
	l, _ := GetLogger("test")

	l.Log(Info(), "hello")
	// Output: INFO | test | goal_test.go:17 | hello
}

func ExampleSetLevel() {
	l, _ := GetLogger("test")

	l.Log(Info(), "start")
	l.SetLevel(Warn())
	l.Log(Info(), "end")	// not logged
	// Output: INFO | test | goal_test.go:24 | start
}

func testGetRoot(t *XT) {
	l, err := GetRootLogger()

	t.assertNotNull(l, err)
}

func testGetRootByName(t *XT) {
	l, err := GetLogger(RootLogger)

	t.assertNotNull(l, err)
}

func testGetCreateLogger(t *XT) {
	l, err := GetLogger("goal")

	t.assertNotNull(l, err)
}

func TestLoggerName(t *testing.T) {
	t.Run("testValidName1", func(t *testing.T) { testValidName(&XT{t}, "test") })
	t.Run("testValidName2", func(t *testing.T) { testValidName(&XT{t}, "parent.child") })
	t.Run("testValidName3", func(t *testing.T) { testValidName(&XT{t}, " ") })
	t.Run("testValidName4", func(t *testing.T) { testValidName(&XT{t}, " . ") })

	t.Run("testInvalidName1", func(t *testing.T) { testInvalidName(&XT{t}, "") })
	t.Run("testInvalidName2", func(t *testing.T) { testInvalidName(&XT{t}, ".") })
	t.Run("testInvalidName3", func(t *testing.T) { testInvalidName(&XT{t}, "parent.") })
	t.Run("testInvalidName4", func(t *testing.T) { testInvalidName(&XT{t}, ".child") })
}

func testValidName(t *XT, name string) {
	l, err := GetLogger(name)
	t.assertNotNull(l, err)
}

func testInvalidName(t *XT, name string) {
	l, err := GetLogger(name)
	t.assertNull(l, err)
}