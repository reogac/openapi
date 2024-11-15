/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NamfCommunication

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "UEContextTransfer",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/transfer",
		Handler: "OnUEContextTransfer",
	},
	{
		Label:   "RegistrationStatusUpdate",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/transfer-update",
		Handler: "OnRegistrationStatusUpdate",
	},
	{
		Label:   "RelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/relocate",
		Handler: "OnRelocateUEContext",
	},
	{
		Label:   "N1N2MessageSubscribe",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions",
		Handler: "OnN1N2MessageSubscribe",
	},
	{
		Label:   "NonUeN2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/transfer",
		Handler: "OnNonUeN2MessageTransfer",
	},
	{
		Label:   "ReleaseUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/release",
		Handler: "OnReleaseUEContext",
	},
	{
		Label:   "NonUeN2InfoUnSubscribe",
		Method:  http.MethodGet,
		Path:    "/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId",
		Handler: "OnNonUeN2InfoUnSubscribe",
	},
	{
		Label:   "AMFStatusChangeSubscribe",
		Method:  http.MethodPost,
		Path:    "/subscriptions",
		Handler: "OnAMFStatusChangeSubscribe",
	},
	{
		Label:   "AMFStatusChangeSubscribeModfy",
		Method:  http.MethodGet,
		Path:    "/subscriptions/:subscriptionId",
		Handler: "OnAMFStatusChangeSubscribeModfy",
	},
	{
		Label:   "N1N2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages",
		Handler: "OnN1N2MessageTransfer",
	},
	{
		Label:   "CancelRelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/cancel-relocate",
		Handler: "OnCancelRelocateUEContext",
	},
	{
		Label:   "N1N2MessageUnSubscribe",
		Method:  http.MethodGet,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
		Handler: "OnN1N2MessageUnSubscribe",
	},
	{
		Label:   "EBIAssignment",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/assign-ebi",
		Handler: "OnEBIAssignment",
	},
	{
		Label:   "NonUeN2InfoSubscribe",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/subscriptions",
		Handler: "OnNonUeN2InfoSubscribe",
	},
	{
		Label:   "CreateUEContext",
		Method:  http.MethodGet,
		Path:    "/ue-contexts/:ueContextId",
		Handler: "OnCreateUEContext",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
