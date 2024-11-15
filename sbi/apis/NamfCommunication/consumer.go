/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:01 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NamfCommunication

import (
	"net/http"
	"sbi"
	"sbi/models"
)

const (
	PATH_ROOT string = "namf-comm/v1"
)

// Summary: Namf_Communication UEContextTransfer service Operation
// Description:
// Path: /ue-contexts/:ueContextId/transfer
// Path Params: ueContextId
func UEContextTransfer(cli sbi.ConsumerClient, ueContextId string, body *models.UEContextTransferRequest) (rsp *models.UEContextTransferResponse, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/transfer", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication RegistrationStatusUpdate service Operation
// Description:
// Path: /ue-contexts/:ueContextId/transfer-update
// Path Params: ueContextId
func RegistrationStatusUpdate(cli sbi.ConsumerClient, ueContextId string, body *models.UeRegStatusUpdateReqData) (rsp *models.UeRegStatusUpdateRspData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/transfer-update", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication RelocateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/relocate
// Path Params: ueContextId
func RelocateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.RelocateUEContextRequest) (rsp *models.UeContextRelocatedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/relocate", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication N1N2 Message Subscribe (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions
// Path Params: ueContextId
func N1N2MessageSubscribe(cli sbi.ConsumerClient, ueContextId string, body *models.UeN1N2InfoSubscriptionCreateData) (rsp *models.UeN1N2InfoSubscriptionCreatedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/n1-n2-messages/subscriptions", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication Non UE N2 Message Transfer service Operation
// Description:
// Path: /non-ue-n2-messages/transfer
// Path Params:
func NonUeN2MessageTransfer(cli sbi.ConsumerClient, body *models.NonUeN2MessageTransferRequest) (rsp *models.N2InformationTransferRspData, ersp *models.N2InformationTransferError, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/non-ue-n2-messages/transfer", PATH_ROOT)
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

// Summary: Namf_Communication ReleaseUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/release
// Path Params: ueContextId
func ReleaseUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.UEContextRelease) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/release", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication Non UE N2 Info UnSubscribe service Operation
// Description:
// Path: /non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId
// Path Params: n2NotifySubscriptionId
func NonUeN2InfoUnSubscribe(cli sbi.ConsumerClient, n2NotifySubscriptionId string) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(n2NotifySubscriptionId) == 0 {
		err = fmt.Errorf("n2NotifySubscriptionId is required")
		return
	}

	request.Path = fmt.Sprintf("%s/non-ue-n2-messages/subscriptions/:n2NotifySubscriptionId", PATH_ROOT, n2NotifySubscriptionId)
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

// Summary: Namf_Communication AMF Status Change Subscribe service Operation
// Description:
// Path: /subscriptions
// Path Params:
func AMFStatusChangeSubscribe(cli sbi.ConsumerClient, body *models.SubscriptionData) (rsp *models.SubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/subscriptions", PATH_ROOT)
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

// Summary: Namf_Communication AMF Status Change Subscribe Modify service Operation
// Description:
// Path: /subscriptions/:subscriptionId
// Path Params: subscriptionId
func AMFStatusChangeSubscribeModfy(cli sbi.ConsumerClient, subscriptionId string, body *models.SubscriptionData) (rsp *models.SubscriptionData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(subscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/subscriptions/%s", PATH_ROOT, subscriptionId)
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

// Summary: Namf_Communication N1N2 Message Transfer (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages
// Path Params: ueContextId
func N1N2MessageTransfer(cli sbi.ConsumerClient, ueContextId string, body *models.N1N2MessageTransferRequest) (rsp *models.N1N2MessageTransferRspData, ersp *models.N1N2MessageTransferError, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/n1-n2-messages", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication CancelRelocateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId/cancel-relocate
// Path Params: ueContextId
func CancelRelocateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.CancelRelocateUEContextRequest) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/cancel-relocate", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication N1N2 Message UnSubscribe (UE Specific) service Operation
// Description:
// Path: /ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId
// Path Params: ueContextId, subscriptionId
type N1N2MessageUnSubscribeParams struct {
	UeContextId    string
	SubscriptionId string
}

func N1N2MessageUnSubscribe(cli sbi.ConsumerClient, params N1N2MessageUnSubscribeParams) (err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(params.UeContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}

	if len(params.SubscriptionId) == 0 {
		err = fmt.Errorf("subscriptionId is required")
		return
	}

	request.Path = fmt.Sprintf("%s/ue-contexts/:ueContextId/n1-n2-messages/subscriptions/:subscriptionId", PATH_ROOT, params.UeContextId, params.SubscriptionId)
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

// Summary: Namf_Communication EBI Assignment service Operation
// Description:
// Path: /ue-contexts/:ueContextId/assign-ebi
// Path Params: ueContextId
func EBIAssignment(cli sbi.ConsumerClient, ueContextId string, body *models.AssignEbiData) (rsp *models.AssignedEbiData, ersp *models.AssignEbiError, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s/assign-ebi", PATH_ROOT, ueContextId)
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

// Summary: Namf_Communication Non UE N2 Info Subscribe service Operation
// Description:
// Path: /non-ue-n2-messages/subscriptions
// Path Params:
func NonUeN2InfoSubscribe(cli sbi.ConsumerClient, body *models.NonUeN2InfoSubscriptionCreateData) (rsp *models.NonUeN2InfoSubscriptionCreatedData, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodPost
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/non-ue-n2-messages/subscriptions", PATH_ROOT)
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

// Summary: Namf_Communication CreateUEContext service Operation
// Description:
// Path: /ue-contexts/:ueContextId
// Path Params: ueContextId
func CreateUEContext(cli sbi.ConsumerClient, ueContextId string, body *models.CreateUEContextRequest) (rsp *models.CreateUEContextResponse, ersp *models.UeContextCreateError, err error) {

	request := sbi.DefaultRequest()

	request.Method = http.MethodGet
	if len(ueContextId) == 0 {
		err = fmt.Errorf("ueContextId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request.Body = body

	request.Path = fmt.Sprintf("%s/ue-contexts/%s", PATH_ROOT, ueContextId)
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
