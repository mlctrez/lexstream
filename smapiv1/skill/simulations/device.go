package simulations

/*
Device Model of a virtual device used for simulation. This device object emulates attributes associated with a real Alexa enabled device.
*/
type Device struct {
	/*
	   A valid locale (e.g "en-US") for the virtual device used in simulation.
	*/
	Locale string `json:"locale,omitempty"`
}

/*
{
 "description": "Model of a virtual device used for simulation. This device object emulates attributes associated with a real Alexa enabled device.\n",
 "properties": {
  "locale": {
   "description": "A valid locale (e.g \"en-US\") for the virtual device used in simulation.\n",
   "type": "string"
  }
 },
 "required": [
  "locale"
 ],
 "type": "object"
}
*/
