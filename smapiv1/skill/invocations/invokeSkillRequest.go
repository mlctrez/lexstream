package invocations

type InvokeSkillRequest struct {
	EndpointRegion *EndPointRegions `json:"endpointRegion"`
	SkillRequest   *SkillRequest    `json:"skillRequest"`
}

/*
{
 "properties": {
  "endpointRegion": {
   "$ref": "#/definitions/v1.skill.invocations.EndPointRegions",
   "x-isEnum": true
  },
  "skillRequest": {
   "$ref": "#/definitions/v1.skill.invocations.SkillRequest"
  }
 },
 "required": [
  "endpointRegion",
  "skillRequest"
 ],
 "type": "object"
}
*/
