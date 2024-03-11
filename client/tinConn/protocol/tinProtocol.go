package tinPro

import (
	"client/tinConn/lib"
	"time"
)

type TinProtocol struct {
	Header tpHeader
	Body   tpBody
	Tail   tpTail
}

type tpHeader struct {
	Command       string
	Path          string
	Version       string
	
	SecretKey     string
	MasterKey     string

	Date          time.Time

	BodyType      string
	BodyLength    int64

	TailLength    int64

}

type tpBody struct {
	Data interface{}
}

type tpTail struct {
	Message string
	Description string
}

func CreateDefaultTinProtocol() *TinProtocol {
    p := &TinProtocol{}
	p.SetAccess(lib.Command.LOOK, "/", "", lib.VERSION_1_0);
    return p
}

func CreateCustomTinProtocol(command, path, secretKey, version string) *TinProtocol {
    p := &TinProtocol{}
    p.SetAccess(command, path, secretKey, version)
    return p
}

func (p *TinProtocol) SetAccess(command, path, secretKey, version string) {
    p.Header = tpHeader{
        Command:   command,
        Path:      path,
        SecretKey: secretKey,
        Version:   version,
    }
}

func (p *TinProtocol) SetBody(data interface{}) {
	p.Body.Data = data
}

func (p *TinProtocol) SetTail(message string,description   string) {
	p.Tail.Description = description
	p.Tail.Message = message
}

func (p *TinProtocol) IsValid() bool {
	switch p.Header.Command {
	case lib.Command.LOOK, lib.Command.SEND, lib.Command.EDIT, lib.Command.TINY:
		return true
	default:
		return false
	}
}