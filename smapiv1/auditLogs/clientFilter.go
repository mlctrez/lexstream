package auditlogs

/*
ClientFilter Identifier for the application the developer used to manage their skills and skill-related resources. For OAuth applications, this is the OAuth Client Id.
*/
type ClientFilter struct {
	Id string `json:"id"`
}

/*
{
 "description": "Identifier for the application the developer used to manage their skills and skill-related resources. For OAuth applications, this is the OAuth Client Id.",
 "properties": {
  "id": {
   "format": "UUID",
   "type": "string"
  }
 },
 "type": "object"
}
*/
