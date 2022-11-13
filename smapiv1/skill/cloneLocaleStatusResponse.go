package skill

/*
CloneLocaleStatusResponse A mapping of statuses per locale detailing progress of resource or error if encountered.
*/
type CloneLocaleStatusResponse struct {
	Status *CloneLocaleRequestStatus `json:"status,omitempty"`
	Errors []*StandardizedError      `json:"errors,omitempty"`
	// Source locale which is cloned to target locales.
	SourceLocale string `json:"sourceLocale,omitempty"`
	// Mapping of statuses per locale.
	TargetLocales map[string]CloneLocaleResourceStatus `json:"targetLocales,omitempty"`
}

/*
{
 "description": "A mapping of statuses per locale detailing progress of resource or error if encountered.",
 "properties": {
  "errors": {
   "items": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   },
   "type": "array"
  },
  "sourceLocale": {
   "description": "Source locale which is cloned to target locales.",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.CloneLocaleRequestStatus",
   "x-isEnum": true
  },
  "targetLocales": {
   "additionalProperties": {
    "$ref": "#/definitions/v1.skill.CloneLocaleResourceStatus"
   },
   "description": "Mapping of statuses per locale.",
   "type": "object"
  }
 },
 "type": "object"
}
*/
