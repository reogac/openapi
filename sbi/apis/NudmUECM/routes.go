/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmUECM

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
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
		Label:   "GetSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations",
		Handler: "OnGetSmfRegistration",
	},
	{
		Label:   "Get3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-3gpp-access",
		Handler: "OnGet3GppSmsfRegistration",
	},
	{
		Label:   "GetNon3GppSmsfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smsf-non-3gpp-access",
		Handler: "OnGetNon3GppSmsfRegistration",
	},
	{
		Label:   "TriggerPCSCFRestoration",
		Method:  http.MethodPost,
		Path:    "/restore-pcscf",
		Handler: "OnTriggerPCSCFRestoration",
	},
	{
		Label:   "GetRegistrations",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations",
		Handler: "OnGetRegistrations",
	},
	{
		Label:   "SendRoutingInfoSm",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/send-routing-info-sm",
		Handler: "OnSendRoutingInfoSm",
	},
	{
		Label:   "Get3GppRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/amf-3gpp-access",
		Handler: "OnGet3GppRegistration",
	},
	{
		Label:   "DeregAMF",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/dereg-amf",
		Handler: "OnDeregAMF",
	},
	{
		Label:   "UpdateRoamingInformation",
		Method:  http.MethodPost,
		Path:    "/:ueId/registrations/amf-3gpp-access/roaming-info-update",
		Handler: "OnUpdateRoamingInformation",
	},
	{
		Label:   "RetrieveSmfRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/smf-registrations/:pduSessionId",
		Handler: "OnRetrieveSmfRegistration",
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
		Label:   "GetNwdafRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/nwdaf-registrations",
		Handler: "OnGetNwdafRegistration",
	},
	{
		Label:   "NwdafRegistration",
		Method:  http.MethodGet,
		Path:    "/:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId",
		Handler: "OnNwdafRegistration",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
