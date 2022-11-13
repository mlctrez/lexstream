package evaluations

import smapiv1 "github.com/mlctrez/lexstream/smapiv1"

type Skill struct {
	Stage *smapiv1.StageType `json:"stage"`
	// skill locale in bcp 47 format
	Locale string `json:"locale"`
}

/*
{
 "properties": {
  "locale": {
   "description": "skill locale in bcp 47 format",
   "type": "string"
  },
  "stage": {
   "$ref": "#/definitions/v1.StageType",
   "x-isEnum": true
  }
 },
 "required": [
  "locale",
  "stage"
 ],
 "type": "object"
}
*/
