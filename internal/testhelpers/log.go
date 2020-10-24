package testhelpers

import (
	"bytes"
	"fmt"

	"github.com/go-kit/kit/log"
)

type TestLogger struct {
}

func (logger TestLogger) Log(args ...interface{}) error {
	fmt.Println(args...)
	return nil
}

func FakeLogger() log.Logger {
	return TestLogger{}
}

func LoggerWithWriter() (log.Logger, *bytes.Buffer) {
	writer := &bytes.Buffer{}
	return log.NewLogfmtLogger(writer), writer
}
