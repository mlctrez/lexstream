package evaluations

type ResolutionsPerAuthorityItems struct {
	/*
	   The name of the authority for the slot values. For custom slot types, this authority label incorporates your skill ID and the slot type name.
	*/
	Authority string                         `json:"authority"`
	Status    *ResolutionsPerAuthorityStatus `json:"status"`
	// An array of resolved values for the slot.
	Values []*ResolutionsPerAuthorityValueItems `json:"values"`
}

/*
{
 "properties": {
  "authority": {
   "description": "The name of the authority for the slot values. For custom slot types, this authority label incorporates your skill ID and the slot type name.\n",
   "type": "string"
  },
  "status": {
   "$ref": "#/definitions/v1.skill.evaluations.ResolutionsPerAuthorityStatus"
  },
  "values": {
   "description": "An array of resolved values for the slot.",
   "items": {
    "$ref": "#/definitions/v1.skill.evaluations.ResolutionsPerAuthorityValueItems"
   },
   "type": "array"
  }
 },
 "type": "object"
}
*/