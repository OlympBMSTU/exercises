package logger

import (
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	CONSOLE = "console"
	FILE = "file"
	BOTH = "console_and_file"
	FULL_LOG = 2 
	
	MIN_LOG = 0
)


var (
	LogE = log.New(LogWriter{}, "ERROR: ", 0)
	LogW = log.New(LogWriter{}, "WARN: ", 0)
	LogI = log.New(LogWriter{}, "INFO: ", 0)
)

type LogWriter struct{}

func (f LogWriter) Write(p []byte) (n int, err error) {
	pc, file, line, ok := runtime.Caller(4)
	if !ok {
		file = "?"
		line = 0
	}

	fn := runtime.FuncForPC(pc)
	var fnName string
	if fn == nil {
		fnName = "?()"
	} else {
		dotName := filepath.Ext(fn.Name())
		fnName = strings.TrimLeft(dotName, ".") + "()"
	}

	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, string(p))
	return len(p), nil
}

func infoFunc() {
	LogI.Println("information message")
}

func warnFunc() {
	LogW.Println("warning message")
}

func errorFunc() {
	LogE.Println("error message")
}

type ILogger interface {
	InfoString(string)
	Info(...interface{})
	WarnString(string)
	Warn(...interface{}) 
	Error(string, error)
}

type ConsoleLogger struct {
}

type FileLogger struct {
}

type BothLogger struct {
	logger FileLogger
	logger ConsoleLogger 

}

type Logger struct {
	// file *os.File
	// pathFile string
	// level int
	*logger ILogger 
}

func New(config config.Config) (Logger, error) {
	if config.LoggerType() == BOTH {
		
	}

	if config.LoggerType() == CONSOLE {

	}

	if config.LoggerType() == FILE {}
}

func (logger Logger) Warn(warn Stirng) {
	if level == FULL_LOG {
		logger.Warn()
	}
}

func (logger Logger) Warn(warn string, args ...interface{} ) {
	if level == FULL_LOG {
		warnSting = 
	}
}
}
// import (
// 	"io"
// 	"log"
// 	"path/filepath"
// 	"runtime"
// 	"strings"
// )

// const (
// 	LOG_ERROR  = 0
// 	LOG_WARING = 1
// 	LOG_INFO   = 2
// )

// var (
// 	LogE *log.Logger
// 	LogW *log.Logger
// 	LogI *log.Logger
// )

// func InitLogger(out io.Writer, prefix string, level int, typeLog int) *log.Logger {
// 	return log.New(out, prefix, level)
// }

// func GetLogger

// type LogWriter struct{}

// func (f LogWriter) Write(p []byte) (n int, err error) {
// 	pc, file, line, ok := runtime.Caller(4)
// 	if !ok {
// 		file = "?"
// 		line = 0
// 	}

// 	fn := runtime.FuncForPC(pc)
// 	var fnName string
// 	if fn == nil {
// 		fnName = "?()"
// 	} else {
// 		dotName := filepath.Ext(fn.Name())
// 		fnName = strings.TrimLeft(dotName, ".") + "()"
// 	}

// 	log.Printf("%s:%d %s: %s", filepath.Base(file), line, fnName, p)
// 	return len(p), nil
// }

// func LogError(err error, params ...interface{}) {

// }
// func infoFunc() {
// 	LogI.Println("information message")
// }

// func warnFunc() {
// 	LogW.Println("warning message")
// }

// func errorFunc() {
// 	LogE.Println("error message")
// }
