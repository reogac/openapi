package NudmUEAU

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "GenerateAv",
		Method:  http.MethodPost,
		Path:    "/:supi/hss-security-information/:hssAuthType/generate-av",
		Handler: "OnGenerateAv",
	},
	{
		Label:   "DeleteAuth",
		Method:  http.MethodGet,
		Path:    "/:supi/auth-events/:authEventId",
		Handler: "OnDeleteAuth",
	},
	{
		Label:   "GenerateGbaAv",
		Method:  http.MethodPost,
		Path:    "/:supi/gba-security-information/generate-av",
		Handler: "OnGenerateGbaAv",
	},
	{
		Label:   "GenerateProseAV",
		Method:  http.MethodPost,
		Path:    "/:supiOrSuci/prose-security-information/generate-av",
		Handler: "OnGenerateProseAV",
	},
	{
		Label:   "GenerateAuthData",
		Method:  http.MethodPost,
		Path:    "/:supiOrSuci/security-information/generate-auth-data",
		Handler: "OnGenerateAuthData",
	},
	{
		Label:   "GetRgAuthData",
		Method:  http.MethodGet,
		Path:    "/:supiOrSuci/security-information-rg",
		Handler: "OnGetRgAuthData",
	},
	{
		Label:   "ConfirmAuth",
		Method:  http.MethodPost,
		Path:    "/:supi/auth-events",
		Handler: "OnConfirmAuth",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
