package client

import (
	auditLogs_ "github.com/mlctrez/lexstream/smapiv1/auditLogs"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
QueryDevelopmentAuditLogsV1 The SMAPI Audit Logs API provides customers with an audit history of all SMAPI calls made by a developer or developers with permissions on that account.

	getAuditLogsRequest - Request object encompassing vendorId, optional request filters and optional pagination context.
*/
func (s *Client) QueryDevelopmentAuditLogsV1(getAuditLogsRequest *auditLogs_.AuditLogsRequest) (response *auditLogs_.AuditLogsResponse, err error) {
	h := swaggerlt.NewRequestHelper("post", s.Endpoint, "/v1/developmentAuditLogs/query")
	h.Body = getAuditLogsRequest
	response = &auditLogs_.AuditLogsResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "The SMAPI Audit Logs API provides customers with an audit history of all SMAPI calls made by a developer or developers with permissions on that account.",
 "parameters": [
  {
   "description": "Request object encompassing vendorId, optional request filters and optional pagination context.",
   "in": "body",
   "name": "getAuditLogsRequest",
   "required": true,
   "schema": {
    "$ref": "#/definitions/v1.auditLogs.AuditLogsRequest"
   }
  }
 ],
 "responses": {
  "200": {
   "description": "Returns a list of audit logs for the given vendor.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type; only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.auditLogs.AuditLogsResponse"
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
  "404": {
   "description": "Not Found",
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
 "x-operation-name": "queryDevelopmentAuditLogsV1"
}
*/
