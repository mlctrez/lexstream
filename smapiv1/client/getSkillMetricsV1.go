package client

import (
	metrics_ "github.com/mlctrez/lexstream/smapiv1/skill/metrics"
	swaggerlt "github.com/mlctrez/swaggerlt"
)

/*
GetSkillMetricsV1 Get analytic metrics report of skill usage.

	skillId - The skill ID.
	startTime - The start time of query.
	endTime - The end time of query (The maximum time duration is 1 week)
	period - The aggregation period to use when retrieving the metric, follows ISO_8601#Durations format.
	metric - A distinct set of logic which predictably returns a set of data.
	stage - The stage of the skill (live, development).
	skillType - The type of the skill (custom, smartHome and flashBriefing).
	intent - The intent of the skill.
	locale - The locale for the skill. e.g. en-GB, en-US, de-DE and etc.
	maxResults - Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.
	nextToken - When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.
*/
func (s *Client) GetSkillMetricsV1(skillId string, startTime string, endTime string, period string, metric string, stage string, skillType string, intent string, locale string, maxResults int, nextToken string) (response *metrics_.GetMetricDataResponse, err error) {
	h := swaggerlt.NewRequestHelper("get", s.Endpoint, "/v1/skills/{skillId}/metrics")
	h.Path("skillId", skillId)
	h.Param("startTime", startTime)
	h.Param("endTime", endTime)
	h.Param("period", period)
	h.Param("metric", metric)
	h.Param("stage", stage)
	h.Param("skillType", skillType)
	h.Param("intent", intent)
	h.Param("locale", locale)
	h.Param("maxResults", maxResults)
	h.Param("nextToken", nextToken)
	response = &metrics_.GetMetricDataResponse{}
	h.Response = response
	err = h.Execute(s.Client)
	return
}

/*
{
 "description": "Get analytic metrics report of skill usage.",
 "parameters": [
  {
   "description": "The skill ID.",
   "in": "path",
   "name": "skillId",
   "required": true,
   "type": "string"
  },
  {
   "description": "The start time of query.",
   "format": "date-time",
   "in": "query",
   "name": "startTime",
   "required": true,
   "type": "string"
  },
  {
   "description": "The end time of query (The maximum time duration is 1 week)",
   "format": "date-time",
   "in": "query",
   "name": "endTime",
   "required": true,
   "type": "string"
  },
  {
   "description": "The aggregation period to use when retrieving the metric, follows ISO_8601#Durations format.",
   "in": "query",
   "name": "period",
   "required": true,
   "type": "string"
  },
  {
   "description": "A distinct set of logic which predictably returns a set of data.",
   "in": "query",
   "name": "metric",
   "required": true,
   "type": "string"
  },
  {
   "description": "The stage of the skill (live, development).",
   "in": "query",
   "name": "stage",
   "required": true,
   "type": "string"
  },
  {
   "description": "The type of the skill (custom, smartHome and flashBriefing).",
   "in": "query",
   "name": "skillType",
   "required": true,
   "type": "string"
  },
  {
   "description": "The intent of the skill.",
   "in": "query",
   "name": "intent",
   "required": false,
   "type": "string"
  },
  {
   "description": "The locale for the skill. e.g. en-GB, en-US, de-DE and etc.",
   "in": "query",
   "name": "locale",
   "required": false,
   "type": "string"
  },
  {
   "description": "Sets the maximum number of results returned in the response body. If you want to retrieve fewer than upper limit of 50 results, you can add this parameter to your request. maxResults should not exceed the upper limit. The response might contain fewer results than maxResults, but it will never contain more. If there are additional results that satisfy the search criteria, but these results were not returned, the response contains isTruncated = true.",
   "in": "query",
   "maximum": 50,
   "minimum": 1,
   "multipleOf": 1,
   "name": "maxResults",
   "required": false,
   "type": "number"
  },
  {
   "description": "When response to this API call is truncated (that is, isTruncated response element value is true), the response also includes the nextToken element. The value of nextToken can be used in the next request as the continuation-token to list the next set of objects. The continuation token is an opaque value that Skill Management API understands. Token has expiry of 24 hours.",
   "in": "query",
   "name": "nextToken",
   "required": false,
   "type": "string"
  }
 ],
 "responses": {
  "200": {
   "description": "Get analytic metrics report successfully.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.metrics.GetMetricDataResponse"
   }
  },
  "400": {
   "description": "Bad request due to invalid or missing data.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.BadRequestError"
   }
  },
  "401": {
   "description": "The auth token is invalid/expired or doesn't have access to the resource.",
   "headers": {
    "WWW-Authenticate": {
     "description": "Defines the authentication method that should be used to gain access to a resource.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "403": {
   "description": "The operation being requested is not allowed.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "404": {
   "description": "The resource being requested is not found.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "429": {
   "description": "Exceed the permitted request limit. Throttling criteria includes total requests, per API, ClientId, and CustomerId.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "500": {
   "description": "Internal Server Error.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  },
  "503": {
   "description": "Service Unavailable.",
   "headers": {
    "Content-Type": {
     "description": "Returned content type, only application/json supported.",
     "type": "string"
    }
   },
   "schema": {
    "$ref": "#/definitions/v1.skill.StandardizedError"
   }
  }
 },
 "tags": [
  "skillManagement"
 ],
 "x-operation-name": "getSkillMetricsV1"
}
*/
