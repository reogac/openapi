package models

type PduSessionContext struct {
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	Dnn                        string             `json:"dnn"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	IsmfBinding                string             `json:"ismfBinding,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	SmfBinding                 string             `json:"smfBinding,omitempty"`
	VsmfBinding                string             `json:"vsmfBinding,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SmContextRef               string             `json:"smContextRef"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AccessType                 AccessType         `json:"accessType"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	AdditionalAccessType       AccessType         `json:"additionalAccessType,omitempty"`
}
