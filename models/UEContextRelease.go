package models

type UEContextRelease struct {
	NgapCause           NgApCause `json:"ngapCause"`
	Supi                string    `json:"supi,omitempty"`
	UnauthenticatedSupi *bool     `json:"unauthenticatedSupi,omitempty"`
}
