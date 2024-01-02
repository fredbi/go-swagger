// Copyright 2015 go-swagger maintainers
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"

	colorizedlog "github.com/charmbracelet/log"
)

var (
	// Debug when the env var DEBUG or SWAGGER_DEBUG is not empty
	// the generators will be very noisy about what they are doing
	Debug = os.Getenv("DEBUG") != "" || os.Getenv("SWAGGER_DEBUG") != ""

	// debugLogger is a debug logger for this package
	debugLogger genLogger

	// defaultLogger is a logger to report about codegen actions
	defaultLogger genLogger
)

type (
	genLogger interface {
		Printf(format string, v ...any)
		Println(v ...any)

		Fatal(v ...any)
		Fatalf(format string, v ...any)
		Fatalln(v ...any)

		SetOutput(w io.Writer)

		/*
			Debug(v ...any)
			Debugf(format string, v ...any)
		*/
		Info(v ...any)
		Infof(format string, v ...any)

		Warn(v ...any)
		Warnf(format string, v ...any)
	}

	logger struct {
		*colorizedlog.Logger
	}

	loggerOptions struct {
	}
)

func (l logger) Fatal(v ...any) {
	if len(v) == 0 {
		l.Logger.Fatal("")
	}
	l.Logger.Fatal(v[0], v[1:])
}

func (l logger) Warn(v ...any) {
	if len(v) == 0 {
		l.Logger.Warn("")
	}
	l.Logger.Warn(v[0], v[1:])
}

func (l logger) Info(v ...any) {
	if len(v) == 0 {
		l.Logger.Info("")
	}
	l.Logger.Info(v[0], v[1:])
}

func (l logger) Fatalln(v ...any) {
	l.Fatal(v...)
}
func (l logger) Println(v ...any) {
	if len(v) == 0 {
		l.Logger.Print("")
	}
	l.Logger.Print(v[0], v[1:])
}

/*
func (l logger) Debug(v ...any) {
}
func (l logger) Debugf(format string, v ...any) {
}
*/

func newLogger(w io.Writer) genLogger {
	return logger{
		Logger: colorizedlog.New(w),
	}
}

func initDebug() {
	// debugLogger = log.New(os.Stdout, "generator:", log.LstdFlags)
	debugLogger = newLogger(os.Stdout)
}

func initLogger() {
	// defaultLogger = log.Default()
	defaultLogger = newLogger(os.Stderr)
}

// debugLog wraps log.Printf with a debug-specific logger
func debugLog(frmt string, args ...interface{}) {
	if Debug {
		_, file, pos, _ := runtime.Caller(1)
		debugLogger.Printf("%s:%d: %s", filepath.Base(file), pos,
			fmt.Sprintf(frmt, args...))
	}
}

// debugLogAsJSON unmarshals its last arg as pretty JSON
func debugLogAsJSON(frmt string, args ...interface{}) {
	if Debug {
		var dfrmt string
		_, file, pos, _ := runtime.Caller(1)
		dargs := make([]interface{}, 0, len(args)+2)
		dargs = append(dargs, filepath.Base(file), pos)
		if len(args) > 0 {
			dfrmt = "%s:%d: " + frmt + "\n%s"
			bbb, _ := json.MarshalIndent(args[len(args)-1], "", " ")
			dargs = append(dargs, args[0:len(args)-1]...)
			dargs = append(dargs, string(bbb))
		} else {
			dfrmt = "%s:%d: " + frmt
		}
		debugLogger.Printf(dfrmt, dargs...)
	}
}
