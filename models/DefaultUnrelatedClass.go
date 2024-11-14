package models
type DefaultUnrelatedClass struct {
	 PrivacyCheckRelatedAction	PrivacyCheckRelatedAction	`json:"privacyCheckRelatedAction,omitempty"`
	 CodeWordInd	CodeWordInd	`json:"codeWordInd,omitempty"`
	 ValidTimePeriod	*ValidTimePeriod	`json:"validTimePeriod,omitempty"`
	 CodeWordList	[]string	`json:"codeWordList,omitempty"`
	 AllowedGeographicArea	[]GeographicArea	`json:"allowedGeographicArea,omitempty"`
}
