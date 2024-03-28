package compiler

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type SetOfFunc struct {
	Handlers []func([][]string) ([][]string, error)
}

func InitSetOfFunc(handler []func([][]string) ([][]string, error)) *SetOfFunc {
	sof := SetOfFunc{
		Handlers: handler,
	}
	return &sof
}

func getFunctionName(fn interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	lastIndex := strings.LastIndex(fullName, ".")
	if lastIndex == -1 {
		return strings.TrimSpace(fullName)
	}
	s := strings.TrimSpace(fullName[lastIndex+1:])
	ss := strings.Split(s, "-")
	return ss[0]
}

func (s *SetOfFunc) executeHandler(name string, values [][]string) ([][]string, error) {
	for _, handler := range s.Handlers {
		name = strings.TrimSpace(name)
		if getFunctionName(handler) == name {
			return handler(values)
		}
	}
	fmt.Println("handler not found")
	return nil, nil
}
