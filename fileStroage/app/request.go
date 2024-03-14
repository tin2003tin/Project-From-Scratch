package app

import (
	"fmt"
	"net"
	"system/app/lib"
	tinPro "system/app/protocol"
)

type Request struct {
	Conn *net.Conn
	Protocol *tinPro.TinReqProtocol
	Params map[string]string
}

func (app *App) handleClient(conn net.Conn) {
	var protocol tinPro.TinReqProtocol;
	//Read Header
	err := tinPro.ReadHeader(conn, &protocol)
	if (err != nil) {
		err = ErrorToClient(conn,&protocol,err)
		if (err != nil) {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(protocol.GetHeader())
	//Read Header && Tail
	if (protocol.GetHeader().Command == lib.Command.SEND) {
		tinPro.ReadBody(conn, &protocol)
		tinPro.ReadTail(conn, &protocol)
	}
	fmt.Println(protocol.GetBody())
	//Send to route handler
	handler,params,err := app.handleCommand(&protocol);
	if (err != nil) {
		err = ErrorToClient(conn,&protocol,err)
		if (err != nil) {
			fmt.Println(err)
		}
		return
	}
	err =  handler(Request{Conn: &conn,Protocol: &protocol, Params: params}, 
		           Response{Conn: &conn,ReqProtocol: &protocol,ResProtocol: &tinPro.TinResProtocol{}})
	if (err != nil) {
		err = ErrorToClient(conn, &protocol,err)
		if (err != nil) {
			fmt.Println(err)
		}
		return
	}
}
