package auditlogs

/*
OperationFilter Name and version of the operation that the developer performed. For example, 'deleteSkill' and 'v1'. This is the same name used in the SMAPI SDK.
*/
type OperationFilter struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

/*
{
 "description": "Name and version of the operation that the developer performed. For example, 'deleteSkill' and 'v1'. This is the same name used in the SMAPI SDK.",
 "properties": {
  "name": {
   "type": "string"
  },
  "version": {
   "type": "string"
  }
 },
 "type": "object"
}
*/
