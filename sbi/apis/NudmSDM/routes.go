/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:53 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmSDM

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "UnsubscribeForSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data-subscriptions/:subscriptionId",
		Handler: "OnUnsubscribeForSharedData",
	},
	{
		Label:   "GetGroupIdentifiers",
		Method:  http.MethodGet,
		Path:    "/group-data/group-identifiers",
		Handler: "OnGetGroupIdentifiers",
	},
	{
		Label:   "GetNSSAI",
		Method:  http.MethodGet,
		Path:    "/:supi/nssai",
		Handler: "OnGetNSSAI",
	},
	{
		Label:   "GetLcsPrivacyData",
		Method:  http.MethodGet,
		Path:    "/:ueId/lcs-privacy-data",
		Handler: "OnGetLcsPrivacyData",
	},
	{
		Label:   "GetLcsBcaData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-bca-data",
		Handler: "OnGetLcsBcaData",
	},
	{
		Label:   "GetMbsData",
		Method:  http.MethodGet,
		Path:    "/:supi/5mbs-data",
		Handler: "OnGetMbsData",
	},
	{
		Label:   "UpuAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/upu-ack",
		Handler: "OnUpuAck",
	},
	{
		Label:   "SubscribeToSharedData",
		Method:  http.MethodPost,
		Path:    "/shared-data-subscriptions",
		Handler: "OnSubscribeToSharedData",
	},
	{
		Label:   "CAGAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/cag-ack",
		Handler: "OnCAGAck",
	},
	{
		Label:   "GetMultipleIdentifiers",
		Method:  http.MethodGet,
		Path:    "/multiple-identifiers",
		Handler: "OnGetMultipleIdentifiers",
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
		Label:   "GetSmfSelData",
		Method:  http.MethodGet,
		Path:    "/:supi/smf-select-data",
		Handler: "OnGetSmfSelData",
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
		Label:   "GetSupiOrGpsi",
		Method:  http.MethodGet,
		Path:    "/:ueId/id-translation-result",
		Handler: "OnGetSupiOrGpsi",
	},
	{
		Label:   "Unsubscribe",
		Method:  http.MethodGet,
		Path:    "/:ueId/sdm-subscriptions/:subscriptionId",
		Handler: "OnUnsubscribe",
	},
	{
		Label:   "SNSSAIsAck",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/subscribed-snssais-ack",
		Handler: "OnSNSSAIsAck",
	},
	{
		Label:   "UpdateSORInfo",
		Method:  http.MethodPost,
		Path:    "/:supi/am-data/update-sor",
		Handler: "OnUpdateSORInfo",
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
		Label:   "GetTraceConfigData",
		Method:  http.MethodGet,
		Path:    "/:supi/trace-data",
		Handler: "OnGetTraceConfigData",
	},
	{
		Label:   "GetSmsMngtData",
		Method:  http.MethodGet,
		Path:    "/:supi/sms-mng-data",
		Handler: "OnGetSmsMngtData",
	},
	{
		Label:   "GetLcsMoData",
		Method:  http.MethodGet,
		Path:    "/:supi/lcs-mo-data",
		Handler: "OnGetLcsMoData",
	},
	{
		Label:   "Subscribe",
		Method:  http.MethodPost,
		Path:    "/:ueId/sdm-subscriptions",
		Handler: "OnSubscribe",
	},
	{
		Label:   "SorAckInfo",
		Method:  http.MethodGet,
		Path:    "/:supi/am-data/sor-ack",
		Handler: "OnSorAckInfo",
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
		Label:   "GetV2xData",
		Method:  http.MethodGet,
		Path:    "/:supi/v2x-data",
		Handler: "OnGetV2xData",
	},
	{
		Label:   "GetProseData",
		Method:  http.MethodGet,
		Path:    "/:supi/prose-data",
		Handler: "OnGetProseData",
	},
	{
		Label:   "GetUcData",
		Method:  http.MethodGet,
		Path:    "/:supi/uc-data",
		Handler: "OnGetUcData",
	},
	{
		Label:   "GetSharedData",
		Method:  http.MethodGet,
		Path:    "/shared-data",
		Handler: "OnGetSharedData",
	},
	{
		Label:   "GetDataSets",
		Method:  http.MethodGet,
		Path:    "/:supi",
		Handler: "OnGetDataSets",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
