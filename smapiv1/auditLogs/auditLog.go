package auditlogs

import "time"

type AuditLog struct {
	// UUID that identifies a specific request.
	XAmznRequestId string `json,omitempty:"xAmznRequestId"`
	// Date-Time when the request was made.
	Timestamp time.Time  `json,omitempty:"timestamp"`
	Client    *Client    `json,omitempty:"client"`
	Operation *Operation `json,omitempty:"operation"`
	// Contains information about the resources affected in this request.
	Resources []*Resource `json,omitempty:"resources"`
	Requester *Requester  `json,omitempty:"requester"`
	// HTTP Status Code returned by this request.
	HttpResponseCode int `json,omitempty:"httpResponseCode"`
}

/*
{
 "properties": {
  "client": {
   "$ref": "#/definitions/v1.auditLogs.Client"
  },
  "httpResponseCode": {
   "description": "HTTP Status Code returned by this request.",
   "type": "integer"
  },
  "operation": {
   "$ref": "#/definitions/v1.auditLogs.Operation"
  },
  "requester": {
   "$ref": "#/definitions/v1.auditLogs.Requester"
  },
  "resources": {
   "description": "Contains information about the resources affected in this request.",
   "items": {
    "$ref": "#/definitions/v1.auditLogs.Resource"
   },
   "type": "array"
  },
  "timestamp": {
   "description": "Date-Time when the request was made.",
   "format": "date-time",
   "type": "string"
  },
  "xAmznRequestId": {
   "description": "UUID that identifies a specific request.",
   "format": "UUID",
   "type": "string"
  }
 },
 "type": "object"
}
*/
