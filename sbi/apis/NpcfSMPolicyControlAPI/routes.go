/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:57 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfSMPolicyControlAPI

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "CreateSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies",
		Handler: "OnCreateSMPolicy",
	},
	{
		Label:   "GetSMPolicy",
		Method:  http.MethodGet,
		Path:    "/sm-policies/:smPolicyId",
		Handler: "OnGetSMPolicy",
	},
	{
		Label:   "UpdateSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies/:smPolicyId/update",
		Handler: "OnUpdateSMPolicy",
	},
	{
		Label:   "DeleteSMPolicy",
		Method:  http.MethodPost,
		Path:    "/sm-policies/:smPolicyId/delete",
		Handler: "OnDeleteSMPolicy",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
