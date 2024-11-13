package models

type PduSessionContext struct {
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	Dnn                        string             `json:"dnn"`
	AccessType                 AccessType         `json:"accessType"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SmContextRef               string             `json:"smContextRef"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
}
