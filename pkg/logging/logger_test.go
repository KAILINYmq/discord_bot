package logging

import "testing"

func TestInitLogger(t *testing.T) {
	logger := InitLogger(WithLogPath("../../log/backend.log"), WithFormat(TextFormat))
	logger.Info("hello world 你好 👋")
}
