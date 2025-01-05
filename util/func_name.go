package util

import (
	"runtime"
)

func GetCurrentFuncName() string {
	var pc uintptr
	pc, _, _, _ = runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
