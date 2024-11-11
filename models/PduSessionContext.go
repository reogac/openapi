package models

type PduSessionContext struct {
	IsmfServiceSetId     string             `json:"ismfServiceSetId,omitempty"`
	PduSessionId         int                `json:"pduSessionId"`
	AccessType           string             `json:"accessType"`
	HsmfSetId            string             `json:"hsmfSetId,omitempty"`
	VsmfBinding          string             `json:"vsmfBinding,omitempty"`
	IsmfId               string             `json:"ismfId,omitempty"`
	IsmfSetId            string             `json:"ismfSetId,omitempty"`
	CnAssistedRanPara    *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	SmContextRef         string             `json:"smContextRef"`
	HsmfId               string             `json:"hsmfId,omitempty"`
	VsmfSetId            string             `json:"vsmfSetId,omitempty"`
	VsmfServiceSetId     string             `json:"vsmfServiceSetId,omitempty"`
	SmfServiceInstanceId string             `json:"smfServiceInstanceId,omitempty"`
	MaPduSession         *bool              `json:"maPduSession,omitempty"`
	Dnn                  string             `json:"dnn"`
	AdditionalAccessType string             `json:"additionalAccessType,omitempty"`
	VsmfId               string             `json:"vsmfId,omitempty"`
	IsmfBinding          string             `json:"ismfBinding,omitempty"`
	SmfBinding           string             `json:"smfBinding,omitempty"`
	NsInstance           string             `json:"nsInstance,omitempty"`
	SNssai               Snssai             `json:"sNssai"`
	SelectedDnn          string             `json:"selectedDnn,omitempty"`
	AllocatedEbiList     []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfServiceSetId     string             `json:"hsmfServiceSetId,omitempty"`
}
