package jutil

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
)

func WriteJson(path string, data any, indent bool) (err error) {

	var buf *bytes.Buffer

	buf, err = Marshal(data, indent)
	if err != nil {
		return
	}

	err = os.WriteFile(path, buf.Bytes(), 0644)
	return
}

func DecodeObject(target any, object any) error {
	buf, err := Marshal(object, false)
	if err != nil {
		return err
	}
	return DecodeReader(bytes.NewBuffer(buf.Bytes()), target)
}

func DecodePath(path string, target any) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	return DecodeReadCloser(f, target)
}

func DecodeReadCloser(r io.ReadCloser, target any) error {
	defer func() { _ = r.Close() }()
	return DecodeReader(r, target)
}

func DecodeReader(r io.Reader, target any) error {
	return json.NewDecoder(r).Decode(target)
}

func Marshal(in any, indent bool) (buf *bytes.Buffer, err error) {
	var data []byte

	if indent {
		data, err = json.MarshalIndent(in, "", "  ")
	} else {
		data, err = json.Marshal(in)
	}
	if err != nil {
		return
	}

	buf = bytes.NewBuffer(data)
	return
}
