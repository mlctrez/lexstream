package skill

type ImportResponseSkill struct {
	SkillId   string                  `json:"skillId,omitempty"`
	ETag      string                  `json:"eTag,omitempty"`
	Resources []*ResourceImportStatus `json:"resources,omitempty"`
}

/*
{
 "properties": {
  "eTag": {
   "type": "string"
  },
  "resources": {
   "items": {
    "$ref": "#/definitions/v1.skill.ResourceImportStatus"
   },
   "type": "array"
  },
  "skillId": {
   "type": "string"
  }
 },
 "required": [
  "resources"
 ],
 "type": "object"
}
*/
