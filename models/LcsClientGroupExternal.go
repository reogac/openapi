package models
type LcsClientGroupExternal struct {
	 LcsClientGroupId	string	`json:"lcsClientGroupId,omitempty"`
	 AllowedGeographicArea	[]GeographicArea	`json:"allowedGeographicArea,omitempty"`
	 PrivacyCheckRelatedAction	PrivacyCheckRelatedAction	`json:"privacyCheckRelatedAction,omitempty"`
	 ValidTimePeriod	*ValidTimePeriod	`json:"validTimePeriod,omitempty"`
}
