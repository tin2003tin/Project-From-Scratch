package tinCreate

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type File struct {
	Content []byte
	Size    float64
	Name    string
}

func Importfile(path string) *File {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}
	parts := strings.Split(path, "/")
	name := parts[len(parts)-1]

	file := &File{
		Content: data,
		Size:    float64(len(data)),
		Name:    name,
	}
	return file
}
