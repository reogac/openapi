package models

type PduSessionContext struct {
	SNssai                     Snssai             `json:"sNssai"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	Dnn                        string             `json:"dnn"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	AccessType                 AccessType         `json:"accessType"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
}
