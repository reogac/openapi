package models

type NsiInformation struct {
	NrfOauth2Required map[string]bool `json:"nrfOauth2Required,omitempty"`
	NrfId             string          `json:"nrfId"`
	NsiId             string          `json:"nsiId,omitempty"`
	NrfNfMgtUri       string          `json:"nrfNfMgtUri,omitempty"`
	NrfAccessTokenUri string          `json:"nrfAccessTokenUri,omitempty"`
}
