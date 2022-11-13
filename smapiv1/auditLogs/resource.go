package auditlogs

/*
Resource Resource that the developer operated on. This includes both the type and ID of the resource.
*/
type Resource struct {
	Id    string `json:"id"`
	Type_ string `json:"type"`
}

/*
{
 "description": "Resource that the developer operated on. This includes both the type and ID of the resource.",
 "properties": {
  "id": {
   "format": "UUID",
   "type": "string"
  },
  "type": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
