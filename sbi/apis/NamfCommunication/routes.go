package NamfCommunication

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "NonUeN2InfoSubscribe",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/subscriptions",
		Handler: "OnNonUeN2InfoSubscribe",
	},
	{
		Label:   "AMFStatusChangeSubscribe",
		Method:  http.MethodPost,
		Path:    "/subscriptions",
		Handler: "OnAMFStatusChangeSubscribe",
	},
	{
		Label:   "EBIAssignment",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/assign-ebi",
		Handler: "OnEBIAssignment",
	},
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
		Label:   "N1N2MessageUnSubscribe",
		Method:  http.MethodGet,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
		Handler: "OnN1N2MessageUnSubscribe",
	},
	{
		Label:   "NonUeN2InfoUnSubscribe",
		Method:  http.MethodGet,
		Path:    "/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId",
		Handler: "OnNonUeN2InfoUnSubscribe",
	},
	{
		Label:   "CancelRelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/cancel-relocate",
		Handler: "OnCancelRelocateUEContext",
	},
	{
		Label:   "NonUeN2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/transfer",
		Handler: "OnNonUeN2MessageTransfer",
	},
	{
		Label:   "CreateUEContext",
		Method:  http.MethodGet,
		Path:    "/ue-contexts/:ueContextId",
		Handler: "OnCreateUEContext",
	},
	{
		Label:   "ReleaseUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/release",
		Handler: "OnReleaseUEContext",
	},
	{
		Label:   "RelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/relocate",
		Handler: "OnRelocateUEContext",
	},
	{
		Label:   "N1N2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages",
		Handler: "OnN1N2MessageTransfer",
	},
	{
		Label:   "N1N2MessageSubscribe",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions",
		Handler: "OnN1N2MessageSubscribe",
	},
	{
		Label:   "AMFStatusChangeSubscribeModfy",
		Method:  http.MethodGet,
		Path:    "/subscriptions/:subscriptionId",
		Handler: "OnAMFStatusChangeSubscribeModfy",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "NamfCommunication",
		Routes:  _routes,
		Handler: p,
	}
}
