package models

type NsiInformation struct {
	NrfId             string             `json:"nrfId"`
	NsiId             string             `json:"nsiId,omitempty"`
	NrfNfMgtUri       string             `json:"nrfNfMgtUri,omitempty"`
	NrfAccessTokenUri string             `json:"nrfAccessTokenUri,omitempty"`
	NrfOauth2Required *nrfOauth2Required `json:"nrfOauth2Required,omitempty"`
}
