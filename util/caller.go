package util

import (
	"fmt"
	"path"
	"runtime"
)

type CallerInfo struct {
	File   string
	Line   int
	FnName string
}

func (x *CallerInfo) String() string {
	return fmt.Sprintf("%s:%d:%s", x.File, x.Line, x.FnName)
}

func getCallerInfo(nCallStackSkip int) (*CallerInfo, error) {
	pc, file, line, ok := runtime.Caller(nCallStackSkip)
	if !ok {
		return nil, fmt.Errorf("GetCallerInfo: failed to get caller pc")
	}

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return nil, fmt.Errorf("GetCallerInfo: failed to get caller func detail")
	}

	fnName := "<unknown func>"
	if fn != nil {
		fnName = fn.Name()
	}

	return &CallerInfo{
		File:   path.Base(file),
		Line:   line,
		FnName: path.Base(fnName),
	}, nil
}

func GetCallerInfo() (*CallerInfo, error) {
	return getCallerInfo(3)
}

func GetCallerInfoStr() string {
	caller, err := getCallerInfo(3)
	if err != nil {
		return ""
	}
	return caller.String()
}
