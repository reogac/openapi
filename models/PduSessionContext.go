package models

type PduSessionContext struct {
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	Dnn                        string             `json:"dnn"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
}
