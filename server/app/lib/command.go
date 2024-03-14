package lib

var Command = struct {
	LOOK string
	SEND string
	KILL string
	EDIT string
	TINY string
}{
	LOOK: "LOOK",
	SEND: "SEND",
	KILL: "KILL",
	EDIT: "EDIT",
	TINY: "TINY",
}

const (
	VERSION_1_0 = "TinProtocolV1.0"
)