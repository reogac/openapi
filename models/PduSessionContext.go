package models

type PduSessionContext struct {
	AllocatedEbiList           []EbiArpMapping    `json:"allocatedEbiList,omitempty"`
	NsInstance                 string             `json:"nsInstance,omitempty"`
	VsmfBinding                string             `json:"vsmfBinding,omitempty"`
	NrfAccessTokenUri          string             `json:"nrfAccessTokenUri,omitempty"`
	AnchorSmfSupportedFeatures string             `json:"anchorSmfSupportedFeatures,omitempty"`
	VsmfSetId                  string             `json:"vsmfSetId,omitempty"`
	SmfBindingInfo             string             `json:"smfBindingInfo,omitempty"`
	IsmfBindingInfo            string             `json:"ismfBindingInfo,omitempty"`
	AnchorSmfOauth2Required    *bool              `json:"anchorSmfOauth2Required,omitempty"`
	AdditionalAccessType       string             `json:"additionalAccessType,omitempty"`
	NrfDiscoveryUri            string             `json:"nrfDiscoveryUri,omitempty"`
	PgwIpAddr                  *IpAddress         `json:"pgwIpAddr,omitempty"`
	PlmnId                     *PlmnId            `json:"plmnId,omitempty"`
	HsmfId                     string             `json:"hsmfId,omitempty"`
	VsmfServiceSetId           string             `json:"vsmfServiceSetId,omitempty"`
	IsmfBinding                string             `json:"ismfBinding,omitempty"`
	SmfServiceInstanceId       string             `json:"smfServiceInstanceId,omitempty"`
	NrfManagementUri           string             `json:"nrfManagementUri,omitempty"`
	AdditionalSnssai           *Snssai            `json:"additionalSnssai,omitempty"`
	PgwFqdn                    string             `json:"pgwFqdn,omitempty"`
	VsmfBindingInfo            string             `json:"vsmfBindingInfo,omitempty"`
	SmContextRef               string             `json:"smContextRef"`
	AccessType                 string             `json:"accessType"`
	HsmfServiceSetId           string             `json:"hsmfServiceSetId,omitempty"`
	SmfBinding                 string             `json:"smfBinding,omitempty"`
	VsmfId                     string             `json:"vsmfId,omitempty"`
	IsmfSetId                  string             `json:"ismfSetId,omitempty"`
	CnAssistedRanPara          *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	MaPduSession               *bool              `json:"maPduSession,omitempty"`
	PduSessionId               int                `json:"pduSessionId"`
	SNssai                     Snssai             `json:"sNssai"`
	Dnn                        string             `json:"dnn"`
	SelectedDnn                string             `json:"selectedDnn,omitempty"`
	HsmfSetId                  string             `json:"hsmfSetId,omitempty"`
	IsmfId                     string             `json:"ismfId,omitempty"`
	IsmfServiceSetId           string             `json:"ismfServiceSetId,omitempty"`
	InterPlmnApiRoot           string             `json:"interPlmnApiRoot,omitempty"`
}
