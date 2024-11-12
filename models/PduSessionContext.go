package models

type PduSessionContext struct {
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	Dnn                        string             `json:"dnn"`
	AccessType                 AccessType         `json:"accessType"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	VsmfBinding                SbiBindingLevel    `json:"vsmfBinding,omitempty"`
	IsmfBinding                SbiBindingLevel    `json:"ismfBinding,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	SNssai                     Snssai             `json:"sNssai"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	SmfBinding                 SbiBindingLevel    `json:"smfBinding,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
}
