package auditlogs

type AuditLogsRequest struct {
	// Vendor Id. See developer.amazon.com/mycid.html.
	VendorId          string                    `json:"vendorId,omitempty"`
	RequestFilters    *RequestFilters           `json:"requestFilters,omitempty"`
	SortDirection     *SortDirection            `json:"sortDirection,omitempty"`
	SortField         *SortField                `json:"sortField,omitempty"`
	PaginationContext *RequestPaginationContext `json:"paginationContext,omitempty"`
}

/*
{
 "properties": {
  "paginationContext": {
   "$ref": "#/definitions/v1.auditLogs.RequestPaginationContext"
  },
  "requestFilters": {
   "$ref": "#/definitions/v1.auditLogs.RequestFilters"
  },
  "sortDirection": {
   "$ref": "#/definitions/v1.auditLogs.SortDirection",
   "x-isEnum": true
  },
  "sortField": {
   "$ref": "#/definitions/v1.auditLogs.SortField",
   "x-isEnum": true
  },
  "vendorId": {
   "description": "Vendor Id. See developer.amazon.com/mycid.html.",
   "type": "string"
  }
 },
 "required": [
  "vendorId"
 ],
 "type": "object"
}
*/
