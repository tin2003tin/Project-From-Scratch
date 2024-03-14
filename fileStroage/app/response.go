package app

import (
	"fmt"
	"net"
	tinPro "system/app/protocol"
)

const (
	SuccessStatusCode = 200
	ErrorStatusCode   = 400
)

type Response struct {
	Conn *net.Conn
	ReqProtocol *tinPro.TinReqProtocol
	ResProtocol *tinPro.TinResProtocol
}
func (res *Response) SetMessage(message string) (*Response)  {
	res.ResProtocol.SetHeader(*res.ReqProtocol,SuccessStatusCode,message,"text",0,0);
	return res;
}

func (res *Response) SetBody(data interface{}) (*Response)  {
	res.ResProtocol.SetBody(data)
	return res;
}

func ErrorToClient(conn net.Conn, protocol *tinPro.TinReqProtocol, err error) error {
	if (err != nil) {
		tr := tinPro.TinResProtocol{}
		tr.SetHeader(*protocol,ErrorStatusCode,err.Error(),"Error",0,0);
		//Marshal the Protocol
		JsonHeader, JsonBody, JsonTail, err := tinPro.MarshalProtocol(&tr);
		if (err != nil) {
			return err
		}
		//Write the connection
		_, err = conn.Write(append(JsonHeader, '\n'))
		if err != nil {
			return err
		}
		_, err = conn.Write(JsonBody)
		if err != nil {
			return err
		}
		_, err = conn.Write(JsonTail)
		if err != nil {
			return err
		}
	}
	return nil;
}

func (res *Response) Send() error {
	//Marshal the Protocol
	JsonHeader, JsonBody, JsonTail, err := tinPro.MarshalProtocol(res.ResProtocol);
	if (err != nil) {
		return err
	}

	fmt.Println("try to send to response...")

	//Write the connection
	_, err = (*res.Conn).Write(append(JsonHeader, '\n'))
	if err != nil {
		return err
	}

	_, err = (*res.Conn).Write(JsonBody)
	if err != nil {
		return err
	}

	_, err = (*res.Conn).Write(JsonTail)
	if err != nil {
		return err
	}
	return nil
}