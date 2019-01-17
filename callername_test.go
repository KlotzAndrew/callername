package callername_test

import (
	"net/http"
	"testing"

	"callername"
)

func TestGetName(t *testing.T) {
	fnName := foo()

	expected := "callername_test.foo"
	if fnName != expected {
		t.Fatalf("%s != %s", fnName, expected)
	}
}

func foo() string {
	return callername.CallerName()
}

type testTripper struct {
	t *testing.T
}

func (t *testTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fnName := callername.MiddlewareCallerName("net/http/client.go")

	expected := "callername_test.bar"
	if fnName != expected {
		t.t.Fatalf("%s != %s", fnName, expected)
	}

	// callername.PrintStack()

	return &http.Response{}, nil
}

// NOTE: looking for first func name after exiting net/http file:
// /usr/local/go/src/runtime/extern.go
// /home/potato/code/go/src/callername/callername.go
// /home/potato/code/go/src/callername/callername_test.go
// /usr/local/go/src/net/http/client.go
// /usr/local/go/src/net/http/client.go
// /usr/local/go/src/net/http/client.go
// /usr/local/go/src/net/http/client.go
// /usr/local/go/src/net/http/client.go
// /home/potato/code/go/src/callername/callername_test.go
// /home/potato/code/go/src/callername/callername_test.go
// /usr/local/go/src/testing/testing.go
// /usr/local/go/src/runtime/asm_amd64.s
func TestGetTripper(t *testing.T) {
	bar(t)
}

func bar(t *testing.T) {
	tripper := &testTripper{t: t}
	client := &http.Client{Transport: tripper}

	_, err := client.Get("https://google.ca")
	if err != nil {
		panic(err)
	}
}
