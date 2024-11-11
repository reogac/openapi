package models

type NsiInformation struct {
	NrfAccessTokenUri string          `json:"nrfAccessTokenUri,omitempty"`
	NrfOauth2Required map[string]bool `json:"nrfOauth2Required,omitempty"`
	NrfId             string          `json:"nrfId"`
	NsiId             string          `json:"nsiId,omitempty"`
	NrfNfMgtUri       string          `json:"nrfNfMgtUri,omitempty"`
}
