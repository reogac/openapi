/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:11:56 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfAMPolicyControl

import (
	"fmt"
	"sbi"
	"sbi/models"
)

func OnReportObservedEventTriggersForIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
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
	rsp, prob := prod.HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(polAssoId, body)

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

func OnCreateIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error

	// decode request body
	body := new(models.PolicyAssociationRequest)
	if err = ctx.DecodeRequest(body); err != nil {
		response.SetBody(400, models.NewSimpleProblem(400, fmt.Sprintf("Fail to decode request body: %+v", err)))
		return
	}

	// call application handler
	rsp, prob := prod.HandleCreateIndividualAMPolicyAssociation(body)

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

func OnReadIndividualAMPolicyAssociation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)

	// read 'polAssoId'
	var polAssoId string
	polAssoId = ctx.Param("polAssoId")
	if len(polAssoId) == 0 {
		response.SetBody(400, models.CreateProblemDetails(400, "polAssoId is required"))
		return
	}

	// call application handler
	rsp, prob := prod.HandleReadIndividualAMPolicyAssociation(polAssoId)

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
	HandleReportObservedEventTriggersForIndividualAMPolicyAssociation(string, *models.PolicyAssociationUpdateRequest) (*models.PolicyUpdate, *models.ProblemDetails)

	HandleCreateIndividualAMPolicyAssociation(*models.PolicyAssociationRequest) (*models.PolicyAssociation, *models.ProblemDetails)

	HandleReadIndividualAMPolicyAssociation(string) (*models.PolicyAssociation, *models.ProblemDetails)
}
