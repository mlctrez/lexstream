package auditlogs

import "time"

/*
RequestFilters Request Filters for filtering audit logs.
*/
type RequestFilters struct {
	// List of Client IDs for filtering.
	Clients []*ClientFilter `json,omitempty:"clients"`
	// Filters for a list of operation names and versions.
	Operations []*OperationFilter `json,omitempty:"operations"`
	// Filters for a list of resources and/or their types. See documentation for allowed types.
	Resources  []*ResourceFilter  `json,omitempty:"resources"`
	Requesters []*RequesterFilter `json,omitempty:"requesters"`
	// Sets the start time for this search. Any audit logs with timestamps after this time (inclusive) will be included in the response.
	StartTime time.Time `json,omitempty:"startTime"`
	// Sets the end time for this search. Any audit logs with timestamps before this time (exclusive) will be included in the result.
	EndTime time.Time `json,omitempty:"endTime"`
	// Filters for HTTP response codes. For example, '200' or '503'
	HttpResponseCodes []string `json,omitempty:"httpResponseCodes"`
}

/*
{
 "description": "Request Filters for filtering audit logs.",
 "properties": {
  "clients": {
   "description": "List of Client IDs for filtering.",
   "items": {
    "$ref": "#/definitions/v1.auditLogs.ClientFilter"
   },
   "type": "array"
  },
  "endTime": {
   "description": "Sets the end time for this search. Any audit logs with timestamps before this time (exclusive) will be included in the result.",
   "format": "date-time",
   "type": "string"
  },
  "httpResponseCodes": {
   "description": "Filters for HTTP response codes. For example, '200' or '503'",
   "items": {
    "type": "string"
   },
   "type": "array"
  },
  "operations": {
   "description": "Filters for a list of operation names and versions.",
   "items": {
    "$ref": "#/definitions/v1.auditLogs.OperationFilter"
   },
   "type": "array"
  },
  "requesters": {
   "items": {
    "$ref": "#/definitions/v1.auditLogs.RequesterFilter"
   },
   "type": "array"
  },
  "resources": {
   "description": "Filters for a list of resources and/or their types. See documentation for allowed types.",
   "items": {
    "$ref": "#/definitions/v1.auditLogs.ResourceFilter"
   },
   "type": "array"
  },
  "startTime": {
   "description": "Sets the start time for this search. Any audit logs with timestamps after this time (inclusive) will be included in the response.",
   "format": "date-time",
   "type": "string"
  }
 },
 "type": "object"
}
*/
