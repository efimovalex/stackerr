package stackerr

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

// Err error
type Err struct {
	Message    string
	StatusCode int

	stackCount int
	stacktrace *Stack
}

func (e *Err) Error() string {
	return e.Message
}

// New - creates a new Err struct
func New(message string) *Err {
	err := Err{
		Message: message,
	}
	return err.Stack()
}

// Stack - Adds the current function and link to the previous Stacktrace
func (e *Err) Stack() *Err {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[1])
	file, line := f.FileLine(pc[1])

	if os.Getenv("GOPATH") != "" {
		file = strings.Replace(file, os.Getenv("GOPATH")+"/src/", "", -1)
	}

	e.stacktrace = &Stack{File: file, Function: getFunctionName(f.Name()), Line: line, CallbackStack: e.stacktrace}
	e.stackCount++
	return e
}

// Sprint - returns a pretty printed string of the Stacktrace ready for printng
func (e *Err) Sprint() string {
	return e.stacktrace.Sprint()
}

// Print - prints the Stacktrace
func (e *Err) Print() {
	fmt.Print(e.stacktrace.Sprint())
}

// Log - logs the Stacktrace using the native log package
func (e *Err) Log() {
	log.Print(e.stacktrace.Sprint())
}

func getFunctionName(name string) string {
	parts := strings.Split(name, "/")

	return parts[len(parts)-1]
}
