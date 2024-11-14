package NpcfAMPolicyControl

const (
	PATH_ROOT string = "npcf-am-policy-control/v1"
)

// Summary: Create individual AM policy association.
// Description:
// Path: /policies
// Path Template: /policies
// Path Params:
func CreateIndividualAMPolicyAssociation(cli sbi.ConsumerClient, body *models.PolicyAssociationRequest) (rsp *models.PolicyAssociation, err error) {
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
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
// Path Template: /policies/%s
// Path Params: polAssoId
func ReadIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string) (rsp *models.PolicyAssociation, err error) {
	request := sbi.DefaultRequest()
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

//Summary: Report observed event triggers and obtain updated policies for an individual AM policy association.

// Description:
// Path: /policies/:polAssoId/update
// Path Template: /policies/%s/update
// Path Params: polAssoId
func ReportObservedEventTriggersForIndividualAMPolicyAssociation(cli sbi.ConsumerClient, polAssoId string, body *models.PolicyAssociationUpdateRequest) (rsp *models.PolicyUpdate, err error) {
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
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
