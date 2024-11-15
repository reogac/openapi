/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:54 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmUEAU

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "ConfirmAuth",
		Method:  http.MethodPost,
		Path:    "/:supi/auth-events",
		Handler: "OnConfirmAuth",
	},
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
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
