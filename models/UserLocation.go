package models

type UserLocation struct {
	GeraLocation  *GeraLocation  `json:"geraLocation,omitempty"`
	EutraLocation *EutraLocation `json:"eutraLocation,omitempty"`
	NrLocation    *NrLocation    `json:"nrLocation,omitempty"`
	N3gaLocation  *N3gaLocation  `json:"n3gaLocation,omitempty"`
	UtraLocation  *UtraLocation  `json:"utraLocation,omitempty"`
}
