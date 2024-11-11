package models

type UEContextRelease struct {
	Supi                string    `json:"supi,omitempty"`
	UnauthenticatedSupi *bool     `json:"unauthenticatedSupi,omitempty"`
	NgapCause           NgApCause `json:"ngapCause"`
}
