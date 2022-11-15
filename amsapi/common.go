package amsapi

import (
	"encoding/json"
	"fmt"
	"github.com/mlctrez/lexstream/amsapi/header"
)

type Request struct {
	Header       *header.Header   `json:"header"`
	Payload      *json.RawMessage `json:"payload,omitempty"`
	BoundPayload any              `json:"bound_payload,omitempty"`
}

// Bind unmarshalls the request payload into p
func Bind[K any](j *Request, value K) K {
	if j.Payload == nil {
		panic(fmt.Errorf("no payload to bind"))
	}
	marshalJSON, err := j.Payload.MarshalJSON()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(marshalJSON, value)
	if err != nil {
		panic(err)
	}
	j.BoundPayload = value
	j.Payload = nil
	return value
}

type Response struct {
	Header  *header.Header `json:"header"`
	Payload any            `json:"payload"`
}
