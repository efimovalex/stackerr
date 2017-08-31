# StackErr
[![Build Status](https://travis-ci.org/efimovalex/stackerr.svg?branch=master)](https://travis-ci.org/efimovalex/stackerr)
[![Go Report Card](https://goreportcard.com/badge/github.com/efimovalex/stackerr)](https://goreportcard.com/report/github.com/efimovalex/stackerr) [![codecov](https://codecov.io/gh/efimovalex/stackerr/branch/master/graph/badge.svg)](https://codecov.io/gh/efimovalex/stackerr)

An error implementation with StatusCode and Stacktrace

## Install

```console
$ go get github.com/tools/godep
```

## Usage

```Go
package main

import "github.com/efimovalex/stackerr"

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

	err.Log()
}
```
Output:

```console
2017/08/31 11:59:54 Error Stacktrage:
-> github.com/efimovalex/stackerr/example/main.go:24 (main.main)
-> github.com/efimovalex/stackerr/example/main.go:18 (main.(*t1).f3)
-> github.com/efimovalex/stackerr/example/main.go:11 (main.f2)
-> github.com/efimovalex/stackerr/example/main.go:6 (main.f1)
```
## Authors

Created and maintained by

Efimov Alex - @efimovalex

## License

MIT