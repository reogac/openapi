package models

type PduSessionContext struct {
	SmContextRef               string             `json:"smContextRef"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	Dnn                        string             `json:"dnn"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
}
