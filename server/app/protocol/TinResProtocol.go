package tinPro

import (
	"time"
)

type TinResProtocol struct {
	header TpResHeader
	body   TpResBody
	tail   TpResTail
}

type TpResHeader struct {
	StatusCode int 	 	 `json:"StatusCode"`
	Message    string 	 `json:"Message"`
	Command    string    `json:"Command"`
	Path       string    `json:"Path"`
	Version    string    `json:"Version"`
	Date       time.Time `json:"Date"`
	BodyType   string    `json:"BodyType"`
	BodyLength int64     `json:"BodyLength"`
	TailLength int64     `json:"TailLength"`
}
type TpResBody struct {
	Data interface{}      `json:"Data"`
}
type TpResTail struct {
	Message     string    `json:"Message"`
	Description string    `json:"Description"`
}

func (trp *TinResProtocol) GetHeader() *TpResHeader {
	return &trp.header
}

func (trp *TinResProtocol) GetBody() *TpResBody {
	return &trp.body
}

func (trp *TinResProtocol) GetTail() *TpResTail {
	return &trp.tail
}

func (tr *TinResProtocol) SetHeader(protocol TinReqProtocol,statusCode int,message string,bodyType string, bodyLength int64, tailLength int64) {
	tr.GetHeader().StatusCode = statusCode;
	tr.GetHeader().Message = message
	tr.GetHeader().Command = protocol.GetHeader().Command 
	tr.GetHeader().Path = protocol.GetHeader().Path 
	tr.GetHeader().Version = protocol.GetHeader().Version
	tr.GetHeader().Date = time.Now() 
	tr.GetHeader().BodyType = bodyType
	tr.GetHeader().BodyLength = bodyLength
	tr.GetHeader().TailLength = tailLength
}

func (tr *TinResProtocol) SetBody(data interface{}) {
	tr.body.Data = data
}
