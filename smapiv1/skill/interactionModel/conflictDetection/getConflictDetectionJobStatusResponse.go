package conflictdetection

type GetConflictDetectionJobStatusResponse struct {
	Status *ConflictDetectionJobStatus `json:"status"`
	// The total number of conflicts within skill model.
	TotalConflicts int `json:"totalConflicts"`
}

/*
{
 "properties": {
  "status": {
   "$ref": "#/definitions/v1.skill.interactionModel.conflictDetection.ConflictDetectionJobStatus",
   "x-isEnum": true
  },
  "totalConflicts": {
   "description": "The total number of conflicts within skill model.",
   "format": "long",
   "type": "number"
  }
 },
 "required": [
  "status"
 ],
 "type": "object"
}
*/
