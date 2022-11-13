package manifest

/*
SSLCertificateType The SSL certificate type of the skill's HTTPS endpoint. Only valid for HTTPS endpoint not for AWS Lambda ARN.
*/
type SSLCertificateType string

func SSLCertificateType_SelfSigned() SSLCertificateType {
	return "SelfSigned"
}
func SSLCertificateType_Wildcard() SSLCertificateType {
	return "Wildcard"
}
func SSLCertificateType_Trusted() SSLCertificateType {
	return "Trusted"
}
