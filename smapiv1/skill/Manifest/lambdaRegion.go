package manifest

/*
LambdaRegion Defines the structure of a regional information.
*/
type LambdaRegion struct {
	Endpoint *LambdaEndpoint `json:"endpoint"`
}

/*
{
 "description": "Defines the structure of a regional information.",
 "properties": {
  "endpoint": {
   "$ref": "#/definitions/v1.skill.Manifest.LambdaEndpoint"
  }
 },
 "type": "object"
}
*/
