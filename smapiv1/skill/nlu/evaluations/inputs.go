package evaluations

import "time"

type Inputs struct {
	Utterance string `json,omitempty:"utterance"`
	// Datetime to use to base date operations on.
	ReferenceTimestamp time.Time `json,omitempty:"referenceTimestamp"`
}

/*
{
 "properties": {
  "referenceTimestamp": {
   "description": "Datetime to use to base date operations on.",
   "example": "2018-10-25T23:50:02.135Z",
   "format": "date-time",
   "type": "string"
  },
  "utterance": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
