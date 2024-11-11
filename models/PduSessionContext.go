package models

type PduSessionContext struct {
	SNssai               Snssai             `json:"sNssai"`
	HsmfId               string             `json:"hsmfId,omitempty"`
	HsmfSetId            string             `json:"hsmfSetId,omitempty"`
	IsmfSetId            string             `json:"ismfSetId,omitempty"`
	SmfServiceInstanceId string             `json:"smfServiceInstanceId,omitempty"`
	PduSessionId         int                `json:"pduSessionId"`
	SelectedDnn          string             `json:"selectedDnn,omitempty"`
	AdditionalAccessType string             `json:"additionalAccessType,omitempty"`
	SmfBinding           string             `json:"smfBinding,omitempty"`
	VsmfServiceSetId     string             `json:"vsmfServiceSetId,omitempty"`
	IsmfId               string             `json:"ismfId,omitempty"`
	IsmfBinding          string             `json:"ismfBinding,omitempty"`
	AllocatedEbiList     []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfServiceSetId     string             `json:"hsmfServiceSetId,omitempty"`
	VsmfId               string             `json:"vsmfId,omitempty"`
	VsmfSetId            string             `json:"vsmfSetId,omitempty"`
	VsmfBinding          string             `json:"vsmfBinding,omitempty"`
	CnAssistedRanPara    *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	SmContextRef         string             `json:"smContextRef"`
	Dnn                  string             `json:"dnn"`
	AccessType           string             `json:"accessType"`
	IsmfServiceSetId     string             `json:"ismfServiceSetId,omitempty"`
	NsInstance           string             `json:"nsInstance,omitempty"`
	MaPduSession         *bool              `json:"maPduSession,omitempty"`
}
