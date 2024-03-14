package tinPro

import "encoding/json"

func MarshalProtocol(protocol *TinResProtocol) ([]byte, []byte, []byte, error) {
	JsonTail, err := marshalJsonTail(protocol)
	if err != nil {
		return nil, nil, nil, err
	}

	JsonBody, err := marshalJsonBody(protocol)
	if err != nil {
		return nil, nil, nil, err
	}

	JsonHeader, err := marshalJsonHeader(protocol)
	if err != nil {
		return nil, nil, nil, err
	}
	return JsonHeader, JsonBody, JsonTail, nil
}

func marshalJsonHeader(protocol *TinResProtocol) ([]byte, error) {
	JsonHeader, err := json.Marshal(protocol.GetHeader())
	if err != nil {
		return nil, err
	}
	return JsonHeader, nil
}

func marshalJsonBody(protocol *TinResProtocol) ([]byte, error) {
	JsonBody, err := json.Marshal(protocol.GetBody())
	if err != nil {
		return nil, err
	}
	if len(JsonBody) != 0 {

		protocol.GetHeader().BodyLength = int64(len(JsonBody))
		protocol.GetHeader().BodyType = "text"

	}
	return JsonBody, nil
}

func marshalJsonTail(protocol *TinResProtocol) ([]byte, error) {
	JsonTail, err := json.Marshal(protocol.GetTail())
	if err != nil {
		return nil, err
	}
	if len(JsonTail) != 0 {

		protocol.GetHeader().TailLength = int64(len(JsonTail))
	}
	return JsonTail, nil
}