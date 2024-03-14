package main

import (
	"fmt"
	"server/app"
	kw "server/route/KILL/warehouse"
	sw "server/route/LOOK/warehouse"
	lw "server/route/SEND/warehouse"
	tt "server/route/TINY/tin"
)

func main() {
	app := app.CreateApp();
	
	app.AddRoute("SEND","/warehouse",lw.SaveFileToWH)
	app.AddRoute("LOOK","/warehouse",sw.LookAllFileWH)
    app.AddRoute("LOOK","/warehouse/{file}",sw.LoadFileToWH)
	app.AddRoute("KILL","/warehouse/{file}",kw.DestroyFile)
	app.AddRoute("TINY","/tin",tt.TinSayHello)

	app.OpenServer("8080", handleError)
}

func handleError(err error) {
	fmt.Print("(Server) Error:",err)
}


