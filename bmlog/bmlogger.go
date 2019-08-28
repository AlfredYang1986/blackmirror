// Package bmlog is log-lib in BlackMirror's GoLibs. Depends on "github.com/sirupsen/logrus".
package bmlog

import (
	"blackmirror/bmerror"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var (
	once       = sync.Once{}
	loggerUser = "default"
	bmLogger   = logrus.StandardLogger()
)

func StandardLogger() *logrus.Entry {
	once.Do(func(){
		//Set display logging position.
		bmLogger.SetReportCaller(true)
		//Format logging position
		//bmLogger.Formatter = &logrus.TextFormatter{
		//	CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		//		filename := strings.Replace(frame.File, os.Getenv("GOPATH"), "", -1)
		//		return fmt.Sprintf("%s", frame.Func), fmt.Sprintf("%s:%d", filename, frame.Line)
		//	},
		//}

		loggerDebugInEnv := os.Getenv("LOGGER_DEBUG")
		if loggerDebugInEnv != "true" {
			loggerPathInEnv := os.Getenv("LOG_PATH")
			file, err := os.OpenFile(loggerPathInEnv, os.O_CREATE|os.O_WRONLY, 0666)
			bmerror.PanicError(err)
			bmLogger.SetOutput(file)
		}

		loggerUserInEnv := os.Getenv("LOGGER_USER")
		if loggerUserInEnv != "" {
			loggerUser = loggerUserInEnv
		}

	})
	return bmLogger.WithFields(logrus.Fields{
		"LoggerUser" : loggerUser,
	})
}