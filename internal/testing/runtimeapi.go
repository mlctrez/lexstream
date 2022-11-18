package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/mlctrez/lexstream/amsapi"
	"github.com/mlctrez/lexstream/amsapi/header"
	"github.com/mlctrez/lexstream/internal/jutil"
	"net/http"
	"time"
)

func main() {

	// These environment variables need to be set when invoking the lambda
	// os.Setenv("LEXSTREAM_BUCKET_NAME", "xxxxx")
	// os.Setenv("AWS_LAMBDA_RUNTIME_API", "localhost:9990")

	http.HandleFunc("/2018-06-01/runtime/invocation/next", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL.Path)
		time.Sleep(4 * time.Second)
		writer.Header().Set("Lambda-Runtime-Aws-Request-Id", uuid.NewString())
		writer.Header().Set("Lambda-Runtime-Deadline-Ms", fmt.Sprintf("%d", time.Now().Add(time.Millisecond*500).UnixMilli()))

		payload := json.RawMessage("{}")

		req := &amsapi.Request{
			Header: &header.Header{
				Namespace:      header.AlexaMediaSearch,
				Name:           header.GetPlayableContent,
				MessageId:      uuid.NewString(),
				PayloadVersion: "1.0",
			},
			Payload: &payload,
		}
		json.NewEncoder(writer).Encode(req)
	})
	http.HandleFunc("/2018-06-01/runtime/invocation/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method, request.URL.Path)
		m := make(map[string]any)
		json.NewDecoder(request.Body).Decode(&m)
		buf, _ := jutil.Marshal(m, true)
		fmt.Println(buf.String())
		marshal, _ := jutil.Marshal(StatusResponse{Status: "OK"}, false)
		writer.WriteHeader(202)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(marshal.Bytes())

	})

	err := http.ListenAndServe(":9990", nil)
	if err != nil {
		panic(err)
	}

}

type StatusResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorType    string `json:"errorType"`
}
