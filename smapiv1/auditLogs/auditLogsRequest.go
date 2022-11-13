package auditlogs

type AuditLogsRequest struct {
	// Vendor Id. See developer.amazon.com/mycid.html.
	VendorId          string                    `json:"vendorId"`
	RequestFilters    *RequestFilters           `json:"requestFilters"`
	SortDirection     *SortDirection            `json:"sortDirection"`
	SortField         *SortField                `json:"sortField"`
	PaginationContext *RequestPaginationContext `json:"paginationContext"`
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
