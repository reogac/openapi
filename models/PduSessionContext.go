package models

type PduSessionContext struct {
	SNssai                     Snssai             `json:"sNssai"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	Dnn                        string             `json:"dnn"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SmContextRef               string             `json:"smContextRef"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
}
