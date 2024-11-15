package NudmUECM

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "GetRegistrations",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations",
		Handler: "OnGetRegistrations",
	},
	{
		Label:   "Get3GppRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: "OnGet3GppRegistration",
	},
	{
		Label:   "PeiUpdate",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/pei-update",
		Handler: "OnPeiUpdate",
	},
	{
		Label:   "GetNon3GppRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/amf-non-3gpp-access",
		Handler: "OnGetNon3GppRegistration",
	},
	{
		Label:   "TriggerPCSCFRestoration",
		Method:  http.MethodPost,
		Path:    "/restore-pcscf",
		Handler: "OnTriggerPCSCFRestoration",
	},
	{
		Label:   "UpdateRoamingInformation",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/roaming-info-update",
		Handler: "OnUpdateRoamingInformation",
	},
	{
		Label:   "GetSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations",
		Handler: "OnGetSmfRegistration",
	},
	{
		Label:   "GetNon3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: "OnGetNon3GppSmsfRegistration",
	},
	{
		Label:   "NwdafRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: "OnNwdafRegistration",
	},
	{
		Label:   "SendRoutingInfoSm",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/send-routing-info-sm",
		Handler: "OnSendRoutingInfoSm",
	},
	{
		Label:   "DeregAMF",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/dereg-amf",
		Handler: "OnDeregAMF",
	},
	{
		Label:   "GetIpSmGwRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/ip-sm-gw",
		Handler: "OnGetIpSmGwRegistration",
	},
	{
		Label:   "GetLocationInfo",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/location",
		Handler: "OnGetLocationInfo",
	},
	{
		Label:   "RetrieveSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: "OnRetrieveSmfRegistration",
	},
	{
		Label:   "Get3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: "OnGet3GppSmsfRegistration",
	},
	{
		Label:   "GetNwdafRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/nwdaf-registrations",
		Handler: "OnGetNwdafRegistration",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "NudmUECM",
		Routes:  _routes,
		Handler: p,
	}
}
