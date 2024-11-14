/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:58 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfUEPolicyControl

func OnReadIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleReadIndividualUEPolicyAssociation()
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.PolicyAssociationUpdateRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(body)
		if rsp != nil {
			response.SetBody(200, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnCreateIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.PolicyAssociationRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleCreateIndividualUEPolicyAssociation(body)
		if rsp != nil {
			response.SetBody(201, rsp)
			return
		}
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
