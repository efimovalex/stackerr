package stackerr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSprint(t *testing.T) {
	stack := Stack{}

	assert.Equal(t, stack.Sprint(), "Error Stacktrage:\n-> out of context\n")

	stack2 := Stack{File: "f2.go", Function: "f2.f2", Line: 3, CallbackStack: &stack}

	stack3 := Stack{File: "f1.go", Function: "f1.f2", Line: 3, CallbackStack: &stack2}
	assert.Equal(t, stack3.Sprint(),
		`Error Stacktrage:
-> f1.go:3 (f1.f2)
-> f2.go:3 (f2.f2)
-> out of context
`)
}
