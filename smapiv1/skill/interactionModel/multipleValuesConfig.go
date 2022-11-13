package interactionmodel

/*
MultipleValuesConfig Configuration object for multiple values capturing behavior for this slot.
*/
type MultipleValuesConfig struct {
	// Boolean that indicates whether this slot can capture multiple values.
	Enabled bool `json:"enabled"`
}

/*
{
 "description": "Configuration object for multiple values capturing behavior for this slot.",
 "properties": {
  "enabled": {
   "description": "Boolean that indicates whether this slot can capture multiple values.",
   "type": "boolean"
  }
 },
 "type": "object"
}
*/
