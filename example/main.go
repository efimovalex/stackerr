package main

import "github.com/efimovalex/stackerr"
import "fmt"

func f1() *stackerr.Err {
	err := stackerr.Error("message")
	return err.Stack()
}

func f2() *stackerr.Err {
	err := f1()
	return err.Stack()
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

	err.Log()
}
