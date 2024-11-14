package NpcfSMPolicyControlAPI

const (
	PATH_ROOT string = "npcf-smpolicycontrol/v1"
)

// Summary: Create a new Individual SM Policy
// Description:
// Path: /sm-policies
// Path Template: /sm-policies
// Path Params:
func CreateSMPolicy(cli sbi.ConsumerClient, body *models.SmPolicyContextData) (rsp *models.SmPolicyDecision, err error) {
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
		rsp = new(SmPolicyDecision)
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

// Summary: Read an Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId
// Path Template: /sm-policies/%s
// Path Params: smPolicyId
func GetSMPolicy(cli sbi.ConsumerClient, smPolicyId string) (rsp *models.SmPolicyControl, err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmPolicyControl)
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

// Summary: Update an existing Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId/update
// Path Template: /sm-policies/%s/update
// Path Params: smPolicyId
func UpdateSMPolicy(cli sbi.ConsumerClient, smPolicyId string, body *models.SmPolicyUpdateContextData) (rsp *models.SmPolicyDecision, err error) {
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
		rsp = new(SmPolicyDecision)
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

// Summary: Delete an existing Individual SM Policy
// Description:
// Path: /sm-policies/:smPolicyId/delete
// Path Template: /sm-policies/%s/delete
// Path Params: smPolicyId
func DeleteSMPolicy(cli sbi.ConsumerClient, smPolicyId string, body *models.SmPolicyDeleteData) (err error) {
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
	case 204:
		return
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
