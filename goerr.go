package stackerr

import (
	"fmt"
	"log"
	"net/http"
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

// New - creates a new Err struct with 500 Error Code
func New(message string) *Err {
	err := Err{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
	}

	return err.Stack()
}

// NewWithStatusCode - creates a new Err struct with a custom Error Code
func NewWithStatusCode(message string, statusCode int) *Err {
	err := Err{
		Message:    message,
		StatusCode: statusCode,
	}

	return err.Stack()
}

// NewFromError - creates a new Err struct with a custom Error Code
func NewFromError(e error) *Err {
	err := Err{
		Message:    e.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	return err.Stack()
}

// IsNotFound - return true if the error type is resource not found.
func (e *Err) IsNotFound() bool {
	return e.StatusCode == http.StatusNotFound
}

// Stack - Adds the current function to Err and a link to the previous Stacktrace
func (e *Err) Stack() *Err {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[1])
	file, line := f.FileLine(pc[1])

	if os.Getenv("GOPATH") != "" {
		file = strings.Replace(file, os.Getenv("GOPATH")+"/src/", "", -1)
	}

	e.stacktrace = &Stack{File: file, Function: getFunctionName(f), Line: line, CallbackStack: e.stacktrace}
	e.stackCount++
	return e
}

// StackWithContext - Adds the current function and context to Err and a link to the previous Stacktrace
func (e *Err) StackWithContext(context string) *Err {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[1])
	file, line := f.FileLine(pc[1])

	if os.Getenv("GOPATH") != "" {
		file = strings.Replace(file, os.Getenv("GOPATH")+"/src/", "", -1)
	}
	e.stacktrace.Context = context
	e.stacktrace = &Stack{File: file, Function: getFunctionName(f), Line: line, CallbackStack: e.stacktrace}
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

func getFunctionName(f *runtime.Func) string {
	parts := strings.Split(f.Name(), "/")

	return parts[len(parts)-1]
}
