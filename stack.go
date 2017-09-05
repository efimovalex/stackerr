package stackerr

import "fmt"

// Stack - contains stack info
type Stack struct {
	File          string
	Line          int
	Function      string
	Context       string
	CallbackStack *Stack
}

// Sprint returns a pretty printed string of the Stacktrace ready for printng
func (s *Stack) Sprint() string {
	stack := "Error Stacktrace:\n"
	stackTrace := s
	for stackTrace != nil {
		if stackTrace.File != "" && stackTrace.Function != "" {
			stack += fmt.Sprintf("-> %s:%d (%s) %s\n", stackTrace.File, stackTrace.Line, stackTrace.Function, stackTrace.Context)
		} else {
			stack += fmt.Sprint("-> out of context\n")
		}

		stackTrace = stackTrace.CallbackStack
	}

	return stack
}
