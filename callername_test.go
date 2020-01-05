package callername_test

import (
	"net/http"
	"testing"

	"github.com/klotzandrew/callername"
)

func TestGetName(t *testing.T) {
	fnName := foo()

	expected := "github.com/klotzandrew/callername_test.foo"
	if fnName != expected {
		t.Fatalf("%s != %s", fnName, expected)
	}
}

func foo() string {
	return callername.CallerName()
}

type testTripper struct {
	t         *testing.T
	middlware string
}

func (t *testTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fnName := callername.MiddlewareCallerName(t.middlware)

	expected := "callername_test.bar"
	if fnName != expected {
		t.t.Fatalf("%s != %s", fnName, expected)
	}

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
func TestGetTripperSuffix(t *testing.T) {
	bar(t, "net/http/client.go")
}

func TestGetTripperContains(t *testing.T) {
	bar(t, "net/http")
}

func bar(t *testing.T, middlware string) {
	tripper := &testTripper{t: t, middlware: middlware}
	client := &http.Client{Transport: tripper}

	_, err := client.Get("https://google.ca")
	if err != nil {
		panic(err)
	}
}
