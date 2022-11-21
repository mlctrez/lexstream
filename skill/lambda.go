package skill

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/google/uuid"
	"github.com/mlctrez/bolt"
	"github.com/mlctrez/lexstream/amsapi"
	"github.com/mlctrez/lexstream/amsapi/header"
	errorResponse "github.com/mlctrez/lexstream/amsapi/response"
	"github.com/mlctrez/lexstream/amsapi/types"
	"github.com/mlctrez/lexstream/internal/jutil"
	"go.etcd.io/bbolt"
	"os"
	"time"
)

func Handle(_ context.Context, request *amsapi.Request) (res *amsapi.Response, err error) {

	if request == nil || request.Header == nil || request.Header.Namespace == "" {
		res = internalError(fmt.Errorf("missing request headers"))
		return
	}

	defer func() {
		if pe, ok := recover().(error); ok {
			res = internalError(pe)
		}
	}()

	switch request.Header.Namespace {

	case header.AlexaAudioPlayQueue:

		switch request.Header.Name {
		case header.GetNextItem:
			res = GetNextItem(amsapi.Bind(request, &amsapi.GetNextItem{}))
		case header.GetPreviousItem:
			res = GetPreviousItem(amsapi.Bind(request, &amsapi.GetPreviousItem{}))
		default:
			res = invalidName(request.Header.Name)
		}

	case header.AlexaMediaPlayback:

		switch request.Header.Name {
		case header.Initiate:
			res = Initiate(amsapi.Bind(request, &amsapi.Initiate{}))
		case header.Reinitiate:
			// TODO: Reinitiate directive handling
			res = invalidName(request.Header.Name)
		default:
			res = invalidName(request.Header.Name)
		}

	case header.AlexaMediaPlayQueue:
		// TODO: implement
	case header.AlexaMediaSearch:
		res = GetPlayableContent(amsapi.Bind(request, &amsapi.GetPlayableContent{}))
	case header.AlexaUserPreference:
		// TODO: implement
	}

	// no res created, return invalid directive error res
	if res == nil {
		res = invalidNamespace(request.Header.Namespace)
	}

	if os.Getenv("LEXSTREAM_QUEUE_NAME") != "" {
		pushPayloadToSQS(map[string]any{"request": request, "response": res}, request.Header.MessageId)
	}

	return
}

func GetPlayableContent(payload *amsapi.GetPlayableContent) (response *amsapi.Response) {

	_ = payload

	_, imageUrl, errResponse := media("")
	if errResponse != nil {
		return errResponse
	}

	responsePayload := amsapi.GetPlayableContentResponse{
		Content: amsapi.Content{
			Id:      "track.001",
			Actions: amsapi.ContentActions{Playable: true, Browsable: false},
			Metadata: amsapi.MediaMetadata{
				Type: "TRACK",
				Name: amsapi.MetadataName{
					Speech: amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "home"}, Display: "Home"},
				Authors: []amsapi.EntityMetadata{
					{Name: amsapi.MetadataNameProperty{
						Speech: amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "toby fox"}, Display: "Toby Fox"},
					},
				},
				Album: amsapi.EntityMetadata{Name: amsapi.MetadataNameProperty{
					Speech:  amsapi.SpeechInfo{Type: "PLAIN_TEXT", Text: "under tale soundtrack"},
					Display: "Undertale Soundtrack",
				}},
				Art: amsapi.Art{Sources: []amsapi.ArtSource{{Url: imageUrl}}},
			},
		},
	}
	response = &amsapi.Response{
		Header:  header.GetPlayableContentResponseHeader(),
		Payload: responsePayload,
	}
	return
}

func Initiate(payload *amsapi.Initiate) (response *amsapi.Response) {
	_ = payload

	mediaUrl, imageUrl, errResponse := media("")
	if errResponse != nil {
		return errResponse
	}

	trackID := "track.001"

	responsePayload := &amsapi.InitiateResponse{
		PlaybackMethod: amsapi.PlaybackMethod{
			Type: amsapi.PlaybackMethodType,
			Id:   uuid.NewString(),
			Controls: []amsapi.QueueControl{
				{Type: "COMMAND", Name: "NEXT", Enabled: true},
				{Type: "COMMAND", Name: "PREVIOUS", Enabled: true},
			},
			Rules: amsapi.QueueRules{Feedback: amsapi.QueueFeedbackRule{Type: "PREFERENCE", Enabled: false}},
			FirstItem: amsapi.Item{
				Id:           trackID,
				PlaybackInfo: amsapi.PlaybackInfo{Type: amsapi.PlaybackInfoTypeDefault},
				Metadata: amsapi.MediaMetadata{
					Type: types.TRACK,

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
				Stream: amsapi.Stream{Id: trackID, Uri: mediaUrl, ValidUntil: time.Now().Add(10 * time.Minute).UTC()},
			},
		},
	}
	response = &amsapi.Response{
		Header:  header.InitiateResponseHeader(),
		Payload: responsePayload,
	}

	return
}

func GetNextItem(payload *amsapi.GetNextItem) (response *amsapi.Response) {

	_ = payload

	mediaUrl, imageUrl, errResponse := media("")
	if errResponse != nil {
		return errResponse
	}

	trackID := "track.001"
	streamValidUntil := time.Now().Add(10 * time.Minute).UTC()

	item := &amsapi.Item{
		Id:           trackID,
		PlaybackInfo: amsapi.PlaybackInfo{Type: amsapi.PlaybackInfoTypeDefault},
		Metadata: amsapi.MediaMetadata{
			Type:    types.TRACK,
			Name:    amsapi.BuildMetadataName("home", "Home"),
			Art:     amsapi.Art{Sources: []amsapi.ArtSource{{Url: imageUrl}}},
			Authors: amsapi.BuildAuthors("toby fox", "Toby Fox"),
		},
		Controls: []amsapi.ItemControl{
			amsapi.BuildItemControlNext(true),
			amsapi.BuildItemControlPrevious(true),
		},
		Rules:  amsapi.ItemRules{FeedbackEnabled: false},
		Stream: amsapi.Stream{Id: trackID, Uri: mediaUrl, ValidUntil: streamValidUntil},
	}

	response = &amsapi.Response{
		Header:  header.GetNextItemResponseHeader(),
		Payload: &amsapi.GetNextItemResponse{IsQueueFinished: false, Item: item},
	}

	return
}

func GetPreviousItem(payload *amsapi.GetPreviousItem) (response *amsapi.Response) {

	_ = payload

	mediaUrl, imageUrl, errResponse := media("")
	if errResponse != nil {
		return errResponse
	}

	trackID := "track.001"
	streamValidUntil := time.Now().Add(10 * time.Minute).UTC()
	item := &amsapi.Item{
		Id:           trackID,
		PlaybackInfo: amsapi.PlaybackInfo{Type: "DEFAULT"},
		Metadata: amsapi.MediaMetadata{
			Type:    types.TRACK,
			Name:    amsapi.BuildMetadataName("home", "Home"),
			Art:     amsapi.Art{Sources: []amsapi.ArtSource{{Url: imageUrl}}},
			Authors: amsapi.BuildAuthors("toby fox", "Toby Fox"),
		},
		Controls: []amsapi.ItemControl{
			amsapi.BuildItemControlPrevious(true),
			amsapi.BuildItemControlNext(true),
		},
		Rules:  amsapi.ItemRules{FeedbackEnabled: false},
		Stream: amsapi.Stream{Id: trackID, Uri: mediaUrl, ValidUntil: streamValidUntil},
	}

	response = &amsapi.Response{
		Header:  header.GetPreviousItemResponseHeader(),
		Payload: &amsapi.GetPreviousItemResponse{Item: item},
	}

	return
}

func media(id string) (mediaUrl, imageUrl string, errResponse *amsapi.Response) {

	_ = id

	var err error

	if mediaUrl, err = signedUrlToItem("Undertale_OST__012_-_Home.ogg", 15*time.Minute); err != nil {
		errResponse = internalError(err)
		return
	}

	if imageUrl, err = signedUrlToItem("undertale_soundtrack_cover.jpg", 15*time.Minute); err != nil {
		errResponse = internalError(err)
		return
	}
	return
}

var _ = boltTest

func boltTest() (err error) {
	// testing bolt db access from within lambda execution
	options := &bbolt.Options{ReadOnly: true, Timeout: 50 * time.Millisecond}

	var db *bolt.Bolt
	if db, err = bolt.NewWithOptions("bolt.db", options); err != nil {
		return
	}
	defer func() { _ = db.Close() }()

	v := &bolt.Value{K: "keyOne"}
	if err = db.Get("sample", v); err != nil {
		return
	}
	fmt.Println("value from db is: ", string(v.V))
	return
}

func internalError(err error) *amsapi.Response {
	return errorResponse.AlexaError(types.ErrorInternalError, uuid.NewString(), err.Error())
}

func pushPayloadToSQS(payload map[string]any, messageId string) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	buf, err := jutil.Marshal(payload, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	messageInput := &sqs.SendMessageInput{
		MessageBody:            aws.String(buf.String()),
		QueueUrl:               aws.String(os.Getenv("LEXSTREAM_QUEUE_NAME")),
		DelaySeconds:           0,
		MessageGroupId:         aws.String("lexstream"),
		MessageDeduplicationId: aws.String(messageId),
	}
	client := sqs.NewFromConfig(cfg)
	_, err = client.SendMessage(context.TODO(), messageInput)
	if err != nil {
		fmt.Println(err)
		return
	}
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

func invalidName(name header.Name) *amsapi.Response {
	message := fmt.Sprintf("unhandled header.name %q", name)
	return errorResponse.AlexaError(types.ErrorInvalidDirective, uuid.NewString(), message)
}

func invalidNamespace(name header.Namespace) *amsapi.Response {
	message := fmt.Sprintf("unhandled header.namespace %q", name)
	return errorResponse.AlexaError(types.ErrorInvalidDirective, uuid.NewString(), message)
}
