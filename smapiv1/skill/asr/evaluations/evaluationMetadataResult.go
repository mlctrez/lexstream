package evaluations

/*
EvaluationMetadataResult indicate the result of the evaluation. This field would be present if the evaluation status is `COMPLETED`
*/
type EvaluationMetadataResult struct {
	Status  *EvaluationResultStatus `json,omitempty:"status"`
	Metrics *Metrics                `json,omitempty:"metrics"`
}

/*
{
 "description": "indicate the result of the evaluation. This field would be present if the evaluation status is `COMPLETED`",
 "properties": {
  "metrics": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.Metrics"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.EvaluationResultStatus",
   "x-isEnum": true
  }
 },
 "required": [
  "metrics",
  "status"
 ],
 "type": "object"
}
*/
