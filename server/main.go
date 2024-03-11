package main

import (
	"fmt"
	"server/app"
	kw "server/route/KILL/warehouse"
	sw "server/route/LOOK/warehouse"
	lw "server/route/SEND/warehouse"
)

func main() {
	app := app.CreateApp();
	
	app.AddRoute("SEND","/warehouse",lw.SaveFileToWH)
	app.AddRoute("LOOK","/warehouse",sw.LookAllFileWH)
    app.AddRoute("LOOK","/warehouse/{file}",sw.LoadFileToWH)
	app.AddRoute("KILL","/warehouse/{file}",kw.DestroyFile)

	app.OpenServer("8080", handleError)
}

func handleError(err error) {
	fmt.Print(err)
}

