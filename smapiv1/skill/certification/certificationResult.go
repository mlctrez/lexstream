package certification

/*
CertificationResult Structure for the result for the outcomes of certification review for the skill. Currently provides the distribution information of a skill if the certification SUCCEEDED.
*/
type CertificationResult struct {
	DistributionInfo *DistributionInfo `json:"distributionInfo,omitempty"`
}

/*
{
 "description": "Structure for the result for the outcomes of certification review for the skill. Currently provides the distribution information of a skill if the certification SUCCEEDED.\n",
 "properties": {
  "distributionInfo": {
   "$ref": "#/definitions/v1.skill.certification.DistributionInfo"
  }
 },
 "type": "object"
}
*/
