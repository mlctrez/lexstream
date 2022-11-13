package private

/*
PrivateDistributionAccount Contains information of the private distribution account with given id.
*/
type PrivateDistributionAccount struct {
	// 12-digit numerical account ID for AWS account holders.
	Principal    string        `json,omitempty:"principal"`
	AcceptStatus *AcceptStatus `json,omitempty:"acceptStatus"`
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
