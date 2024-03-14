package tiny

import "server/app"

func TinSayHello(req app.Request, res app.Response) error {
	res.SetMessage("Hello World").Send()	
	return nil
}
