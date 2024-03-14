package warehouse

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"system/app"
)

func LoadFileToWH(req app.Request,res app.Response) error {
	filename := req.Params["file"];
	 filePath := filepath.Join("warehouse", filename)
	 fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
    fmt.Println(string(fileContent));
	err = res.SetMessage("Found").SetBody((string(fileContent))).Send()
	if err != nil {
		fmt.Println("Error sending file contents to client:", err)
		return err
	}
    return nil
}

func LookAllFileWH(req app.Request,res app.Response) error {
	dir := "./warehouse"

    d, err := os.Open(dir)
    if err != nil {
        fmt.Println("Error opening directory:", err)
        return err
    }
    defer d.Close()

    files, err := d.Readdir(-1)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return err
    }

	var warehouseFiles []string
    for _, file := range files {
        if file.Mode().IsRegular() {
            warehouseFiles = append(warehouseFiles, file.Name()+", ")
        }
    }
    fmt.Println(warehouseFiles)
    err = res.SetMessage("Found").SetBody(strings.Join(warehouseFiles, ", ")).Send()
    if err != nil {
        fmt.Println("Error sending file contents to client:", err)
        return err
    } 
    return nil
}