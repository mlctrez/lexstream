package amsapi

// GetNextItem is sent to the skill when a content queue exists and playback has started on the Alexa device.
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-audio-playqueue.html#getnextitem-directive
type GetNextItem struct {
	// Context information about the request.
	RequestContext *RequestContext `json:"requestContext"`
	// The currently playing item.
	CurrentItemReference *ItemReference `json:"currentItemReference"`
	// true if the user explicitly asked to skip to the next track (for music) or program (for podcasts),
	// false if the current track or program will soon end and the next item is needed.
	IsUserInitiated bool `json:"isUserInitiated,omitempty"`
}

// GetNextItemResponse is the reply to a GetNextItem request.
// If your skill has a next item to return, return the next item. If the currently playing item is the last item
// in the queue, your skill should send "isQueueFinished": true to indicate that there are no more items.
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-audio-playqueue.html#getnextitem-response-event
type GetNextItemResponse struct {
	// A flag that indicates whether the currently playing item is the last item in the queue.
	// When this value is false, the item field must be present and contain a non-null value.
	// When this value is true, the item field must be absent or set to null.
	IsQueueFinished bool `json:"isQueueFinished"`
	// The requested item in the play queue.
	Item *Item `json:"item"`
}

// GetPreviousItem request is sent to the skill when a content queue exists and playback has started on the Alexa device.
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-audio-playqueue.html#getpreviousitem-directive
type GetPreviousItem struct {
	// Context information about the request.
	RequestContext *RequestContext `json:"requestContext"`
	// The currently playing item.
	CurrentItemReference *ItemReference `json:"currentItemReference"`
}

// GetPreviousItemResponse should be returned when the skill has a previous item to return.
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-audio-playqueue.html#getpreviousitem-response-event
type GetPreviousItemResponse struct {
	// The previous item in the play queue.
	Item *Item `json:"item"`
}

// JumpToItem request is sent to the skill when audio content is playing from a music skill,
// and the user views the active queue of music (resulting from a GetView response) in the Alexa app
// then clicks on a track in the list to play.
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-audio-playqueue.html#jumptoitem-directive
type JumpToItem struct {
	// Context information about the request.
	RequestContext *RequestContext `json:"requestContext"`
	// The currently playing item.
	CurrentItemReference *ItemReference `json:"currentItemReference"`
	// The identifier of the item to jump to.
	TargetItemId string `json:"targetItemId"`
}

// JumpToItemResponse is returned when the requested item is valid and the jump is allowed.
// https://developer.amazon.com/en-US/docs/alexa/device-apis/alexa-audio-playqueue.html#jumptoitem-response-event
type JumpToItemResponse struct {
	// The requested item in the play queue.
	Item *Item `json:"item"`
}
