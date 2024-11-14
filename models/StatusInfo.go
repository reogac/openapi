package models
type StatusInfo struct {
	 Cause	Cause	`json:"cause,omitempty"`
	 CnAssistedRanPara	*CnAssistedRanPara	`json:"cnAssistedRanPara,omitempty"`
	 AnType	AccessType	`json:"anType,omitempty"`
	 ResourceStatus	ResourceStatus	`json:"resourceStatus"`
}
