package models

type PduSessionContext struct {
	SmContextRef               string             `json:"smContextRef"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	Dnn                        string             `json:"dnn"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
}
