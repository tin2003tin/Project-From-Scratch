package tinPro

import (
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
    Data interface{}      `json:"Data"`
}
type TpReqTail struct {
	Message string        `json:"Message"`
	Description string    `json:"Description"`
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

