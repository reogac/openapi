package models
type LcsClientGroupExternal struct {
	 AllowedGeographicArea	[]GeographicArea	`json:"allowedGeographicArea,omitempty"`
	 PrivacyCheckRelatedAction	PrivacyCheckRelatedAction	`json:"privacyCheckRelatedAction,omitempty"`
	 ValidTimePeriod	*ValidTimePeriod	`json:"validTimePeriod,omitempty"`
	 LcsClientGroupId	string	`json:"lcsClientGroupId,omitempty"`
}
