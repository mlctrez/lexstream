package auditlogs

/*
RequesterFilter Filter for the requester of the operation.
*/
type RequesterFilter struct {
	// LoginWithAmazon User ID.
	UserId string `json:"userId"`
}

/*
{
 "description": "Filter for the requester of the operation.",
 "properties": {
  "userId": {
   "description": "LoginWithAmazon User ID.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
