package callername

import (
	"fmt"
	"runtime"
	"strings"
)

// CallerName returns name of caller function
func CallerName() string {
	pc, _, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}
	return ""
}

// MiddlewareCallerName returns name of caller function from inside middleware
func MiddlewareCallerName(middlware string) string {
	var ptrs [80]uintptr // NOTE: usually need a length ~8, choosing 10x larger
	callerCount := runtime.Callers(0, ptrs[:])
	frames := runtime.CallersFrames(ptrs[:])

	entered := false
	var functionName string
	for i := 0; i < callerCount; i++ {
		frame, _ := frames.Next()
		match := strings.HasSuffix(frame.File, middlware)

		// if we entered and no longer matched we have exited
		if entered && !match {
			functionName = fnName(frame.Function)
			break
		}
		entered = match
	}
	return functionName
}

// PrintStack prints out all frames for inspecting/debugging
func PrintStack() {
	var ptrs [32]uintptr // NOTE: usually need a length ~8, choosing 10x larger
	callerCount := runtime.Callers(0, ptrs[:])
	frames := runtime.CallersFrames(ptrs[:])

	for i := 0; i < callerCount; i++ {
		frame, _ := frames.Next()
		fmt.Println("-----")
		fmt.Println(frame.File)
		fmt.Println(frame.Function)
	}
}

func fnName(s string) string {
	split := strings.Split(s, "/")
	name := split[len(split)-1]
	if name == "" {
		return "unknown"
	}
	return name
}
