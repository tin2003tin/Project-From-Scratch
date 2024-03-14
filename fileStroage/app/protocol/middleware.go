package tinPro

import "fmt"

func Middleware(header *TpReqHeader) error {
	if header.Command == "" {
		return fmt.Errorf("the command is required")
	}
	if header.SecretKey != "1234" {
		return fmt.Errorf("the SecretKey is incorrect")
	}
	if header.Command == "TINY" && header.MasterKey == "master" {
		return fmt.Errorf("the TINY command need correctly masterKey")
	}
	return nil
}
