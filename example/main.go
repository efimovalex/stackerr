package main

import (
	"fmt"
	"net/http"

	"github.com/efimovalex/stackerr"
)

func f1() *stackerr.Err {
	err := stackerr.NewWithStatusCode("message", http.StatusNotFound)
	return err.Stack()
}

func f2() *stackerr.Err {
	err := f1()
	return err.StackWithContext("context")
}

type t1 struct{}

func (t *t1) f3() *stackerr.Err {
	err := f2()
	return err.Stack()
}

func main() {
	ts := t1{}
	err := ts.f3()

	fmt.Println(err.Sprint())

	fmt.Println(err.Error())

	fmt.Println(err.StatusCode)

	if err.IsNotFound() {
		fmt.Println("Resource is not found")
	}

	newErr := stackerr.NewFromError(err)

	err.Log()

	newErr.Log()
}
