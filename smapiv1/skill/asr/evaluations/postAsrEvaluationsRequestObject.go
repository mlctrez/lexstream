package evaluations

type PostAsrEvaluationsRequestObject struct {
	Skill *Skill `json:"skill,omitempty"`
	// ID for annotation set
	AnnotationSetId string `json:"annotationSetId,omitempty"`
}

/*
{
 "properties": {
  "annotationSetId": {
   "description": "ID for annotation set",
   "type": "string"
  },
  "skill": {
   "$ref": "#/definitions/v1.skill.asr.evaluations.Skill"
  }
 },
 "required": [
  "annotationSetId",
  "skill"
 ],
 "type": "object"
}
*/
