package NudmSDM

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "GetUeCtxInSmsfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-smsf-data",
		Handler: "OnGetUeCtxInSmsfData",
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
		Label:   "UpdateSORInfo",
		Method:  http.MethodPost,
		Path:    "/:supi/am-data/update-sor",
		Handler: "OnUpdateSORInfo",
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: "OnSubscribeToSharedData",
	},
	{
		Label:   "GetAmData",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data",
		Handler: "OnGetAmData",
	},
	{
		Label:   "GetSmData",
		Method:  http.MethodGet,
		Path:    "/:supi/sm-data",
		Handler: "OnGetSmData",
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
		Label:   "Subscribe",
		Method:  http.MethodPost,
		Path:    "/:ueId/sdm-subscriptions",
		Handler: "OnSubscribe",
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: "OnGetDataSets",
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: "OnGetMultipleIdentifiers",
	},
	{
		Label:   "Unsubscribe",
		Method:  http.MethodGet,
		Path:    "/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: "OnUnsubscribe",
	},
	{
		Label:   "UpuAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/upu-ack",
		Handler: "OnUpuAck",
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
		Label:   "GetSmsData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-data",
		Handler: "OnGetSmsData",
	},
	{
		Label:   "GetLcsMoData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-mo-data",
		Handler: "OnGetLcsMoData",
	},
	{
		Label:   "GetProseData",
		Method:  http.MethodGet,
		Path:    "/:supi/prose-data",
		Handler: "OnGetProseData",
	},
	{
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/:ueId/id-translation-result",
		Handler: "OnGetSupiOrGpsi",
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: "OnSNSSAIsAck",
	},
	{
		Label:   "GetIndividualSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data/:sharedDataId",
		Handler: "OnGetIndividualSharedData",
	},
	{
		Label:   "GetUeCtxInAmfData",
		Method:  http.MethodGet,
		Path:    "/:supi/ue-context-in-amf-data",
		Handler: "OnGetUeCtxInAmfData",
	},
	{
		Label:   "GetLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-bca-data",
		Handler: "OnGetLcsBcaData",
	},
	{
		Label:   "CAGAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/cag-ack",
		Handler: "OnCAGAck",
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/group-data/group-identifiers",
		Handler: "OnGetGroupIdentifiers",
	},
	{
		Label:   "GetSmfSelData",
		Method:  http.MethodGet,
		Path:    "/:supi/smf-select-data",
		Handler: "OnGetSmfSelData",
	},
	{
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: "OnGetNSSAI",
	},
	{
		Label:   "GetSmsMngtData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-mng-data",
		Handler: "OnGetSmsMngtData",
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
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: "OnUnsubscribeForSharedData",
	},
	{
		Label:   "GetTraceConfigData",
		Method:  http.MethodGet,
		Path:    "/:supi/trace-data",
		Handler: "OnGetTraceConfigData",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
