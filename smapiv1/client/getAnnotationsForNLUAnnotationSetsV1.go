package client

import swaggerlt "github.com/mlctrez/swaggerlt"

/*
GetAnnotationsForNLUAnnotationSetsV1

	skillId - The skill ID.
	annotationId - Identifier of the NLU annotation set.
	accept - Standard HTTP. Pass `application/json` or `test/csv` for GET calls.
*/
func (s *Client) GetAnnotationsForNLUAnnotationSetsV1(skillId string, annotationId string, accept string) (err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/nluAnnotationSets/{annotationId}/annotations")
	h.Path("skillId", skillId)
	h.Path("annotationId", annotationId)
	h.Header("Accept", accept)
	err = h.Execute(s.Client)
	return
}

/*
{
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Identifier of the NLU annotation set.",
   "in": "path",
   "name": "annotationId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Standard HTTP. Pass `application/json` or `test/csv` for GET calls.\n",
   "in": "header",
   "name": "Accept",
   "required": true,
   "type": "string"
  }
 ],
 "produces": [
  "application/json",
  "text/csv"
 ],
 "responses": {
  "200": {
   "description": "The specific version of a NLU annotation set has the content.",
   "examples": {
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
   "headers": {
    "Content-Type": {
     "description": "Returned content type, application/json supported",
     "type": "string"
    }
   }
  },
  "400": {
   "description": "Server cannot process the request due to a client error.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "Define the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "summary": "Get the annotations of an NLU annotation set",
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getAnnotationsForNLUAnnotationSetsV1"
}
*/
