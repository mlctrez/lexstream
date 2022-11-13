package metrics

/*
GetMetricDataResponse Response object for the API call which contains metrics data.
*/
type GetMetricDataResponse struct {
	// The name of metric which customer requested.
	Metric string `json,omitempty:"metric"`
	// The timestamps for the data points.
	Timestamps []string `json,omitempty:"timestamps"`
	// The data points for the metric corresponding to Timestamps.
	Values []int `json,omitempty:"values"`
	// A token that marks the next batch of returned results.
	NextToken string `json,omitempty:"nextToken"`
}

/*
{
 "description": "Response object for the API call which contains metrics data.",
 "properties": {
  "metric": {
   "description": "The name of metric which customer requested.",
   "type": "string"
  },
  "nextToken": {
   "description": "A token that marks the next batch of returned results.",
   "type": "string"
  },
  "timestamps": {
   "description": "The timestamps for the data points.",
   "items": {
    "format": "date-time",
    "type": "string"
   },
   "type": "array"
  },
  "values": {
   "description": "The data points for the metric corresponding to Timestamps.",
   "items": {
    "type": "number"
   },
   "type": "array"
  }
 },
 "required": [
  "metric",
  "timestamps",
  "values"
 ],
 "type": "object"
}
*/
