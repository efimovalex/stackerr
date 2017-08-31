package stackerr

import (
	"bytes"
	"log"
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
	return err.Stack()
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
	assert.Equal(t, err.Error(), "message")
	assert.Equal(t, err.Sprint(),
		`Error Stacktrage:
-> github.com/efimovalex/stackerr/goerr_test.go:30 (stackerr.TestStackTrace)
-> github.com/efimovalex/stackerr/goerr_test.go:25 (stackerr.(*t1).f3)
-> github.com/efimovalex/stackerr/goerr_test.go:18 (stackerr.f2)
-> github.com/efimovalex/stackerr/goerr_test.go:13 (stackerr.f1)
`)
}

func TestError(t *testing.T) {
	err := Error("message")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "message")
	expected := Err{
		Message: "message",
	}

	assert.Equal(t, err.Error(), expected.Error())
}
func TestErrorWS(t *testing.T) {
	err := ErrorWS("message", 200)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "message")
	expected := Err{
		Message:    "message",
		StatusCode: 200,
	}

	assert.Equal(t, err.Error(), expected.Error())
	assert.Equal(t, err.StatusCode, expected.StatusCode)
}
func TestLog(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	err := f2()
	err.Log()
	log.SetOutput(os.Stderr)
	assert.Contains(t, buf.String(),
		`Error Stacktrage:
-> github.com/efimovalex/stackerr/goerr_test.go:58 (stackerr.TestLog)
-> github.com/efimovalex/stackerr/goerr_test.go:18 (stackerr.f2)
-> github.com/efimovalex/stackerr/goerr_test.go:13 (stackerr.f1)
`)
}
