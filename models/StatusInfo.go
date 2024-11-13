package models

type StatusInfo struct {
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            AccessType         `json:"anType,omitempty"`
	ResourceStatus    ResourceStatus     `json:"resourceStatus"`
	Cause             Cause              `json:"cause,omitempty"`
}
