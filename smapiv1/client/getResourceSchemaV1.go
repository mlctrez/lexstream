package client

import (
	resourceSchema_ "github.com/mlctrez/lexstream/smapiv1/skill/resourceSchema"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetResourceSchemaV1 GetResourceSchema API provides schema for skill related resources. The schema returned by this API will be specific to vendor because it considers public beta features allowed for the vendor.

	resource - Name of the ASK resource for which schema is requested.
	vendorId - The vendor ID.
	operation - This parameter is required when resource is manifest because skill manifest schema differs based on operation. For example, submit for certification schema has more validations than create skill schema.
*/
func (s *Client) GetResourceSchemaV1(resource string, vendorId string, operation string) (response *resourceSchema_.GetResourceSchemaResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/resourceSchema/{resource}")
	h.Path("resource", resource)
	h.Param("vendorId", vendorId)
	h.Param("operation", operation)
	response = &resourceSchema_.GetResourceSchemaResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "GetResourceSchema API provides schema for skill related resources. The schema returned by this API will be specific to vendor because it considers public beta features allowed for the vendor.",
 "parameters": [
  {
   "description": "Name of the ASK resource for which schema is requested.",
   "enum": [
    "manifest",
    "skillPackageStructure"
   ],
   "in": "path",
   "name": "resource",
   "required": true,
   "type": "string"
  },
  {
   "description": "The vendor ID.",
   "in": "query",
   "name": "vendorId",
   "required": true,
   "type": "string"
  },
  {
   "description": "This parameter is required when resource is manifest because skill manifest schema differs based on operation. For example, submit for certification schema has more validations than create skill schema.",
   "enum": [
    "CREATE_SKILL",
    "UPDATE_SKILL",
    "SUBMIT_SKILL",
    "ENABLE_SKILL"
   ],
   "in": "query",
   "name": "operation",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Returns a S3 presigned URL to location of schema",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.resourceSchema.getResourceSchemaResponse"
   }
  },
  "400": {
   "description": "Invalid request",
   "headers": {
    "Content-Type": {
     "description": "returned content type; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "Unauthorized",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "403": {
   "description": "Forbidden",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "429": {
   "description": "Too Many Requests",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
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
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.Error"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
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
 "x-operation-name": "getResourceSchemaV1"
}
*/
