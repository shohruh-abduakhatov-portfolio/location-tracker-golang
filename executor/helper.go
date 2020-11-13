package executor

import (
	"fmt"
	"runtime"
	"strings"
)

func getCallingFunc() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	fmt.Printf(frame.Function)
	arr := strings.Split(frame.Function, ".")
	return arr[len(arr)-1]
}
