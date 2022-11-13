package client

import (
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
SetSSLCertificatesV1 Updates the ssl certificates associated with this skill.

	skillId - The skill ID.
	sslCertificatePayload - Defines the input/output of the ssl certificates api for a skill.
*/
func (s *Client) SetSSLCertificatesV1(skillId string, sslCertificatePayload *skill_.SSLCertificatePayload) (err error) {
	h := swaggerlt.NewRequestHelper("put", s.Endpoint, "/v1/skills/{skillId}/sslCertificateSets/~latest")
	h.Path("skillId", skillId)
	h.Body = sslCertificatePayload
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Updates the ssl certificates associated with this skill.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "Defines the input/output of the ssl certificates api for a skill.",
   "in": "body",
   "name": "sslCertificatePayload",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.SSLCertificatePayload"
   }
  }
 ],
 "responses": {
  "204": {
   "description": "Accepted; Request was successful and get will now result in the new values.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
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
     "description": "Defines the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceeds the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "setSSLCertificatesV1"
}
*/
