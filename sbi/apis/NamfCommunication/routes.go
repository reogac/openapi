package NamfCommunication

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
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
		Label:   "N1N2MessageUnSubscribe",
		Method:  http.MethodGet,
		Path:    "/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId",
		Handler: "OnN1N2MessageUnSubscribe",
	},
	{
		Label:   "AMFStatusChangeSubscribeModfy",
		Method:  http.MethodGet,
		Path:    "/subscriptions/:subscriptionId",
		Handler: "OnAMFStatusChangeSubscribeModfy",
	},
	{
		Label:   "CreateUEContext",
		Method:  http.MethodGet,
		Path:    "/ue-contexts/:ueContextId",
		Handler: "OnCreateUEContext",
	},
	{
		Label:   "UEContextTransfer",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/transfer",
		Handler: "OnUEContextTransfer",
	},
	{
		Label:   "CancelRelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/cancel-relocate",
		Handler: "OnCancelRelocateUEContext",
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
		Label:   "RegistrationStatusUpdate",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/transfer-update",
		Handler: "OnRegistrationStatusUpdate",
	},
	{
		Label:   "NonUeN2InfoSubscribe",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/subscriptions",
		Handler: "OnNonUeN2InfoSubscribe",
	},
	{
		Label:   "ReleaseUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/release",
		Handler: "OnReleaseUEContext",
	},
	{
		Label:   "EBIAssignment",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/assign-ebi",
		Handler: "OnEBIAssignment",
	},
	{
		Label:   "RelocateUEContext",
		Method:  http.MethodPost,
		Path:    "/ue-contexts/:ueContextId/relocate",
		Handler: "OnRelocateUEContext",
	},
	{
		Label:   "NonUeN2MessageTransfer",
		Method:  http.MethodPost,
		Path:    "/non-ue-n2-messages/transfer",
		Handler: "OnNonUeN2MessageTransfer",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
