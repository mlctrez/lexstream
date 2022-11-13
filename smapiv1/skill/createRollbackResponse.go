package skill

/*
CreateRollbackResponse Defines the response body when a rollback request is created
*/
type CreateRollbackResponse struct {
	// Defines the identifier for a rollback request.
	RollbackRequestId string `json:"rollbackRequestId,omitempty"`
}

/*
{
 "description": "Defines the response body when a rollback request is created",
 "properties": {
  "rollbackRequestId": {
   "description": "Defines the identifier for a rollback request.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
