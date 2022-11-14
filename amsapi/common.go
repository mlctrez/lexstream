package amsapi

import (
	"encoding/json"
	"fmt"
)

type Header struct {
	Namespace      string `json:"namespace"`
	Name           string `json:"name"`
	MessageId      string `json:"messageId"`
	PayloadVersion string `json:"payloadVersion"`
}

type Request struct {
	Header  Header      `json:"header"`
	Payload interface{} `json:"payload"`
}

func (m *Request) BindPayload(p interface{}) (err error) {
	var marshal []byte
	if marshal, err = json.Marshal(m.Payload); err != nil {
		return
	}
	return json.Unmarshal(marshal, p)
}

func (m *Request) LogPayload() {
	if m == nil {
		fmt.Println("Request.LogPayload with nil *Request")
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Request.LogPayload marshal error", err)
		return
	}
	fmt.Println("Request.LogPayload", string(marshal))
}

type Response struct {
	Header  *Header     `json:"header"`
	Payload interface{} `json:"payload"`
}

func (m *Response) LogPayload() {
	if m == nil {
		fmt.Println("Response.LogPayload with nil *Response")
	}
	marshal, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Response.LogPayload marshal error", err)
		return
	}
	fmt.Println("Response.LogPayload", string(marshal))
}
