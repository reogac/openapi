package models

type UEContextRelease struct {
	UnauthenticatedSupi *bool     `json:"unauthenticatedSupi,omitempty"`
	NgapCause           NgApCause `json:"ngapCause"`
	Supi                string    `json:"supi,omitempty"`
}
