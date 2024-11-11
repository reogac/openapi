package models

type DnaiInformation struct {
	NoLocalPsaChangeInd *bool  `json:"noLocalPsaChangeInd,omitempty"`
	Dnai                string `json:"dnai"`
	NoDnaiChangeInd     *bool  `json:"noDnaiChangeInd,omitempty"`
}
