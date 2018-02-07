package utils

import (
	"fmt"
	"runtime"
)

func stack() []byte {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return buf[:n]
}

func GetStackInfo() string {
	stackInfo := stack()
	return fmt.Sprintf("%s", stackInfo)
}
