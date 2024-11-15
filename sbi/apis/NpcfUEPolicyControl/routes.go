/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:59 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfUEPolicyControl

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "ReadIndividualUEPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: "OnReadIndividualUEPolicyAssociation",
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualUEPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies/:polAssoId/update",
		Handler: "OnReportObservedEventTriggersForIndividualUEPolicyAssociation",
	},
	{
		Label:   "CreateIndividualUEPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies",
		Handler: "OnCreateIndividualUEPolicyAssociation",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
