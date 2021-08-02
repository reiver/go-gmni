# go-gemini

Package **gemini** provides **Gemini Protocol** client and server implementations, for the Go programming language.

The **gemini** package provides an API in a style similar to the "net/http" library that is part of the Go standard library, including support for "middleware".


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-gemini

[![GoDoc](https://godoc.org/github.com/reiver/go-gemini?status.svg)](https://godoc.org/github.com/reiver/go-gemini)


## Example

A very very simple Gemini Protocol server is shown in the following code.

This particular Gemini Protocol server just responds to the client with the URI that was in the request, plus the remote address.


```go
package main

import (
	"github.com/reiver/go-gemini"
)

func main() {

	var handler gemini.Handler = gemini.DebugHandler

	err := gemini.ListenAndServe(":1965", handler)
	if nil != err {
		//@TODO: Handle this error better.
		panic(err)
	}
}
```
