package models

type StatusInfo struct {
	ResourceStatus    ResourceStatus     `json:"resourceStatus"`
	Cause             Cause              `json:"cause,omitempty"`
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            AccessType         `json:"anType,omitempty"`
}
