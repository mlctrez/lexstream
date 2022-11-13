package auditlogs

/*
Requester The user that performed the operation.
*/
type Requester struct {
	// LWA User ID. https://developer.amazon.com/docs/login-with-amazon/obtain-customer-profile.html
	UserId string `json:"userId,omitempty"`
}

/*
{
 "description": "The user that performed the operation.",
 "properties": {
  "userId": {
   "description": "LWA User ID. https://developer.amazon.com/docs/login-with-amazon/obtain-customer-profile.html",
   "type": "string"
  }
 },
 "type": "object"
}
*/
