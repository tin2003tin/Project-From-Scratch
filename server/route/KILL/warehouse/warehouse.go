package warehouse

import (
	"fmt"
	"os"
	"server/app"
)

func DestroyFile(req app.Request) {
	filename := req.Params["file"]
	err := os.Remove("warehouse/"+filename)
	if err != nil {
		fmt.Println("Error to delete file or the file is not found ," + filename);
		return;
	}
	fmt.Println(filename + " is Deleted")
}