package main

import (
	"fmt"
	"system/app"
	kw "system/route/KILL/warehouse"
	sw "system/route/LOOK/warehouse"
	lw "system/route/SEND/warehouse"
)

func main() {
	app := app.CreateApp();

	app.AddRoute("SEND","/warehouse",lw.SaveFileToWH)
	app.AddRoute("LOOK","/warehouse",sw.LookAllFileWH)
    app.AddRoute("LOOK","/warehouse/{file}",sw.LoadFileToWH)
	app.AddRoute("KILL","/warehouse/{file}",kw.DestroyFile)
	
	app.OpenServer("8000", handleError)
}

func handleError(err error) {
	fmt.Print(err)
}

