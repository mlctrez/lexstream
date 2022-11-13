package amsapi

import "time"

// Initiate payloads are sent when Alexa receives a content identifier from a skill's GetPlayableContent
// response.
//
// The skill responds with the Stream URI, and playback begins.
//
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-media-playback.html#initiate-directive
type Initiate struct {
	RequestContext       RequestContext `json:"requestContext"`
	Filters              Filter         `json:"filters"`
	ContentId            string         `json:"contentId"`
	CurrentItemReference MediaReference `json:"currentItemReference"`
	PlaybackModes        PlaybackModes  `json:"playbackModes"`
}

// MediaReference is a polymorphic used to identify a specific media item (Item) to be used in cases
// where a media item must be identified in a domain-agnostic manner.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#mediareference
type MediaReference struct {
	Namespace string        `json:"namespace"`
	Name      string        `json:"name"`
	Value     ItemReference `json:"value"`
}

// ItemReference identifies a specific Item.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#itemreference
type ItemReference struct {
	Id        string `json:"id"`
	QueueId   string `json:"queueId"`
	ContentId string `json:"contentId"`
}

// PlaybackModes represent the playback modes requested by the user.
// If the user doesn't request a looped or shuffled queue, this attribute defaults to false
// for all supported playback modes.
type PlaybackModes struct {
	Shuffle bool `json:"shuffle"`
	Loop    bool `json:"loop"`
}

// InitiateResponse is a reply to an Initiate directive.
//
// The skill's Initiate response includes a play queue based on the requested ContentId, the first playable track,
// and enough information for Alexa to manage the queue.
//
// To get the second track, Alexa calls GetNextItem after the first track begins.
//
// Subsequent tracks are also retrieved with GetNextItem. For details, see Alexa.Audio.PlayQueue Interface.
//
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-media-playback.html
type InitiateResponse struct {
	PlaybackMethod PlaybackMethod `json:"playbackMethod"`
}

// PlaybackMethod is a realization of Content.
//
// When Alexa has a reference to some content as a result of GetPlayableContent, and wants to start
// playback for the user, it invokes the skill with the content reference so the skill can realize the content into
// a PlaybackMethod.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#playbackmethod
type PlaybackMethod struct {
	// The type of the playback method. The only allowed value is ALEXA_AUDIO_PLAYER_QUEUE.
	Type string `json:"type"`
	// The identifier for this play queue that is opaque to Alexa. If the skill wants to enforce concurrency limits,
	// then the queue ID should be unique for each user. Note that all Music Skill API requests after a queue is
	// created will report this queue ID back to the skill in all requests.
	Id string `json:"id"`
	// List of QueueControl objects used by Alexa to determine which controls should be clickable in the Alexa app.
	// Note that some queue-level controls might be overridden at the item-level.
	// See the QueueControl object for more information.
	Controls []QueueControl `json:"controls,omitempty"`
	// Identifies rules that apply to all audio items in the play queue.
	// Note that some queue-level rules may be overridden at the item-level.
	Rules QueueRules `json:"rules,omitempty"`
	// The first item from this queue that Alexa should play for the user.
	// See the Item object for more information.
	FirstItem Item `json:"firstItem"`
}

// QueueControl describes a control that the user could apply to a queue or the items in the queue.
//
// Examples are shuffle and loop. Note that some queue controls can be overridden at the item level.
// Note that if the enablement status of a control isn't specified, Alexa assumes the control is disabled.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#queuecontrol
type QueueControl struct {
	// The type of the control. This can be one of the following:
	//  ADJUST - for controls such as seeking
	//  COMMAND - for controls such as skipping
	//
	// QueueControl supports the following types in addition to the base types:
	//  TOGGLE - for controls such as shuffle
	Type string `json:"type"`
	// The name of the control.
	//
	// For type ADJUST, the value for name can be one of the following:
	//  SEEK_POSITION - enables user to seek forward and back within a currently playing track
	//
	// For type COMMAND, the value for name can be one of the following:
	//  NEXT - for skipping to the next item
	//  PREVIOUS - for skipping to the previous item
	//
	// For QueueControl type TOGGLE, the value can be one of the following:
	//  SHUFFLE - for toggling shuffle mode for the queue
	//  LOOP - for toggling loop mode for the queue
	//
	// Note the following controls are overridable at the item level:
	//  SEEK_POSITION, NEXT, PREVIOUS
	Name string `json:"name"`
	// Informs Alexa whether the control is enabled.
	//
	// For some control types, this determines whether the button for the control should be
	// clickable: set the value to true when the control should be clickable by the user in the Alexa app.
	Enabled bool `json:"enabled"`
	// Indicates that a control should render in a selected state.
	// For example, if a user has turned on loop mode in the provider's app, when the queue is cast to an Alexa device,
	// the loop control will have a selected value of true. Defaults to false.
	Selected bool `json:"selected,omitempty"`
}

// QueueRules identifies rules that apply to all items in the play queue.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#queuerules
type QueueRules struct {
	Feedback QueueFeedbackRule `json:"feedback"`
}

// QueueFeedbackRule describes whether feedback of a specific type is allowed at the queue level.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#queuefeedbackrule
type QueueFeedbackRule struct {
	// The type of feedback.
	//  The only supported value is PREFERENCE.
	Type string `json:"type"`
	// Whether feedback should be enabled.
	//
	// If false, Alexa will render error prompts (VUI) or show error messages (GUI) when the user tries to provide
	// feedback. Note that this flag can be overridden on an item-by-item basis.
	Enabled bool `json:"enabled"`
}

// Item is an audio item that contains within it an identifier and metadata
// (for example, art URLs, author name, title, etc.), and a Stream containing the playback URL.
//
// Note that in some cases, the duration of an item needs to be displayed when there is no stream for that item.
// That is why the duration field is a peer to the stream field and not a member of it.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#item
type Item struct {
	// A identifier for the item that is opaque to Alexa.
	// Note that the item should be unique within the queue.
	Id string `json:"id"`
	// Describes which business rules Alexa should run when playing or displaying this item. See the PlaybackInfo object for details.
	PlaybackInfo PlaybackInfo `json:"playbackInfo"`

	// Metadata for the item.
	//
	// For playbackInfo types DEFAULT and SAMPLE, the value is an instance of MediaMetadata, where the metadata
	// type field can be TRACK (for music), STATION (for radio), or PROGRAM (for podcasts).
	// See the MediaMetadata object for details.
	//
	// For playbackInfo type AD, the value is an instance of AdMetadata. See the AdMetadata object for details.
	// (Note: Ads aren't supported for podcasts.)
	// TODO: support AdMetadata also with common interface or generics
	Metadata MediaMetadata `json:"metadata"`
	// Duration of the item in milliseconds. If the item is a live stream, do not return this field.
	DurationInMilliseconds int `json:"durationInMilliseconds,omitempty"`
	// List of ItemControl objects used by Alexa to determine which controls should be clickable in the Alexa app.
	Controls []ItemControl `json:"controls,omitempty"`
	// Object which contains rules for the item. See ItemRules for details.
	Rules ItemRules `json:"rules"`
	// Object that contains stream information for this item. See the Stream object for more information.
	Stream Stream `json:"stream"`

	Feedback Feedback `json:"feedback"`
}

// PlaybackInfo describes which business rules Alexa should run when playing an item.
//
// More specifically for advertisements, Alexa will follow these rules:
//
//	No skipping or seeking allowed
//	No mention of the advertisement in prompts (or display use cases) for what's playing next
//	Alexa may want to inform users (in the app's Now Playing screen) that an advertisement is currently playing
//	Some potential differences in metadata for an advertisement versus a default track
//
// For samples, Alexa will render different prompts (mentioning that an item is a sample and not the full track),
// and it may potentially augment the item's metadata.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#playbackinfo
type PlaybackInfo struct {
	// The type of playback info. The field can have one of the following values:
	//
	//  DEFAULT - describes an item that is neither a sample nor an ad
	//  SAMPLE - describes an item whose stream is a shorter version of a full track
	//  AD - describes an item whose stream is an advertisement
	//  (Note: Ads aren't supported for podcasts.)
	Type string `json:"type"`
}

// ItemControl describes a control that the user can take on an item.
//
// Examples are skip forward and skip backward buttons.
//
// Note that item controls will override any existing queue controls of the same type.
//
// The object has the same structure as BaseControl.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#itemcontrol
type ItemControl struct {

}

// ItemRules describes rules for what the user can do with an item.
//
// One example of a rule is whether the user can provide feedback about an audio item.
//
// Note that item level rules will override queue level rules.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#itemrules
type ItemRules struct {
	FeedbackEnabled bool `json:"feedbackEnabled"`
}

// Stream contains the stream URI that Alexa sends to devices as part of an AudioPlayer play directive, along
// with other stream-related information such as the playback offset in milliseconds,
// the expiration time of the URI, and an identifier for the stream.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#stream
type Stream struct {
	// A unique opaque (to Alexa) identifier for the stream.
	//
	// Note that the same stream (with the same identifier) can be re-used across User and Queue instances.
	Id string `json:"id"`
	// URI of the stream (file or audio-chunk playlist).
	//
	// This URI is sent to the Alexa device media player to play the content.
	//
	// URIs must be HTTPS.
	Uri string `json:"uri"`
	// Offset, in milliseconds, at which to start playback of the stream.
	//
	// Defaults to 0.
	OffsetInMilliseconds int `json:"offsetInMilliseconds,omitempty"`
	// Offset, in milliseconds, at which to stop playback.
	//
	// If endOffsetInMilliseconds is greater than the actual track length, it's ignored, and the
	// stream plays to the end. Must be greater than or equal to offsetInMilliseconds.
	//
	// If endOffsetInMilliseconds is less than offsetInMilliseconds, a validation exception is raised.
	EndOffsetInMilliseconds int `json:"endOffsetInMilliseconds,omitempty"`
	// https://en.wikipedia.org/wiki/ISO_8601 representation of when the stream URI expires.
	//
	// Use the ISO 8601 extended format with UTC offset, for example 2019-01-29TT07:00:00.000Z.
	ValidUntil time.Time `json:"validUntil,omitempty"`
}

// Feedback describes the user's feedback or preference about an item.
//
// For example, if the user had previously indicated "thumbs up" for the item, the skill should set the
// feedback type to PREFERENCE and the feedback value to POSITIVE.
//
// https://developer.amazon.com/en-US/docs/alexa/music-skills/api-components-reference.html#feedback
type Feedback struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
