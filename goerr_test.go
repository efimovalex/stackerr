package stackerr

import (
	"bytes"
	"errors"
	"fmt"
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
	path, _ := os.Getwd()
	ts := t1{}
	err := ts.f3()

	assert.NotNil(t, err)
	assert.Equal(t, "message", err.Error())
	assert.Equal(t,
		fmt.Sprintf(`Error Stacktrace:
-> %[1]s/goerr_test.go:36 (stackerr.TestStackTrace) 
-> %[1]s/goerr_test.go:30 (stackerr.(*t1).f3) 
-> %[1]s/goerr_test.go:23 (stackerr.f2) context
-> %[1]s/goerr_test.go:18 (stackerr.f1) 
`, path), err.Sprint())
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
	path, _ := os.Getwd()
	var buf bytes.Buffer
	log.SetOutput(&buf)
	err := f2()
	err.Log()
	log.SetOutput(os.Stderr)
	assert.Contains(t, buf.String(),
		fmt.Sprintf(`Error Stacktrace:
-> %[1]s/goerr_test.go:83 (stackerr.TestLog) 
-> %[1]s/goerr_test.go:23 (stackerr.f2) context
-> %[1]s/goerr_test.go:18 (stackerr.f1) 
`, path))
}

func TestIsNotFound(t *testing.T) {
	err := NewWithStatusCode("message", http.StatusOK)

	assert.False(t, err.IsNotFound())

	errNotFound := NewWithStatusCode("message", http.StatusNotFound)

	assert.True(t, errNotFound.IsNotFound())
}

func TestPrint(t *testing.T) {
	path, _ := os.Getwd()
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
		fmt.Sprintf(`Error Stacktrace:
-> %[1]s/goerr_test.go:117 (stackerr.TestPrint) 
-> %[1]s/goerr_test.go:23 (stackerr.f2) context
-> %[1]s/goerr_test.go:18 (stackerr.f1) 
`, path))
}
