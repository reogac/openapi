package models

type PlmnRestriction struct {
	ServiceAreaRestriction      *ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`
	CoreNetworkTypeRestrictions []string                `json:"coreNetworkTypeRestrictions,omitempty"`
	PrimaryRatRestrictions      []string                `json:"primaryRatRestrictions,omitempty"`
	SecondaryRatRestrictions    []string                `json:"secondaryRatRestrictions,omitempty"`
	RatRestrictions             []string                `json:"ratRestrictions,omitempty"`
	ForbiddenAreas              []Area                  `json:"forbiddenAreas,omitempty"`
}
