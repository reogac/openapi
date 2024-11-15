/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:02 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package AUSFAPI

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "RgAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/rg-authentications",
		Handler: "OnRgAuthenticationsPost",
	},
	{
		Label:   "ProseAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/prose-authentications",
		Handler: "OnProseAuthenticationsPost",
	},
	{
		Label:   "ProseAuth",
		Method:  http.MethodPost,
		Path:    "/prose-authentications/:authCtxId/prose-auth",
		Handler: "OnProseAuth",
	},
	{
		Label:   "UeAuthenticationsPost",
		Method:  http.MethodPost,
		Path:    "/ue-authentications",
		Handler: "OnUeAuthenticationsPost",
	},
	{
		Label:   "UeAuthenticationsDeregisterPost",
		Method:  http.MethodPost,
		Path:    "/ue-authentications/deregister",
		Handler: "OnUeAuthenticationsDeregisterPost",
	},
	{
		Label:   "UeAuthentications5gAkaConfirmationPut",
		Method:  http.MethodGet,
		Path:    "/ue-authentications/:authCtxId/5g-aka-confirmation",
		Handler: "OnUeAuthentications5gAkaConfirmationPut",
	},
	{
		Label:   "EapAuthMethod",
		Method:  http.MethodPost,
		Path:    "/ue-authentications/:authCtxId/eap-session",
		Handler: "OnEapAuthMethod",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
