package main

import (
	tinConn "client/tinConn"
	"client/tinConn/lib"
	"fmt"
)

func main() {
	tc := tinConn.CreateTinConnection("8080")
	tc.Access(lib.Command.TINY, "/tin", "1234", lib.VERSION_1_0)
	// tc.Body(tinCreate.Importfile("test_object/vocabulary.txt"))
	tc.Error(handleError)
	response := tc.Run()
	fmt.Println(response.GetResponse().Message);
	fmt.Println(response.GetData())
}

func handleError(err error) {
	fmt.Println("Test", err)
}
