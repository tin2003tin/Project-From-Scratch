package tinPro

import (
	"errors"
	"client/tinConn/lib"
	"time"
)

type TinReqProtocol struct {
    header TpReqHeader 
    body   TpReqBody  
    tail   TpReqTail   
}

type TpReqHeader struct {
    Command    string    `json:"Command"`
    Path       string    `json:"Path"`
    Version    string    `json:"Version"`
    SecretKey  string    `json:"SecretKey"`
    MasterKey  string    `json:"MasterKey"`
    Date       time.Time `json:"Date"`
    BodyType   string    `json:"BodyType"`
    BodyLength int64     `json:"BodyLength"`
    TailLength int64     `json:"TailLength"`
}

type TpReqBody struct {
    Data interface{} `json:"Data"`
}

type TpReqTail struct {
    Message     string `json:"Message"`
    Description string `json:"Description"`
}

func (trp *TinReqProtocol) GetHeader() *TpReqHeader {
	return &trp.header
}

func (trp *TinReqProtocol) GetBody() *TpReqBody {
	return &trp.body
}

func (trp *TinReqProtocol) GetTail() *TpReqTail {
	return &trp.tail
}

func (p *TinReqProtocol) SetAccess(command, path, secretKey, version string) {
    p.header = TpReqHeader{
        Command:   command,
        Path:      path,
        SecretKey: secretKey,
        Version:   version,
    }
}

func (p *TinReqProtocol) SetBody(data interface{}) {
	p.body.Data = data
}

func (p *TinReqProtocol) SetTail(message string,description   string) {
	p.tail.Description = description
	p.tail.Message = message
}

func (p *TinReqProtocol) IsValid() error {
	switch p.header.Command {
	case lib.Command.LOOK, lib.Command.SEND, lib.Command.EDIT, lib.Command.TINY, lib.Command.KILL:
		return nil
	default:
		return errors.New("invalid command")
	}
}