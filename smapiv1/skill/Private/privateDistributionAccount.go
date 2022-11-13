package private

/*
PrivateDistributionAccount Contains information of the private distribution account with given id.
*/
type PrivateDistributionAccount struct {
	// 12-digit numerical account ID for AWS account holders.
	Principal    string        `json:"principal,omitempty"`
	AcceptStatus *AcceptStatus `json:"acceptStatus,omitempty"`
}

/*
{
 "description": "Contains information of the private distribution account with given id.",
 "properties": {
  "acceptStatus": {
   "$ref": "#/definitions/v1.skill.Private.AcceptStatus",
   "x-isEnum": true
  },
  "principal": {
   "description": "12-digit numerical account ID for AWS account holders.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
