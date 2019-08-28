package bmlog_test

import (
	"blackmirror/bmlog"
	"os"
)

func ExampleStandardLogger() {
	os.Setenv("LOGGER_USER", "example")
	os.Setenv("LOGGER_DEBUG", "false")
	os.Setenv("LOG_PATH", "/home/jeorch/work/test/temp/go.log")
	bmlog.StandardLogger().Info("Example Test Info")
}
