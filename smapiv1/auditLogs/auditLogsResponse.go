package auditlogs

/*
AuditLogsResponse Response to the Query Audit Logs API. It contains the collection of audit logs for the vendor, nextToken and other metadata related to the search query.
*/
type AuditLogsResponse struct {
	PaginationContext *ResponsePaginationContext `json:"paginationContext"`
	// List of audit logs for the vendor.
	AuditLogs []*AuditLog `json:"auditLogs"`
}

/*
{
 "description": "Response to the Query Audit Logs API. It contains the collection of audit logs for the vendor, nextToken and other metadata related to the search query.",
 "properties": {
  "auditLogs": {
   "description": "List of audit logs for the vendor.",
   "items": {
    "$ref": "#/definitions/v1.auditLogs.AuditLog"
   },
   "type": "array"
  },
  "paginationContext": {
   "$ref": "#/definitions/v1.auditLogs.ResponsePaginationContext"
  }
 },
 "type": "object"
}
*/
