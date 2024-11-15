package NudmSDM

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "CAGAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/cag-ack",
		Handler: "OnCAGAck",
	},
	{
		Label:   "UpdateSORInfo",
		Method:  http.MethodPost,
		Path:    "/:supi/am-data/update-sor",
		Handler: "OnUpdateSORInfo",
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: "OnGetMultipleIdentifiers",
	},
	{
		Label:   "GetSmfSelData",
		Method:  http.MethodGet,
		Path:    "/:supi/smf-select-data",
		Handler: "OnGetSmfSelData",
	},
	{
		Label:   "GetLcsMoData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-mo-data",
		Handler: "OnGetLcsMoData",
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/:ueId/id-translation-result",
		Handler: "OnGetSupiOrGpsi",
	},
	{
		Label:   "GetTraceConfigData",
		Method:  http.MethodGet,
		Path:    "/:supi/trace-data",
		Handler: "OnGetTraceConfigData",
	},
	{
		Label:   "Unsubscribe",
		Method:  http.MethodGet,
		Path:    "/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: "OnUnsubscribe",
	},
	{
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data/:sharedDataId",
		Handler: "OnGetIndividualSharedData",
	},
	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data",
		Handler: "OnGetAmData",
	},
	{
		Label:   "GetSmsMngtData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-mng-data",
		Handler: "OnGetSmsMngtData",
	},
	{
		Label:   "GetLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/:ueId/lcs-privacy-data",
		Handler: "OnGetLcsPrivacyData",
	},
	{
		Label:   "GetV2xData",
		Method:  http.MethodGet,
		Path:    "/:supi/v2x-data",
		Handler: "OnGetV2xData",
	},
	{
		Label:   "GetMbsData",
		Method:  http.MethodGet,
		Path:    "/:supi/5mbs-data",
		Handler: "OnGetMbsData",
	},
	{
		Label:   "GetUcData",
		Method:  http.MethodGet,
		Path:    "/:supi/uc-data",
		Handler: "OnGetUcData",
	},
	{
		Label:   "SorAckInfo",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/sor-ack",
		Handler: "OnSorAckInfo",
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data",
		Handler: "OnGetSharedData",
	},
	{
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: "OnGetNSSAI",
	},
	{
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: "OnUnsubscribeForSharedData",
	},
	{
		Label:   "UpuAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/upu-ack",
		Handler: "OnUpuAck",
	},
	{
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-amf-data",
		Handler: "OnGetUeCtxInAmfData",
	},
	{
		Label:   "GetEcrData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/ecr-data",
		Handler: "OnGetEcrData",
	},
	{
		Label:   "GetUeCtxInSmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smf-data",
		Handler: "OnGetUeCtxInSmfData",
	},
	{
		Label:   "GetUeCtxInSmsfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smsf-data",
		Handler: "OnGetUeCtxInSmsfData",
	},
	{
		Label:   "GetSmData",
		Method:  http.MethodGet,
		Path:    "/:supi/sm-data",
		Handler: "OnGetSmData",
	},
	{
		Label:   "GetSmsData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-data",
		Handler: "OnGetSmsData",
	},
	{
		Label:   "GetProseData",
		Method:  http.MethodGet,
		Path:    "/:supi/prose-data",
		Handler: "OnGetProseData",
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: "OnGetDataSets",
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/group-data/group-identifiers",
		Handler: "OnGetGroupIdentifiers",
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: "OnSNSSAIsAck",
	},
	{
		Label:   "Subscribe",
		Method:  http.MethodPost,
		Path:    "/:ueId/sdm-subscriptions",
		Handler: "OnSubscribe",
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: "OnSubscribeToSharedData",
	},
	{
		Label:   "GetLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-bca-data",
		Handler: "OnGetLcsBcaData",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "NudmSDM",
		Routes:  _routes,
		Handler: p,
	}
}
