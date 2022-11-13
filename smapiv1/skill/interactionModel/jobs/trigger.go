package jobs

/*
Trigger Condition when jobs will be executed.
*/
type Trigger struct {
	// Polymorphic type of the trigger
	Type_ string `json:"type,omitempty"`
}

/*
{
 "description": "Condition when jobs will be executed.",
 "discriminator": "type",
 "properties": {
  "type": {
   "description": "Polymorphic type of the trigger",
   "type": "string",
   "x-isDiscriminator": true
  }
 },
 "type": "object"
}
*/
