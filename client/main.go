package main

import (
	"client/tinConn"
	tinCreate "client/tinConn/create"
	"client/tinConn/lib"
)

func main() {
	tinConn.CreateTinConnection("8080").Access(lib.Command.SEND,"/warehouse","1234",lib.VERSION_1_0).Body(tinCreate.Importfile("C:/Users/com/Downloads/PlatformerTutorial-ep28_finale.zip")).Run()
}
