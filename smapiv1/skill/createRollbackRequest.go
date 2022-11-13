package skill

/*
CreateRollbackRequest Defines the request body to create a rollback request
*/
type CreateRollbackRequest struct {
	// The target skill version to rollback to.
	TargetVersion string `json:"targetVersion,omitempty"`
}

/*
{
 "description": "Defines the request body to create a rollback request",
 "properties": {
  "targetVersion": {
   "description": "The target skill version to rollback to.",
   "type": "string"
  }
 },
 "required": [
  "targetVersion"
 ],
 "type": "object"
}
*/
