package bmlog

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestLogrus(t *testing.T) {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("Trace msg")
	logrus.Debug("Debug msg")
	logrus.Info("Info msg")
	logrus.Warn("Warn msg")
	logrus.Error("Error msg")
	//logrus.Fatal("Fatal msg")
	//logrus.Panic("Panic msg")
}

func TestStandardLogger(t *testing.T) {
	os.Setenv("LOGGER_USER", "debugger")
	os.Setenv("LOGGER_DEBUG", "true")

	StandardLogger().Info("Test Info")
}

func TestLog2File(t *testing.T) {
	os.Setenv("LOGGER_USER", "blackmirror")
	os.Setenv("LOGGER_DEBUG", "false")
	os.Setenv("LOG_PATH", "/home/jeorch/work/test/temp/go.log")
	StandardLogger().Info("TestLog2File")
}
