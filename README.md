# StackErr
[![Build Status](https://travis-ci.org/efimovalex/stackerr.svg?branch=master)](https://travis-ci.org/efimovalex/stackerr)
[![Go Report Card](https://goreportcard.com/badge/github.com/efimovalex/stackerr)](https://goreportcard.com/report/github.com/efimovalex/stackerr) [![codecov](https://codecov.io/gh/efimovalex/stackerr/branch/master/graph/badge.svg)](https://codecov.io/gh/efimovalex/stackerr) [![GoDoc](https://godoc.org/github.com/efimovalex/stackerr?status.svg)](https://godoc.org/github.com/efimovalex/stackerr)

An error implementation with StatusCode and Stacktrace

It implements the Golang error interface

You can use Status Code to better identify the answer you need to report to the client.

Makes debugging easier by logging the functions the error passes through and by adding the ability to log context on each function pass of the error, so that you can create a path of the error through your application.  

## Install

```console
$ go get github.com/efimovalex/stackerr
```

## Usage

```Go
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

	err.Log()
}

```
Output:

```console
Error Stacktrace:
-> github.com/efimovalex/stackerr/example/main.go:29 (main.main)
-> github.com/efimovalex/stackerr/example/main.go:23 (main.(*t1).f3)
-> github.com/efimovalex/stackerr/example/main.go:16 (main.f2) context
-> github.com/efimovalex/stackerr/example/main.go:11 (main.f1)

message

404

Resource is not found

2017/09/05 17:32:15 Error Stacktrace:
-> github.com/efimovalex/stackerr/example/main.go:29 (main.main)
-> github.com/efimovalex/stackerr/example/main.go:23 (main.(*t1).f3)
-> github.com/efimovalex/stackerr/example/main.go:16 (main.f2) context
-> github.com/efimovalex/stackerr/example/main.go:11 (main.f1)
```
## Authors

Created and maintained by

Efimov Alex - @efimovalex

## License

MIT
