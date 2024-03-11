package warehouse

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"server/app"
)

func LoadFileToWH(req app.Request) {
	filename := req.Params["file"];
	 filePath := filepath.Join("warehouse", filename)
	 fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	_, err = (*req.Conn).Write(fileContent)
	if err != nil {
		fmt.Println("Error sending file contents to client:", err)
		return
	}
}

func LookAllFileWH(req app.Request) {
	dir := "./warehouse"

    d, err := os.Open(dir)
    if err != nil {
        fmt.Println("Error opening directory:", err)
        return
    }
    defer d.Close()

    files, err := d.Readdir(-1)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }

	var warehouseFiles []string
    for _, file := range files {
        if file.Mode().IsRegular() {
            warehouseFiles = append(warehouseFiles, file.Name()+", ")
        }
    }

    _, err = (*req.Conn).Write([]byte(fmt.Sprintf("Files in %s:\n %v", dir, warehouseFiles)))
    if err != nil {
        fmt.Println("Error sending file contents to client:", err)
        return
    }
}