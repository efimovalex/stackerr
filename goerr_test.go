package stackerr

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func f1() *Err {
	err := New("message")
	return err.Stack()
}

func f2() *Err {
	err := f1()
	return err.StackWithContext("context")
}

type t1 struct{}

func (t *t1) f3() *Err {
	err := f2()
	return err.Stack()
}

func TestStackTrace(t *testing.T) {
	ts := t1{}
	err := ts.f3()

	assert.NotNil(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t,
		`Error Stacktrace:
-> github.com/efimovalex/stackerr/goerr_test.go:33 (stackerr.TestStackTrace) 
-> github.com/efimovalex/stackerr/goerr_test.go:27 (stackerr.(*t1).f3) 
-> github.com/efimovalex/stackerr/goerr_test.go:20 (stackerr.f2) context
-> github.com/efimovalex/stackerr/goerr_test.go:15 (stackerr.f1) 
`, err.Sprint())
}

func TestNew(t *testing.T) {
	err := New("message")

	assert.NotNil(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
}

func TestNewFromError(t *testing.T) {
	e := errors.New("message")
	err := NewFromError(e)

	assert.NotNil(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t, http.StatusInternalServerError, err.StatusCode)

	errErr := NewFromError(err)
	assert.NotNil(t, errErr)
	assert.Equal(t, "message", errErr.Error())
	assert.Equal(t, http.StatusInternalServerError, errErr.StatusCode)
}

func TestNewWithStatusCode(t *testing.T) {
	err := NewWithStatusCode("message", http.StatusOK)

	assert.NotNil(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t, http.StatusOK, err.StatusCode)
}
func TestLog(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	err := f2()
	err.Log()
	log.SetOutput(os.Stderr)
	assert.Contains(t, buf.String(),
		`Error Stacktrace:
-> github.com/efimovalex/stackerr/goerr_test.go:79 (stackerr.TestLog) 
-> github.com/efimovalex/stackerr/goerr_test.go:20 (stackerr.f2) context
-> github.com/efimovalex/stackerr/goerr_test.go:15 (stackerr.f1) 
`)
}

func TestIsNotFound(t *testing.T) {
	err := NewWithStatusCode("message", http.StatusOK)

	assert.False(t, err.IsNotFound())

	errNotFound := NewWithStatusCode("message", http.StatusNotFound)

	assert.True(t, errNotFound.IsNotFound())
}

func TestPrint(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	err := f2()
	err.Print()

	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	assert.Equal(t, out,
		`Error Stacktrace:
-> github.com/efimovalex/stackerr/goerr_test.go:112 (stackerr.TestPrint) 
-> github.com/efimovalex/stackerr/goerr_test.go:21 (stackerr.f2) context
-> github.com/efimovalex/stackerr/goerr_test.go:16 (stackerr.f1) 
`)
}
