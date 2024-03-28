package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

type Book struct {
	id int
}


func (b *Book) GetId() int {
	return b.id
}

func getFuncName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	lastDotIndex := strings.LastIndex(fullName, ".")
	if lastDotIndex >= 0 {
		fullName = fullName[lastDotIndex+1:]
	}

	dashIndex := strings.Index(fullName, "-")
	if dashIndex >= 0 {
		fullName = fullName[:dashIndex]
	}

	return fullName
}

func main() {
	book := Book{
		id: 101,
	}
	execute := "GetId"
	set := []func() int{
		book.GetId,
	}
	fmt.Println(getFuncName(set[0]))
	if getFuncName(set[0]) == execute {
		id := set[0]()
		fmt.Println("Book ID:", id)
	}
}
