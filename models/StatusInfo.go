package models

type StatusInfo struct {
	Cause             string             `json:"cause,omitempty"`
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            AccessType         `json:"anType,omitempty"`
	ResourceStatus    string             `json:"resourceStatus"`
}
