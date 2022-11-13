package annotationsets

import "time"

type AnnotationSetEntity struct {
	Locale string `json:"locale"`
	// Name of the NLU annotation set
	Name string `json:"name"`
	// Number of entries which represents number of utterances in each NLU annotation set content
	NumberOfEntries int `json:"numberOfEntries"`
	// The lastest updated timestamp for the NLU annotation set
	UpdatedTimestamp time.Time `json:"updatedTimestamp"`
}

/*
{
 "properties": {
  "locale": {
   "type": "string"
  },
  "name": {
   "description": "Name of the NLU annotation set",
   "type": "string"
  },
  "numberOfEntries": {
   "description": "Number of entries which represents number of utterances in each NLU annotation set content",
   "type": "integer"
  },
  "updatedTimestamp": {
   "description": "The lastest updated timestamp for the NLU annotation set",
   "example": "2018-10-25T23:50:02.135Z",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
