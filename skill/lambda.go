package skill

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/mlctrez/bolt"
	"github.com/mlctrez/lexstream/amsapi"
	"github.com/mlctrez/lexstream/amsapi/types"
	"go.etcd.io/bbolt"
	"os"
	"time"
)

func Handle(_ context.Context, request *amsapi.Request) (response interface{}, err error) {

	// testing bolt db access from within lambda execution
	var db *bolt.Bolt
	db, err = bolt.NewWithOptions("bolt.db", &bbolt.Options{ReadOnly: true, Timeout: 50 * time.Millisecond})
	if err != nil {
		return
	}
	defer func() { _ = db.Close() }()

	v := &bolt.Value{K: "keyOne"}
	err = db.Get("sample", v)
	if err != nil {
		return nil, err
	}
	fmt.Println("value from db is: ", string(v.V))

	request.LogPayload()

	switch request.Header.Namespace {
	case "Alexa.Media.Search":
		response, err = AlexaMediaSearch(request)
	case "Alexa.Media.Playback":
		response, err = AlexaMediaPlayback(request)
	case "Alexa.Audio.PlayQueue":
		response, err = AlexaAudioPlayQueue(request)
	default:
		return nil, fmt.Errorf("unhandled namespace %q", request.Header.Namespace)
	}

	return
}

func signedUrlToItem(item string, dur time.Duration) (url string, err error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return "", err
	}
	svc := s3.NewFromConfig(cfg)
	bucket := os.Getenv("LEXSTREAM_BUCKET_NAME")

	signer := s3.NewPresignClient(svc, s3.WithPresignExpires(dur))

	req, err := signer.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if err != nil {
		return "", err
	}
	return req.URL, nil
}

func AlexaAudioPlayQueue(request *amsapi.Request) (interface{}, error) {

	var mediaUrl string
	var imageUrl string
	var err error

	mediaUrl, err = signedUrlToItem("Undertale_OST__012_-_Home.ogg", 15*time.Minute)
	if err != nil {
		return nil, err
	}

	imageUrl, err = signedUrlToItem("undertale_soundtrack_cover.jpg", 15*time.Minute)
	if err != nil {
		return nil, err
	}
	m := &amsapi.Response{
		Header: amsapi.Header{
			Namespace:      request.Header.Namespace,
			Name:           request.Header.Name + ".Response",
			MessageId:      "resp-" + request.Header.MessageId,
			PayloadVersion: "1.0",
		},
	}

	item := &amsapi.Item{
		Id:           "track.001",
		PlaybackInfo: amsapi.PlaybackInfo{Type: "DEFAULT"},
		Metadata: amsapi.MediaMetadata{
			Type: types.TRACK,
			Name: amsapi.MetadataName{Speech: amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "home"}, Display: "home"},
			Art:  amsapi.Art{Sources: []amsapi.ArtSource{{Url: imageUrl}}},
			Authors: []amsapi.EntityMetadata{{Name: amsapi.MetadataNameProperty{
				Speech: amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "toby fox"}, Display: "toby fox",
			}}},
		},
		Controls: []amsapi.ItemControl{
			{Type: "COMMAND", Name: "NEXT", Enabled: true},
			{Type: "COMMAND", Name: "PREVIOUS", Enabled: true},
		},
		Rules:  amsapi.ItemRules{FeedbackEnabled: false},
		Stream: amsapi.Stream{Id: "track.001", Uri: mediaUrl, ValidUntil: time.Now().Add(10 * time.Minute).UTC()},
	}
	switch request.Header.Name {
	case "GetNextItem":
		m.Payload = &amsapi.GetNextItemResponse{IsQueueFinished: false, Item: item}
	case "GetPreviousItem":
		m.Payload = &amsapi.GetPreviousItemResponse{Item: item}
	default:
		return nil, fmt.Errorf("unhandled name %q", request.Header.Name)
	}
	return m, nil
}

func AlexaMediaPlayback(request *amsapi.Request) (interface{}, error) {
	var mediaUrl string
	var imageUrl string
	var err error

	mediaUrl, err = signedUrlToItem("Undertale_OST__012_-_Home.ogg", 15*time.Minute)
	if err != nil {
		return nil, err
	}

	imageUrl, err = signedUrlToItem("undertale_soundtrack_cover.jpg", 15*time.Minute)
	if err != nil {
		return nil, err
	}

	m := &amsapi.Response{
		Header: amsapi.Header{
			Namespace:      request.Header.Namespace,
			Name:           request.Header.Name + ".Response",
			MessageId:      "resp-" + request.Header.MessageId,
			PayloadVersion: "1.0",
		},
	}

	switch request.Header.Name {
	case "Initiate":
		m.Payload = &amsapi.InitiateResponse{
			PlaybackMethod: amsapi.PlaybackMethod{
				Type: "ALEXA_AUDIO_PLAYER_QUEUE",
				Id:   uuid.NewString(),
				Controls: []amsapi.QueueControl{
					{Type: "COMMAND", Name: "NEXT", Enabled: true},
					{Type: "COMMAND", Name: "PREVIOUS", Enabled: true},
				},
				Rules: amsapi.QueueRules{Feedback: amsapi.QueueFeedbackRule{Type: "PREFERENCE", Enabled: false}},
				FirstItem: amsapi.Item{
					Id:           "track.001",
					PlaybackInfo: amsapi.PlaybackInfo{Type: "DEFAULT"},
					Metadata: amsapi.MediaMetadata{
						Type: "TRACK",
						Name: amsapi.MetadataName{
							Speech: amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "home"}, Display: "Home"},
						Art: amsapi.Art{Sources: []amsapi.ArtSource{{Url: imageUrl}}},
						Authors: []amsapi.EntityMetadata{{Name: amsapi.MetadataNameProperty{
							Speech: amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "toby fox"}, Display: "Toby Fox",
						}}},
						Album: amsapi.EntityMetadata{},
					},
					Controls: []amsapi.ItemControl{
						{Type: "COMMAND", Name: "NEXT", Enabled: true},
						{Type: "COMMAND", Name: "PREVIOUS", Enabled: true},
					},
					Rules:  amsapi.ItemRules{FeedbackEnabled: false},
					Stream: amsapi.Stream{Id: "track.001", Uri: mediaUrl, ValidUntil: time.Now().Add(10 * time.Minute).UTC()},
				},
			},
		}
		m.LogPayload()
		return m, nil
	default:
		return nil, fmt.Errorf("unhandled name %q", request.Header.Name)
	}

}

func AlexaMediaSearch(request *amsapi.Request) (interface{}, error) {
	switch request.Header.Name {
	case "GetPlayableContent":
		gpc := &amsapi.GetPlayableContent{}
		err := request.BindPayload(gpc)
		if err != nil {
			return nil, err
		}
		response := &amsapi.Response{
			Header: amsapi.Header{
				Namespace:      "Alexa.Media.Search",
				Name:           "GetPlayableContent.Response",
				MessageId:      "resp-" + request.Header.MessageId,
				PayloadVersion: "1.0",
			},
			Payload: amsapi.GetPlayableContentResponse{
				Content: amsapi.Content{
					// todo: support more than just the one track
					Id:      "track.001",
					Actions: amsapi.ContentActions{Playable: true, Browsable: false},
					Metadata: amsapi.MediaMetadata{
						Type: "TRACK",
						Name: amsapi.MetadataName{
							Speech: amsapi.SpeechInfo{
								Type: "PLAIN_TEXT",
								Text: "home",
							},
							Display: "Home",
						},
						Authors: []amsapi.EntityMetadata{
							{Name: amsapi.MetadataNameProperty{
								Speech: amsapi.SpeechInfo{
									Type: "PLAIN_TEXT",
									Text: "toby fox",
								},
								Display: "Toby Fox",
							}},
						},
						Album: amsapi.EntityMetadata{Name: amsapi.MetadataNameProperty{
							Speech: amsapi.SpeechInfo{
								Type: "PLAIN_TEXT",
								Text: "under tale soundtrack",
							},
							Display: "Undertale Soundtrack",
						}},
						Art: amsapi.Art{Sources: []amsapi.ArtSource{
							{Url: "https://mlctrez-website.s3.amazonaws.com/undertale_soundtrack_cover.jpg"},
						}},
					},
				},
			},
		}
		response.LogPayload()
		return response, nil

	}
	return nil, fmt.Errorf("unhandled name %q", request.Header.Name)
}
