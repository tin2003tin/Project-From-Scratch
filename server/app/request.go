package app

import (
	"fmt"
	"net"
	tinPro "server/app/protocol"
)

type Request struct {
	Conn *net.Conn
	Protocol *tinPro.TinProtocol
	Params map[string]string
}

func (app *App)handleClient(conn net.Conn) {
    defer conn.Close()
	var protocol tinPro.TinProtocol;
	// Read Header
	err := tinPro.ReadHeader(&conn, &protocol)
	if (err != nil) {
		fmt.Println(err);
		return;
	}
	fmt.Println(protocol.Header)
	// Read Middleware
	err = tinPro.Middleware(&protocol.Header);
	if (err != nil) {
		fmt.Println(err);
		return
	}
	// Read Body
	err = tinPro.ReadBody(&conn, &protocol)
	if (err != nil) {
		fmt.Println(err);
		return;
	}
	// Read Tail
	err = tinPro.ReadTail(&conn, &protocol)
	if (err != nil) {
		fmt.Println(err);
		return;
	}

	// Send to route handler
	handler,params,err := app.handleCommand(&protocol);
	if (err != nil) {
		fmt.Println(err);
		return;
	}
	handler(Request{Conn: &conn,Protocol: &protocol, Params: params})
}

//Find handler Route and excute that handler
func (app *App)handleCommand(protocol *tinPro.TinProtocol) (func(Request), map[string]string, error) {
	index, params := app.findHandlerRoute(protocol.Header.Command, protocol.Header.Path)
	if index != -1 {
		return app.Routes[index].Handler, params, nil
	}
	return nil, nil, fmt.Errorf("handler not found for path: %s with command: %s", protocol.Header.Path, protocol.Header.Command)
}

