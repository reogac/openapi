package models

type PduSessionContext struct {
	IsmfId                     string             `json:"ismfId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	Dnn                        string             `json:"dnn"`
	AccessType                 AccessType         `json:"accessType"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	SNssai                     Snssai             `json:"sNssai"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
}
