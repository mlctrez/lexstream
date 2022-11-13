package client

import (
	certification_ "github.com/mlctrez/lexstream/smapiv1/skill/certification"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetCertificationReviewV1 Gets a specific certification resource. The response contains the review tracking information for a skill to show how much time the skill is expected to remain under review by Amazon. Once the review is complete, the response also contains the outcome of the review. Old certifications may not be available, however any ongoing certification would always give a response. If the certification is unavailable the result will return a 404 HTTP status code.

	accept_Language - User's locale/language in context.
	skillId - The skill ID.
	certificationId - Id of the certification. Reserved word identifier of mostRecent can be used to get the most recent certification for the skill. Note that the behavior of the API in this case would be the same as when the actual certification id of the most recent certification is used in the request.
*/
func (s *Client) GetCertificationReviewV1(accept_Language string, skillId string, certificationId string) (response *certification_.CertificationResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/certifications/{certificationId}")
	h.Header("Accept-Language", accept_Language)
	h.Path("skillId", skillId)
	h.Path("certificationId", certificationId)
	response = &certification_.CertificationResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Gets a specific certification resource. The response contains the review tracking information for a skill to show how much time the skill is expected to remain under review by Amazon. Once the review is complete, the response also contains the outcome of the review. Old certifications may not be available, however any ongoing certification would always give a response. If the certification is unavailable the result will return a 404 HTTP status code.\n",
 "parameters": [
  {
   "description": "User's locale/language in context.",
   "in": "header",
   "name": "Accept-Language",
   "required": false,
   "type": "string"
  },
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Id of the certification. Reserved word identifier of mostRecent can be used to get the most recent certification for the skill. Note that the behavior of the API in this case would be the same as when the actual certification id of the most recent certification is used in the request.\n",
   "in": "path",
   "name": "certificationId",
   "required": true,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Successfully retrieved skill certification information.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.certification.CertificationResponse"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Exceeded the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.\n",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. only application/json supported",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "headers": {
    "Content-Language": {
     "description": "Standard HTTP header for language for which the content of the response is intended. Only en-US, ja-JP supported for this API currently.\n",
     "type": "string"
    },
    "Content-Type": {
     "description": "The content type of the response body. Only application/json supported.",
     "enum": [
      "application/json"
     ],
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getCertificationReviewV1"
}
*/
