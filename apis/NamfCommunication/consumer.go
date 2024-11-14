package NamfCommunication

const (
	PATH_ROOT string = "namf-comm/v1"
)

// Summary: Namf_Communication Non UE N2 Message Transfer service Operation
// Description:
// Path: /non-ue-n2-messages/transfer
// Path Template: /non-ue-n2-messages/transfer
// Path Params:
func NonUeN2MessageTransfer(cli sbi.ConsumerClient, body *models.NonUeN2MessageTransferRequest) (rsp *models.N2InformationTransferRspData, ersp *models.N2InformationTransferError, err error) {
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
		rsp = new(N2InformationTransferRspData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
		ersp = new(N2InformationTransferError)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411, 413, 415, 429:
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

// Summary: Namf_Communication Non UE N2 Info UnSubscribe service Operation
// Description:
// Path: /non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId
// Path Template: /non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId
// Path Params: n2NotifySubscriptionId
func NonUeN2InfoUnSubscribe(cli sbi.ConsumerClient, n2NotifySubscriptionId string) (err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 404, 429, 500, 503:
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

// Summary: Namf_Communication AMF Status Change Subscribe Modify service Operation
// Description:
// Path: /subscriptions/:subscriptionId
// Path Template: /subscriptions/%s
// Path Params: subscriptionId
func AMFStatusChangeSubscribeModfy(cli sbi.ConsumerClient, subscriptionId string, body *models.SubscriptionData) (rsp *models.SubscriptionData, err error) {
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
		rsp = new(SubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication RegistrationStatusUpdate service Operation
// Description:
// Path: /ue-contexts/:ueContextId/transfer-update
// Path Template: /ue-contexts/%s/transfer-update
// Path Params: ueContextId
func RegistrationStatusUpdate(cli sbi.ConsumerClient, ueContextId string, body *models.UeRegStatusUpdateReqData) (rsp *models.UeRegStatusUpdateRspData, err error) {
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
		rsp = new(UeRegStatusUpdateRspData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication CreateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId
// Path Template: /ue-contexts/%s
// Path Params: ueContextId
func CreateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.CreateUEContextRequest) (rsp *models.CreateUEContextResponse, ersp *models.UeContextCreateError, err error) {
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
		rsp = new(CreateUEContextResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 500:
		ersp = new(UeContextCreateError)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411, 413, 415, 429, 503:
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

// Summary: Namf_Communication ReleaseUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/release
// Path Template: /ue-contexts/%s/release
// Path Params: ueContextId
func ReleaseUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.UEContextRelease) (err error) {
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
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication RelocateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/relocate
// Path Template: /ue-contexts/%s/relocate
// Path Params: ueContextId
func RelocateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.RelocateUEContextRequest) (rsp *models.UeContextRelocatedData, err error) {
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
		rsp = new(UeContextRelocatedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication CancelRelocateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/cancel-relocate
// Path Template: /ue-contexts/%s/cancel-relocate
// Path Params: ueContextId
func CancelRelocateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.CancelRelocateUEContextRequest) (err error) {
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
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication N1N2 Message Transfer (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages
// Path Template: /ue-contexts/%s/n1-n2-messages
// Path Params: ueContextId
func N1N2MessageTransfer(cli sbi.ConsumerClient, ueContextId string, body *models.N1N2MessageTransferRequest) (rsp *models.N1N2MessageTransferRspData, ersp *models.N1N2MessageTransferError, err error) {
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
		rsp = new(N1N2MessageTransferRspData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 409, 504:
		ersp = new(N1N2MessageTransferError)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication N1N2 Message Subscribe (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions
// Path Template: /ue-contexts/%s/n1-n2-messages/subscriptions
// Path Params: ueContextId
func N1N2MessageSubscribe(cli sbi.ConsumerClient, ueContextId string, body *models.UeN1N2InfoSubscriptionCreateData) (rsp *models.UeN1N2InfoSubscriptionCreatedData, err error) {
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
		rsp = new(UeN1N2InfoSubscriptionCreatedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication EBI Assignment service Operation
// Description:
// Path: /ue-contexts/:ueContextId/assign-ebi
// Path Template: /ue-contexts/%s/assign-ebi
// Path Params: ueContextId
func EBIAssignment(cli sbi.ConsumerClient, ueContextId string, body *models.AssignEbiData) (rsp *models.AssignedEbiData, ersp *models.AssignEbiError, err error) {
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
		rsp = new(AssignedEbiData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 409, 500:
		ersp = new(AssignEbiError)
		response.Body = ersp
		err = cli.DecodeResponse(response)
	case 411, 413, 415, 429, 503:
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

// Summary: Namf_Communication UEContextTransfer service Operation
// Description:
// Path: /ue-contexts/:ueContextId/transfer
// Path Template: /ue-contexts/%s/transfer
// Path Params: ueContextId
func UEContextTransfer(cli sbi.ConsumerClient, ueContextId string, body *models.UEContextTransferRequest) (rsp *models.UEContextTransferResponse, err error) {
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
		rsp = new(UEContextTransferResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId
// Path Template: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId
// Path Params: ueContextId, subscriptionId
type N1N2MessageUnSubscribeParams struct {
	UeContextId    string
	SubscriptionId string
}

func N1N2MessageUnSubscribe(cli sbi.ConsumerClient, params N1N2MessageUnSubscribeParams) (err error) {
	request := sbi.DefaultRequest()
	var response sbi.Response
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication Non UE N2 Info Subscribe service Operation
// Description:
// Path: /non-ue-n2-messages/subscriptions
// Path Template: /non-ue-n2-messages/subscriptions
// Path Params:
func NonUeN2InfoSubscribe(cli sbi.ConsumerClient, body *models.NonUeN2InfoSubscriptionCreateData) (rsp *models.NonUeN2InfoSubscriptionCreatedData, err error) {
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
		rsp = new(NonUeN2InfoSubscriptionCreatedData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 411, 413, 415, 429, 500, 503:
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

// Summary: Namf_Communication AMF Status Change Subscribe service Operation
// Description:
// Path: /subscriptions
// Path Template: /subscriptions
// Path Params:
func AMFStatusChangeSubscribe(cli sbi.ConsumerClient, body *models.SubscriptionData) (rsp *models.SubscriptionData, err error) {
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
		rsp = new(SubscriptionData)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 411, 413, 415, 429, 500, 503:
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
