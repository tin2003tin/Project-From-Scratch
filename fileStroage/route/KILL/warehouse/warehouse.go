package warehouse

import (
	"fmt"
	"os"
	"system/app"
)

func DestroyFile(req app.Request, res app.Response) error {
	filename := req.Params["file"]
	err := os.Remove("warehouse/"+filename)
	if err != nil {
		fmt.Println("Error to delete file or the file is not found ," + filename);
		return err;
	}
	text := filename + " is Deleted"
	fmt.Println(text)
	res.SetMessage(text).Send()
	return nil;
}