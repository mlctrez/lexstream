package header

import "github.com/google/uuid"

type Header struct {
	Namespace      Namespace `json:"namespace"`
	Name           Name      `json:"name"`
	MessageId      string    `json:"messageId"`
	PayloadVersion string    `json:"payloadVersion"`
}

type Namespace string

const AlexaAudioPlayQueue Namespace = "Alexa.Audio.PlayQueue"
const AlexaMediaPlayback Namespace = "Alexa.Media.Playback"
const AlexaMediaPlayQueue Namespace = "Alexa.Media.PlayQueue"
const AlexaMediaSearch Namespace = "Alexa.Media.Search"
const AlexaUserPreference Namespace = "Alexa.UserPreference"

type Name string

const Initiate Name = "Initiate"
const InitiateResponse Name = Initiate + ".Response"
const Reinitiate Name = "Reinitiate"
const ReinitiateResponse Name = Reinitiate + ".Response"

const GetNextItem Name = "GetNextItem"
const GetNextItemResponse Name = GetNextItem + ".Response"
const GetPreviousItem Name = "GetPreviousItem"
const GetPreviousItemResponse Name = GetPreviousItem + ".Response"
const GetPlayableContent Name = "GetPlayableContent"
const GetPlayableContentResponse Name = GetPlayableContent + ".Response"

func uu() string { return uuid.NewString() }

const V1 = "1.0"

func header(ns Namespace, name Name) *Header {
	return &Header{Namespace: ns, Name: name, MessageId: uu(), PayloadVersion: V1}
}

func InitiateResponseHeader() *Header {
	return header(AlexaMediaPlayback, InitiateResponse)
}

func ReinitiateResponseHeader() *Header {
	return header(AlexaMediaPlayback, ReinitiateResponse)
}

func GetNextItemResponseHeader() *Header {
	return header(AlexaAudioPlayQueue, GetNextItemResponse)
}

func GetPreviousItemResponseHeader() *Header {
	return header(AlexaAudioPlayQueue, GetPreviousItemResponse)
}

func GetPlayableContentResponseHeader() *Header {
	return header(AlexaMediaSearch, GetPlayableContentResponse)
}
