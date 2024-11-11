package models

type PduSessionContext struct {
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	SmfBinding                 string             `json:"smfBinding,omitempty"`
	VsmfBinding                string             `json:"vsmfBinding,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	Dnn                        string             `json:"dnn"`
	AdditionalAccessType       string             `json:"additionalAccessType,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	IsmfBinding                string             `json:"ismfBinding,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	AccessType                 string             `json:"accessType"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	SNssai                     Snssai             `json:"sNssai"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
}
