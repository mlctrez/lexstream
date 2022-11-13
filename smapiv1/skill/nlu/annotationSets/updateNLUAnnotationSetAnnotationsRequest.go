package annotationsets

type UpdateNLUAnnotationSetAnnotationsRequest struct {
	// TODO: support serialization of examples
}

/*
{
 "example": {
  "application/json": {
   "data": [
    {
     "expected": [
      {
       "intent": {
        "name": "TravelIntent",
        "slots": {
         "fromCity": {
          "value": "seattle"
         }
        }
       }
      }
     ],
     "inputs": {
      "referenceTimestamp": "2018-10-25T23:50:02.135Z",
      "utterance": "i want to travel from seattle"
     }
    }
   ]
  },
  "text/csv": "utterance,referenceTimestamp,intent,slot[fromCity] i want to travel from seattle,2018-10-25T23:50:02.135Z,TravelIntent,seattle\n"
 },
 "type": "object"
}
*/
