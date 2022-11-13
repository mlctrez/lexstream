package isp

/*
TaxInformationCategory Select tax category that best describes in-skill product. Choice will be validated during certification process.
*/
type TaxInformationCategory string

func TaxInformationCategory_SOFTWARE() TaxInformationCategory {
	return "SOFTWARE"
}
func TaxInformationCategory_STREAMING_AUDIO() TaxInformationCategory {
	return "STREAMING_AUDIO"
}
func TaxInformationCategory_STREAMING_RADIO() TaxInformationCategory {
	return "STREAMING_RADIO"
}
func TaxInformationCategory_INFORMATION_SERVICES() TaxInformationCategory {
	return "INFORMATION_SERVICES"
}
func TaxInformationCategory_VIDEO() TaxInformationCategory {
	return "VIDEO"
}
func TaxInformationCategory_PERIODICALS() TaxInformationCategory {
	return "PERIODICALS"
}
func TaxInformationCategory_NEWSPAPERS() TaxInformationCategory {
	return "NEWSPAPERS"
}
