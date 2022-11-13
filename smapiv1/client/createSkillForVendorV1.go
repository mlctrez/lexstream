package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	skill_ "github.com/mlctrez/lexstream/smapiv1/skill"
	swaggerlt "github.com/mlctrez/swaggerlt"
	"io"
	"net/http"
	"strings"
)

/*
CreateSkillForVendorV1 Creates a new skill for given vendorId.

	createSkillRequest - Defines the request body for createSkill API.
*/
func (s *Client) CreateSkillForVendorV1(createSkillRequest *skill_.CreateSkillRequest) (response *skill_.CreateSkillResponse, err error) {
	u := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/skills")
	u.Body = createSkillRequest
	response = &skill_.CreateSkillResponse{}
	u.Response = response

	// calculate url from parameters
	uri := u.Uri
	for name, param := range u.PathParam {
		uri = strings.ReplaceAll(uri, fmt.Sprintf("{%s}", name), param)
	}
	if len(u.QueryValues) > 0 {
		uri += "?" + u.QueryValues.Encode()
	}

	var body io.Reader

	if u.Body != nil {
		marshal, err := json.Marshal(u.Body)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(marshal)
	} else {
		body = bytes.NewReader([]byte{})
	}

	request, err := http.NewRequest(u.Method, u.Endpoint+uri, body)
	if err != nil {
		return
	}

	if u.Body != nil {
		request.Header.Set("Content-Type", "application/json")
	}
	for header, value := range u.Headers {
		request.Header.Set(header, value)
	}

	res, err := s.Client.Do(request)
	if err != nil {
		return nil, err
	}

	if u.Response != nil {
		var all []byte
		all, err = io.ReadAll(res.Body)
		fmt.Println(string(all))

		err = json.NewDecoder(res.Body).Decode(u.Response)
		if err != nil {
			return
		}
	} else {
		_, _ = io.ReadAll(request.Body)
		_ = request.Body.Close()
	}

	return
}

/*
{
 "description": "Creates a new skill for given vendorId.",
 "parameters": [
  {
   "description": "Defines the request body for createSkill API.",
   "in": "body",
   "name": "createSkillRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.skill.createSkillRequest"
   }
  }
 ],
 "responses": {
  "202": {
   "description": "Accepted; Returns a URL to track the status in 'Location' header.",
   "headers": {
    "Content-Type": {
     "description": "Contains content type of the response; only application/json supported.",
     "type": "string"
    },
    "Location": {
     "description": "Contains relative URL to get skill status.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.CreateSkillResponse"
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
  "403": {
   "description": "The operation being requested is not allowed.",
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
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
 "x-operation-name": "createSkillForVendorV1"
}
*/
