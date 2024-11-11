package models

type StatusInfo struct {
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType            string             `json:"anType,omitempty"`
	ResourceStatus    string             `json:"resourceStatus"`
	Cause             string             `json:"cause,omitempty"`
}
