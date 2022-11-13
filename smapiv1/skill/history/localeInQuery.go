package history

/*
LocaleInQuery A filter used to retrieve items where the locale is equal to the given value.
*/
type LocaleInQuery string

func LocaleInQuery_En_US() LocaleInQuery {
	return "en-US"
}
func LocaleInQuery_En_GB() LocaleInQuery {
	return "en-GB"
}
func LocaleInQuery_En_IN() LocaleInQuery {
	return "en-IN"
}
func LocaleInQuery_En_CA() LocaleInQuery {
	return "en-CA"
}
func LocaleInQuery_En_AU() LocaleInQuery {
	return "en-AU"
}
func LocaleInQuery_De_DE() LocaleInQuery {
	return "de-DE"
}
func LocaleInQuery_Ja_JP() LocaleInQuery {
	return "ja-JP"
}
