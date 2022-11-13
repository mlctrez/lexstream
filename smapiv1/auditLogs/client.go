package auditlogs

/*
Client Contains information about the Client that this request was performed by.
*/
type Client struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

/*
{
 "description": "Contains information about the Client that this request was performed by.",
 "properties": {
  "id": {
   "type": "string"
  },
  "name": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
