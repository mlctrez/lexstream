package amsapi

import "github.com/mlctrez/lexstream/amsapi/types"

// GetPlayableContent supports the GetPlayableContent directive so that customers can request content.
//
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-media-search.html#getplayablecontent-directive
type GetPlayableContent struct {
	RequestContext    RequestContext            `json:"requestContext"`
	Filters           Filter                    `json:"filters"`
	SelectionCriteria ResolvedSelectionCriteria `json:"selectionCriteria"`
}

// RequestContext describes the context of an Alexa Music Skill API request.
// If the music service provider supports premium audio in their skill, requests for GetItem, GetNextItem,
// GetPreviousItem, and JumpToItem contain an Endpoint object in the RequestContext containing information
// about what types of audio quality streams are supported on the target device.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#requestcontext
type RequestContext struct {
	User     User     `json:"user"`
	Location Location `json:"location,omitempty"`
}

// User describes information about the user who initiated a request to the music skill.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#user
type User struct {
	Id          string `json:"id"`
	AccessToken string `json:"accessToken,omitempty"`
}

// Location describes the location of a request.
// Use this information to enforce geographical restrictions on content or decide which language
// version of metadata to return to Alexa for media items.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#location
type Location struct {
	OriginatingLocale string `json:"originatingLocale,omitempty"`
	CountryCode       string `json:"countryCode,omitempty"`
}

// Filter describes filters that the skill should apply to search results (selection criteria and content) before
// returning a response to the Alexa service.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#filters
type Filter struct {
	ExplicitLanguageAllowed bool `json:"explicitLanguageAllowed"`
}

// ResolvedSelectionCriteria is content selection criteria in resolved form.
// A ResolvedSelectionCriteria object identifies attributes that can be resolved (by searching) to a Content object.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#resolvedselectioncriteria
type ResolvedSelectionCriteria struct {
	Attributes []ResolvedSelectionCriteriaAttribute `json:"attributes"`
}

// ResolvedSelectionCriteriaAttribute is a single attribute within a resolved content selection criteria.
//
// A ResolvedSelectionCriteriaAttribute object conveys a single aspect of a user's request in refined
// or processed form, such as the song name "Thrift Shop" (with skill-provided entitytype ID) in Alexa, play
// the song Thrift Shop.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#resolvedselectioncriteriaattribute
type ResolvedSelectionCriteriaAttribute struct {
	Type     types.EntityType `json:"type"`
	EntityId string           `json:"entityId,omitempty"`
	Value    string           `json:"value,omitempty"`
}

// GetPlayableContentResponse is the reply to GetPlayableContent.
// If you handle a GetPlayableContent directive successfully, respond with an Alexa.Response event.
// If your skill can successfully find playable content to satisfy the request, it should respond
// with GetPlayableContentResponse.
type GetPlayableContentResponse struct {
	Content Content `json:"content"`
}

// Content includes an identifier (reference) and metadata (for example, album name, author name, title, type, etc.).
//
// The content object can represent a single track, an album, a playlist, a custom station, a live station, or a
// program. content objects are global and can be shared between different users.
//
// For example, if content object 1234 is the album "King Animal" by Soundgarden, this is true for any user who
// receives content object 1234, and if a user receives content object 1234 a year later, it should still represent
// the album King Animal by Soundgarden.
//
// The content object is the primary object for the Alexa.Media.Search interface.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#content
type Content struct {
	Id       string         `json:"id"`
	Actions  ContentActions `json:"actions"`
	Metadata MediaMetadata  `json:"metadata"`
}

// ContentActions describes whether a content object represents content that's playable, browsable, or both.
//
// The flags in this object are intended to support display use cases where the user can click a content object
// to play all its contents (for example, play an entire album) or the user can click a content object to view
// the items within (for example, see the tracks listing for the album) and perhaps play one item.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#contentactions
type ContentActions struct {
	Playable  bool `json:"playable"`
	Browsable bool `json:"browsable"`
}


