package skill

/*
ValidationDetails Standardized, machine readable structure that wraps all the information about a specific occurrence of an error of the type specified by the code.
*/
type ValidationDetails struct {
	// Set of properties of the image provided by the customer.
	ActualImageAttributes *ImageAttributes `json:"actualImageAttributes,omitempty"`
	// Number of items in an array provided by the customer.
	ActualNumberOfItems int `json:"actualNumberOfItems,omitempty"`
	// Number of characters in a string provided by the customer.
	ActualStringLength int `json:"actualStringLength,omitempty"`
	// List of allowed content types for a resource.
	AllowedContentTypes []string `json:"allowedContentTypes,omitempty"`
	// List of allowed data types for an instance.
	AllowedDataTypes []*ValidationDataTypes `json:"allowedDataTypes,omitempty"`
	// List of set of properties representing all possible allowed images.
	AllowedImageAttributes []*ImageAttributes `json:"allowedImageAttributes,omitempty"`
	// Instance conflicting with another instance.
	ConflictingInstance *Instance `json:"conflictingInstance,omitempty"`
	// Format in which instance value is expected in.
	ExpectedFormat *Format `json:"expectedFormat,omitempty"`
	// Instance that is expected by a related instance.
	ExpectedInstance *Instance `json:"expectedInstance,omitempty"`
	// Regular expression that a string instance is expected to match.
	ExpectedRegexPattern string `json:"expectedRegexPattern,omitempty"`
	// Type of the agreement that the customer must be compliant to.
	AgreementType *AgreementType `json:"agreementType,omitempty"`
	// Properties of a publicly known feature that has restricted access.
	Feature *ValidationFeature `json:"feature,omitempty"`
	// Endpoint which has a different value for property named type when compared to original endpoint.
	InconsistentEndpoint *ValidationEndpoint `json:"inconsistentEndpoint,omitempty"`
	// Minimum allowed value of an integer instance.
	MinimumIntegerValue int `json:"minimumIntegerValue,omitempty"`
	// Minimum allowed number of items in an array.
	MinimumNumberOfItems int `json:"minimumNumberOfItems,omitempty"`
	// Minimum allowed number of characters in a string.
	MinimumStringLength int `json:"minimumStringLength,omitempty"`
	// Maximum allowed value of an integer instance.
	MaximumIntegerValue int `json:"maximumIntegerValue,omitempty"`
	// Maximum allowed number of items in an array.
	MaximumNumberOfItems int `json:"maximumNumberOfItems,omitempty"`
	// Maximum allowed number of characters in a string.
	MaximumStringLength int `json:"maximumStringLength,omitempty"`
	// An Endpoint instance
	OriginalEndpoint *ValidationEndpoint `json:"originalEndpoint,omitempty"`
	// An Instance
	OriginalInstance *Instance `json:"originalInstance,omitempty"`
	// Represents what is wrong in the request.
	Reason *ValidationFailureReason `json:"reason,omitempty"`
	// Property required but missing in the object.
	RequiredProperty string `json:"requiredProperty,omitempty"`
	// Property not expected but present in the object.
	UnexpectedProperty string `json:"unexpectedProperty,omitempty"`
}

/*
{
 "description": "Standardized, machine readable structure that wraps all the information about a specific occurrence of an error of the type specified by the code.",
 "properties": {
  "actualImageAttributes": {
   "$ref": "#/definitions/v1.skill.imageAttributes",
   "description": "Set of properties of the image provided by the customer."
  },
  "actualNumberOfItems": {
   "description": "Number of items in an array provided by the customer.",
   "type": "integer"
  },
  "actualStringLength": {
   "description": "Number of characters in a string provided by the customer.",
   "type": "integer"
  },
  "agreementType": {
   "$ref": "#/definitions/v1.skill.AgreementType",
   "description": "Type of the agreement that the customer must be compliant to.",
   "x-isEnum": true
  },
  "allowedContentTypes": {
   "description": "List of allowed content types for a resource.",
   "items": {
    "description": "Allowed content type in IANA media type format.",
    "type": "string"
   },
   "type": "array"
  },
  "allowedDataTypes": {
   "description": "List of allowed data types for an instance.",
   "items": {
    "$ref": "#/definitions/v1.skill.ValidationDataTypes",
    "description": "Allowed data type."
   },
   "type": "array"
  },
  "allowedImageAttributes": {
   "description": "List of set of properties representing all possible allowed images.",
   "items": {
    "$ref": "#/definitions/v1.skill.imageAttributes",
    "description": "Set of properties of the image provided by the customer."
   },
   "type": "array"
  },
  "conflictingInstance": {
   "$ref": "#/definitions/v1.skill.Instance",
   "description": "Instance conflicting with another instance."
  },
  "expectedFormat": {
   "$ref": "#/definitions/v1.skill.Format",
   "description": "Format in which instance value is expected in.",
   "x-isEnum": true
  },
  "expectedInstance": {
   "$ref": "#/definitions/v1.skill.Instance",
   "description": "Instance that is expected by a related instance."
  },
  "expectedRegexPattern": {
   "description": "Regular expression that a string instance is expected to match.",
   "type": "string"
  },
  "feature": {
   "$ref": "#/definitions/v1.skill.ValidationFeature",
   "description": "Properties of a publicly known feature that has restricted access."
  },
  "inconsistentEndpoint": {
   "$ref": "#/definitions/v1.skill.ValidationEndpoint",
   "description": "Endpoint which has a different value for property named type when compared to original endpoint."
  },
  "maximumIntegerValue": {
   "description": "Maximum allowed value of an integer instance.",
   "type": "integer"
  },
  "maximumNumberOfItems": {
   "description": "Maximum allowed number of items in an array.",
   "type": "integer"
  },
  "maximumStringLength": {
   "description": "Maximum allowed number of characters in a string.",
   "type": "integer"
  },
  "minimumIntegerValue": {
   "description": "Minimum allowed value of an integer instance.",
   "type": "integer"
  },
  "minimumNumberOfItems": {
   "description": "Minimum allowed number of items in an array.",
   "type": "integer"
  },
  "minimumStringLength": {
   "description": "Minimum allowed number of characters in a string.",
   "type": "integer"
  },
  "originalEndpoint": {
   "$ref": "#/definitions/v1.skill.ValidationEndpoint",
   "description": "An Endpoint instance"
  },
  "originalInstance": {
   "$ref": "#/definitions/v1.skill.Instance",
   "description": "An Instance"
  },
  "reason": {
   "$ref": "#/definitions/v1.skill.ValidationFailureReason",
   "description": "Represents what is wrong in the request."
  },
  "requiredProperty": {
   "description": "Property required but missing in the object.",
   "type": "string"
  },
  "unexpectedProperty": {
   "description": "Property not expected but present in the object.",
   "type": "string"
  }
 },
 "type": "object"
}
*/
