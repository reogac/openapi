package AUSFAPI

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
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
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "AUSFAPI",
		Routes:  _routes,
		Handler: p,
	}
}
