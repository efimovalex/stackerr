# StackErr
[![Build Status](https://travis-ci.org/efimovalex/stackerr.svg?branch=master)](https://travis-ci.org/efimovalex/stackerr)
[![Go Report Card](https://goreportcard.com/badge/github.com/efimovalex/stackerr)](https://goreportcard.com/report/github.com/efimovalex/stackerr) [![codecov](https://codecov.io/gh/efimovalex/stackerr/branch/master/graph/badge.svg)](https://codecov.io/gh/efimovalex/stackerr) [![GoDoc](https://godoc.org/github.com/efimovalex/stackerr?status.svg)](https://godoc.org/github.com/efimovalex/stackerr)

An error implementation with StatusCode and Stacktrace 

## Install

```console
$ go get github.com/efimovalex/stackerr
```

## Usage

```Go
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
```
Output:

```console
Error Stacktrace:
-> github.com/efimovalex/stackerr/example/main.go:25 (main.main)
-> github.com/efimovalex/stackerr/example/main.go:19 (main.(*t1).f3)
-> github.com/efimovalex/stackerr/example/main.go:12 (main.f2)
-> github.com/efimovalex/stackerr/example/main.go:7 (main.f1)

message

2017/08/31 12:13:47 Error Stacktrace:
-> github.com/efimovalex/stackerr/example/main.go:25 (main.main)
-> github.com/efimovalex/stackerr/example/main.go:19 (main.(*t1).f3)
-> github.com/efimovalex/stackerr/example/main.go:12 (main.f2)
-> github.com/efimovalex/stackerr/example/main.go:7 (main.f1)
```
## Authors

Created and maintained by

Efimov Alex - @efimovalex

## License

MIT
