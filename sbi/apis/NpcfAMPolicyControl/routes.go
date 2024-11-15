/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfAMPolicyControl

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "CreateIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies",
		Handler: "OnCreateIndividualAMPolicyAssociation",
	},
	{
		Label:   "ReadIndividualAMPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: "OnReadIndividualAMPolicyAssociation",
	},
	{
		Label:   "ReportObservedEventTriggersForIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies/:polAssoId/update",
		Handler: "OnReportObservedEventTriggersForIndividualAMPolicyAssociation",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
