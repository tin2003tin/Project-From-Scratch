package main

import (
	"client/tinConn"
	"client/tinConn/lib"
)

func main() {
	tinConn.CreateTinConnection().Access(lib.Command.LOOK,"/warehouse","1234",lib.VERSION_1_0).Run()
	// conn, err := net.Dial("tcp", "localhost:1000")
	// if err != nil {
	// 	fmt.Println("Error connecting:", err)
	// 	return
	// }
	// JsonBody, err := json.Marshal(interface{}(tinCreate.Importfile("object/test.txt")))
	// _, err = conn.Write(JsonBody)
	// if err != nil {
	// 	fmt.Println("Error writing to connection:", err)
	// 	return
	// }
}
