package warehouse

import (
	"errors"
	"server/app"
	"server/tinConn"
)

func DestroyFile(req app.Request, res app.Response) error {
	tc := tinConn.CreateTinConnection("8000");
	tc.Access(req.Protocol.GetHeader().Command,req.Protocol.GetHeader().Path,req.Protocol.GetHeader().SecretKey,req.Protocol.GetHeader().Version)
    response := tc.Run()
	if (response.GetResponse().StatusCode == 400) {
		app.ErrorToClient(*res.Conn, res.ReqProtocol, errors.New(response.GetResponse().Message))
		return nil
	}
	err := res.SetMessage(response.GetResponse().Message).SetBody(response.GetData().Data).Send()
	if (err != nil) {
		return err
	}
	return nil;
}