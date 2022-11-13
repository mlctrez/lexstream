package skill

/*
ValidationDetails Standardized, machine readable structure that wraps all the information about a specific occurrence of an error of the type specified by the code.
*/
type ValidationDetails struct {
	// Set of properties of the image provided by the customer.
	ActualImageAttributes *ImageAttributes `json,omitempty:"actualImageAttributes"`
	// Number of items in an array provided by the customer.
	ActualNumberOfItems int `json,omitempty:"actualNumberOfItems"`
	// Number of characters in a string provided by the customer.
	ActualStringLength int `json,omitempty:"actualStringLength"`
	// List of allowed content types for a resource.
	AllowedContentTypes []string `json,omitempty:"allowedContentTypes"`
	// List of allowed data types for an instance.
	AllowedDataTypes []*ValidationDataTypes `json,omitempty:"allowedDataTypes"`
	// List of set of properties representing all possible allowed images.
	AllowedImageAttributes []*ImageAttributes `json,omitempty:"allowedImageAttributes"`
	// Instance conflicting with another instance.
	ConflictingInstance *Instance `json,omitempty:"conflictingInstance"`
	// Format in which instance value is expected in.
	ExpectedFormat *Format `json,omitempty:"expectedFormat"`
	// Instance that is expected by a related instance.
	ExpectedInstance *Instance `json,omitempty:"expectedInstance"`
	// Regular expression that a string instance is expected to match.
	ExpectedRegexPattern string `json,omitempty:"expectedRegexPattern"`
	// Type of the agreement that the customer must be compliant to.
	AgreementType *AgreementType `json,omitempty:"agreementType"`
	// Properties of a publicly known feature that has restricted access.
	Feature *ValidationFeature `json,omitempty:"feature"`
	// Endpoint which has a different value for property named type when compared to original endpoint.
	InconsistentEndpoint *ValidationEndpoint `json,omitempty:"inconsistentEndpoint"`
	// Minimum allowed value of an integer instance.
	MinimumIntegerValue int `json,omitempty:"minimumIntegerValue"`
	// Minimum allowed number of items in an array.
	MinimumNumberOfItems int `json,omitempty:"minimumNumberOfItems"`
	// Minimum allowed number of characters in a string.
	MinimumStringLength int `json,omitempty:"minimumStringLength"`
	// Maximum allowed value of an integer instance.
	MaximumIntegerValue int `json,omitempty:"maximumIntegerValue"`
	// Maximum allowed number of items in an array.
	MaximumNumberOfItems int `json,omitempty:"maximumNumberOfItems"`
	// Maximum allowed number of characters in a string.
	MaximumStringLength int `json,omitempty:"maximumStringLength"`
	// An Endpoint instance
	OriginalEndpoint *ValidationEndpoint `json,omitempty:"originalEndpoint"`
	// An Instance
	OriginalInstance *Instance `json,omitempty:"originalInstance"`
	// Represents what is wrong in the request.
	Reason *ValidationFailureReason `json,omitempty:"reason"`
	// Property required but missing in the object.
	RequiredProperty string `json,omitempty:"requiredProperty"`
	// Property not expected but present in the object.
	UnexpectedProperty string `json,omitempty:"unexpectedProperty"`
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
