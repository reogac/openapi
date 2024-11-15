/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:03:42 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfUEPolicyControl

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnReportObservedEventTriggersForIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// decode request body
	body := new(models.PolicyAssociationUpdateRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(polAssoId, body)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnCreateIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PolicyAssociationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualUEPolicyAssociation(body)

	// check for success response
	if rsp != nil {
		response.SetBody(201, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

func OnReadIndividualUEPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualUEPolicyAssociation(polAssoId)

	// check for success response
	if rsp != nil {
		response.SetBody(200, rsp)
		return
	}

	// check for problem
	if prob != nil {
		response.SetBody(models.ProblemDetailsCode(prob), prob)
		return
	}
	return
}

type Producer interface {
	HandleReportObservedEventTriggersForIndividualUEPolicyAssociation(string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)

	HandleCreateIndividualUEPolicyAssociation(*models.PolicyAssociationRequest) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualUEPolicyAssociation(string) (*models.PolicyAssociation, *models.ProblemDetails)
}
