package models

type StatusInfo struct {
	ResourceStatus    string             `json:"resourceStatus"`
	Cause             string             `json:"cause,omitempty"`
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            string             `json:"anType,omitempty"`
}
