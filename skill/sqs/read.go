package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	jp "github.com/buger/jsonparser"
	"github.com/mlctrez/lexstream/internal/jutil"
	"log"
	"path/filepath"
	"time"
)

func main() {

	url := "https://sqs.us-east-1.amazonaws.com/359625541351/lexstream.fifo"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	client := sqs.NewFromConfig(cfg)
	input := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(url), MaxNumberOfMessages: 10, VisibilityTimeout: 0, WaitTimeSeconds: 20,
	}

	var sequence int

	for {
		receiveMessage, msgErr := client.ReceiveMessage(context.TODO(), input)
		if msgErr != nil {
			log.Println(msgErr)
			time.Sleep(30)
		}
		dmb := &sqs.DeleteMessageBatchInput{
			QueueUrl: aws.String(url),
			Entries:  []types.DeleteMessageBatchRequestEntry{},
		}

		for _, message := range receiveMessage.Messages {
			entry := types.DeleteMessageBatchRequestEntry{Id: message.MessageId, ReceiptHandle: message.ReceiptHandle}
			dmb.Entries = append(dmb.Entries, entry)

			messageBody := []byte(*message.Body)

			name, parseErr := jp.GetString(messageBody, "request", "header", "name")
			if parseErr != nil {
				log.Println(parseErr)
				continue
			}

			o := CreateOutput(messageBody)

			//os.WriteFile(fmt.Sprintf("temp/%02d.json", sequence), messageBody, 0644)
			fileName := fmt.Sprintf("%03d_%s.json", sequence, name)
			if mErr := jutil.WriteJson(filepath.Join("temp", fileName), o, true); mErr != nil {
				log.Println(mErr)
			}
			sequence++
		}
		if len(dmb.Entries) > 0 {
			if _, dmbErr := client.DeleteMessageBatch(context.TODO(), dmb); dmbErr != nil {
				log.Println(dmbErr)
			}
		}
	}

}

func CreateOutput(message json.RawMessage) *Output {
	requestHeader, _, _, _ := jp.Get(message, "request", "header")
	requestPayload, _, _, _ := jp.Get(message, "request", "bound_payload")
	responseHeader, _, _, _ := jp.Get(message, "response", "header")
	responsePayload, _, _, _ := jp.Get(message, "response", "payload")

	o := &Output{
		Request:  Body{Header: requestHeader, Payload: requestPayload},
		Response: Body{Header: responseHeader, Payload: responsePayload},
	}
	return o
}

type Output struct {
	Request  Body `json:"request"`
	Response Body `json:"response"`
}
type Body struct {
	Header  json.RawMessage `json:"header"`
	Payload json.RawMessage `json:"payload"`
}
