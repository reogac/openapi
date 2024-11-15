/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 17:41:11 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NpcfAMPolicyControl

import (
	"sbi"
	"sbi/models"
)

const (
	PATH_ROOT string = "npcf-am-policy-control/v1"
)

//Summary: Report observed event triggers and obtain updated policies for an individual AM policy association.

// Description:
// Path: /policies/:polAssoId/update
// Path Params: polAssoId
func ReportObservedEventTriggersForIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string, body *models.PolicyAssociationUpdateRequest) (rsp *models.PolicyUpdate, err error) {

	request := sbi.DefaultRequest()
	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/policies/%s/update", PATH_ROOT, polAssoId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(PolicyUpdate)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

// Summary: Create individual AM policy association.
// Description:
// Path: /policies
// Path Params:
func CreateIndividualAMPolicyAssociation(cli sbi.ConsumerClient, body *models.PolicyAssociationRequest) (rsp *models.PolicyAssociation, err error) {

	request := sbi.DefaultRequest()
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/policies", PATH_ROOT)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 201:
		rsp = new(PolicyAssociation)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 411, 413, 415, 429, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}

// Summary: Read individual AM policy association.
// Description:
// Path: /policies/:polAssoId
// Path Params: polAssoId
func ReadIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string) (rsp *models.PolicyAssociation, err error) {

	request := sbi.DefaultRequest()
	if len(polAssoId) == 0 {
		err = fmt.Errorf("polAssoId is required")
		return
	}

	request.Path = fmt.Sprintf("%s/policies/%s", PATH_ROOT, polAssoId)
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}

	switch response.StatusCode {
	case 200:
		rsp = new(PolicyAssociation)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 401, 403, 404, 429, 500, 503:
		prob := new(ProblemDetails)
		response.Body = prob
		if err = cli.DecodeResponse(response); err == nil {
			err = sbi.ErrorFromProblemDetails(prob)
		}
	default:
		err = fmt.Errorf("%d, %s", response.StatusCode, response.Status)
	}
	return
}
