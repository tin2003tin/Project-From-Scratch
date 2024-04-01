package main

import (
	db_constant "database/db/constant"
	"database/runtime"
	"fmt"
)

func main() {
	app := runtime.CreateApp()
	app.LoadMetaDatabase(db_constant.DATABASE_PATH)
	app.InitCompiler()
	app.OpenServer("5432", HandleError)
}

func HandleError(err error) {
	fmt.Println(err)
	return
}
