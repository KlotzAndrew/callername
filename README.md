# CallerName

[![Build Status](https://travis-ci.com/KlotzAndrew/callername.svg?branch=master)](https://travis-ci.com/KlotzAndrew/callername)

Find the caller name of functions used in middleware

### Use case

You want to insert middleware that logs data about where long requests are and
who is making them

### Install

```
go get github.com/KlotzAndrew/callername
```

### Setup

The main funciton `callername.MiddlewareCallerName` is looking for when a
middleware stack is entered and exited, the function name on exit being the
caller from our source code

To know this, we need to peek at the existing callstack. So for example in your
`foo` function

```golang
package foo

func bar() {
  callername.PrintStack() // add a printstack call to inspect
}
```

we will get some callstack like this:

```
/usr/local/go/src/runtime/extern.go
/home/klotz/code/go/src/callername/callername.go
/home/klotz/code/go/src/callername/callername_test.go
/usr/local/go/src/net/http/client.go
/usr/local/go/src/net/http/client.go
/usr/local/go/src/net/http/client.go
/usr/local/go/src/net/http/client.go
/usr/local/go/src/net/http/client.go
/home/klotz/code/go/src/callername/callername_test.go
/home/klotz/code/go/src/callername/callername_test.go
/usr/local/go/src/testing/testing.go
/usr/local/go/src/runtime/asm_amd64.s
```

We can see that our callstack enters the net/http library in net/http/client.go,
then comes back to our source code. Taking this information, we replace
`callername.PrintStack()` with `callername.MiddlewareCallerName`. The matcher
uses string suffix, so any around "net/http/client.go" will work

```golang
package foo

func bar() {
  name := callername.MiddlewareCallerName("net/http/client.go")
  fmt.Println(name) // foo.bar
}
```
