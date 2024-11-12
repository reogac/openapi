package models

type PduSessionContext struct {
	SNssai                     Snssai             `json:"sNssai"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SmContextRef               string             `json:"smContextRef"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	Dnn                        string             `json:"dnn"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
}
