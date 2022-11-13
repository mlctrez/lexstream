package publication

import "time"

type PublishSkillRequest struct {
	// Used to determine when the skill Publishing should start. It takes the request timestamp as default value. The date range can be a maximum of upto 6 months from the current time stamp. The format should be the  RFC 3399 variant of ISO 8601. e.g 2019-04-12T23:20:50.52Z
	PublishesAtDate time.Time `json,omitempty:"publishesAtDate"`
}

/*
{
 "properties": {
  "publishesAtDate": {
   "description": "Used to determine when the skill Publishing should start. It takes the request timestamp as default value. The date range can be a maximum of upto 6 months from the current time stamp. The format should be the  RFC 3399 variant of ISO 8601. e.g 2019-04-12T23:20:50.52Z",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
