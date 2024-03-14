package tinConn

import (
	tinPro "client/tinConn/protocol"
	"net"
)

type Response struct {
	protocol tinPro.TinResProtocol
}

func (r *Response) GetResponse() *tinPro.TpResHeader {
	return r.protocol.GetHeader();
}
func (r *Response) GetData() *tinPro.TpResBody {
	return r.protocol.GetBody();
}

func (r *Response) readReponse(conn net.Conn) (error) {
	defer conn.Close()
	// Read Header
	err := tinPro.ReadHeader(conn, &r.protocol)
	if (err != nil) {
		return err
	}
	// Read Body 
	err = tinPro.ReadBody(conn, &r.protocol)
	if (err != nil) {
		return err
	}
	// Read Tail 
	err = tinPro.ReadTail(conn, &r.protocol)
	if (err != nil) {
		return err
	}
	return nil;
}