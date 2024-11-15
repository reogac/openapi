package NpcfUEPolicyControl

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
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
	{
		Label:   "ReadIndividualUEPolicyAssociation",
		Method:  http.MethodGet,
		Path:    "/policies/:polAssoId",
		Handler: "OnReadIndividualUEPolicyAssociation",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   "NpcfUEPolicyControl",
		Routes:  _routes,
		Handler: p,
	}
}
