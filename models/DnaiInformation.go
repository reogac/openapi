package models

type DnaiInformation struct {
	NoDnaiChangeInd     *bool  `json:"noDnaiChangeInd,omitempty"`
	NoLocalPsaChangeInd *bool  `json:"noLocalPsaChangeInd,omitempty"`
	Dnai                string `json:"dnai"`
}
