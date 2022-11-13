package auditlogs

import "time"

type AuditLog struct {
	// UUID that identifies a specific request.
	XAmznRequestId string `json:"xAmznRequestId,omitempty"`
	// Date-Time when the request was made.
	Timestamp time.Time  `json:"timestamp,omitempty"`
	Client    *Client    `json:"client,omitempty"`
	Operation *Operation `json:"operation,omitempty"`
	// Contains information about the resources affected in this request.
	Resources []*Resource `json:"resources,omitempty"`
	Requester *Requester  `json:"requester,omitempty"`
	// HTTP Status Code returned by this request.
	HttpResponseCode int `json:"httpResponseCode,omitempty"`
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
