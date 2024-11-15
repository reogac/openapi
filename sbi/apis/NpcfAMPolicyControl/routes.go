package NpcfAMPolicyControl

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "ReportObservedEventTriggersForIndividualAMPolicyAssociation",
		Method:  http.MethodPost,
		Path:    "/policies/:polAssoId/update",
		Handler: "OnReportObservedEventTriggersForIndividualAMPolicyAssociation",
	},
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
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "NpcfAMPolicyControl",
		Routes:  _routes,
		Handler: p,
	}
}
