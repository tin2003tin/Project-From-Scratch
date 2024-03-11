package main

import (
	"client/tinConn"
	"client/tinConn/lib"
)

func main() {
	tinConn.CreateTinConnection().Access(lib.Command.LOOK,"/warehouse","1234",lib.VERSION_1_0).Run()
}
