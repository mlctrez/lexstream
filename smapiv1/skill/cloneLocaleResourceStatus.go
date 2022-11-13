package skill

/*
CloneLocaleResourceStatus an object detailing the status of a locale clone request and if applicable the errors occurred when saving/building resources during clone process.
*/
type CloneLocaleResourceStatus struct {
	Status *CloneLocaleStatus   `json:"status"`
	Errors []*StandardizedError `json:"errors"`
}

/*
{
 "description": "an object detailing the status of a locale clone request and if applicable the errors occurred when saving/building resources during clone process.",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.CloneLocaleStatus",
   "x-isEnum": true
  }
 },
 "type": "object"
}
*/
