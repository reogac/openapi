/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmUECM

const (
	PATH_ROOT string = "nudm-uecm/v1"
)

// Summary: retrieve the target UE's location information
// Description:
// Path: /:ueId/registrations/location
// Path Params: ueId
type GetLocationInfoParams struct {
	UeId              string
	SupportedFeatures string
}

func GetLocationInfo(cli sbi.ConsumerClient, params GetLocationInfoParams) (rsp *models.LocationInfo, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/location", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(LocationInfo)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve the AMF registration for non-3GPP access information
// Description:
// Path: /:ueId/registrations/amf-non-3gpp-access
// Path Params: ueId
type GetNon3GppRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func GetNon3GppRegistration(cli sbi.ConsumerClient, params GetNon3GppRegistrationParams) (rsp *models.AmfNon3GppAccessRegistration, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/amf-non-3gpp-access", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(AmfNon3GppAccessRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve the AMF registration for 3GPP access information
// Description:
// Path: /:ueId/registrations/amf-3gpp-access
// Path Params: ueId
type Get3GppRegistrationParams struct {
	SupportedFeatures string
	UeId              string
}

func Get3GppRegistration(cli sbi.ConsumerClient, params Get3GppRegistrationParams) (rsp *models.Amf3GppAccessRegistration, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/amf-3gpp-access", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(Amf3GppAccessRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: get an SMF registration
// Description:
// Path: /:ueId/registrations/smf-registrations/:pduSessionId
// Path Params: ueId, pduSessionId
type RetrieveSmfRegistrationParams struct {
	UeId         string
	PduSessionId int
}

func RetrieveSmfRegistration(cli sbi.ConsumerClient, params RetrieveSmfRegistrationParams) (rsp *models.SmfRegistration, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/:ueId/registrations/smf-registrations/%s", PATH_ROOT, params.UeId, params.PduSessionId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmfRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: Trigger the Restoration of the P-CSCF
// Description:
// Path: /restore-pcscf
// Path Params:
func TriggerPCSCFRestoration(cli sbi.ConsumerClient, body *models.TriggerRequest) (err error) {
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/restore-pcscf", PATH_ROOT)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 403, 404, 500, 501, 503:
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

// Summary: register as NWDAF
// Description:
// Path: /:ueId/registrations/nwdaf-registrations/:nwdafRegistrationId
// Path Params: ueId, nwdafRegistrationId
type NwdafRegistrationParams struct {
	UeId                string
	NwdafRegistrationId string
}

func NwdafRegistration(cli sbi.ConsumerClient, params NwdafRegistrationParams, body *models.NwdafRegistration) (rsp *models.NwdafRegistration, err error) {
	if len(params.NwdafRegistrationId) == 0 {
		err = fmt.Errorf("nwdafRegistrationId is required")
		return
	}

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/:ueId/registrations/nwdaf-registrations/%s", PATH_ROOT, params.UeId, params.NwdafRegistrationId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(NwdafRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve the SMSF registration for non-3GPP access information
// Description:
// Path: /:ueId/registrations/smsf-non-3gpp-access
// Path Params: ueId
type GetNon3GppSmsfRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func GetNon3GppSmsfRegistration(cli sbi.ConsumerClient, params GetNon3GppSmsfRegistrationParams) (rsp *models.SmsfRegistration, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/smsf-non-3gpp-access", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmsfRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve UE registration data sets
// Description:
// Path: /:ueId/registrations
// Path Params: ueId
type GetRegistrationsParams struct {
	UeId                     string
	SupportedFeatures        string
	RegistrationDatasetNames []string
	SingleNssai              *Snssai
	Dnn                      string
}

func GetRegistrations(cli sbi.ConsumerClient, params GetRegistrationsParams) (rsp *models.RegistrationDataSets, err error) {
	if len(params.RegistrationDatasetNames) == 0 {
		err = fmt.Errorf("registration-dataset-names is required")
		return
	}

	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(RegistrationDataSets)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: Retreive addressing information for SMS delivery
// Description:
// Path: /:ueId/registrations/send-routing-info-sm
// Path Params: ueId
func SendRoutingInfoSm(cli sbi.ConsumerClient, ueId string, body *models.RoutingInfoSmRequest) (rsp *models.RoutingInfoSmResponse, err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/send-routing-info-sm", PATH_ROOT, ueId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(RoutingInfoSmResponse)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: trigger AMF for 3GPP access deregistration
// Description:
// Path: /:ueId/registrations/amf-3gpp-access/dereg-amf
// Path Params: ueId
func DeregAMF(cli sbi.ConsumerClient, ueId string, body *models.AmfDeregInfo) (err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/amf-3gpp-access/dereg-amf", PATH_ROOT, ueId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 403, 404, 500, 503:
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

// Summary: Updates the PEI in the 3GPP access registration context
// Description:
// Path: /:ueId/registrations/amf-3gpp-access/pei-update
// Path Params: ueId
func PeiUpdate(cli sbi.ConsumerClient, ueId string, body *models.PeiUpdateInfo) (err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/amf-3gpp-access/pei-update", PATH_ROOT, ueId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 204:
		return
	case 400, 403, 404, 500, 503:
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

// Summary: Update the Roaming Information
// Description:
// Path: /:ueId/registrations/amf-3gpp-access/roaming-info-update
// Path Params: ueId
func UpdateRoamingInformation(cli sbi.ConsumerClient, ueId string, body *models.RoamingInfoUpdate) (rsp *models.RoamingInfoUpdate, err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	if body == nil {
		err = fmt.Errorf("body is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/amf-3gpp-access/roaming-info-update", PATH_ROOT, ueId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 201:
		rsp = new(RoamingInfoUpdate)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 204:
		return
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve the SMF registration information
// Description:
// Path: /:ueId/registrations/smf-registrations
// Path Params: ueId
type GetSmfRegistrationParams struct {
	SupportedFeatures string
	UeId              string
	SingleNssai       *Snssai
	Dnn               string
}

func GetSmfRegistration(cli sbi.ConsumerClient, params GetSmfRegistrationParams) (rsp *models.SmfRegistrationInfo, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/smf-registrations", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmfRegistrationInfo)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve the SMSF registration for 3GPP access information
// Description:
// Path: /:ueId/registrations/smsf-3gpp-access
// Path Params: ueId
type Get3GppSmsfRegistrationParams struct {
	UeId              string
	SupportedFeatures string
}

func Get3GppSmsfRegistration(cli sbi.ConsumerClient, params Get3GppSmsfRegistrationParams) (rsp *models.SmsfRegistration, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/smsf-3gpp-access", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(SmsfRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: Retrieve the IP-SM-GW registration information
// Description:
// Path: /:ueId/registrations/ip-sm-gw
// Path Params: ueId
func GetIpSmGwRegistration(cli sbi.ConsumerClient, ueId string) (rsp *models.IpSmGwRegistration, err error) {
	if len(ueId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/ip-sm-gw", PATH_ROOT, ueId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(IpSmGwRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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

// Summary: retrieve the NWDAF registration
// Description:
// Path: /:ueId/registrations/nwdaf-registrations
// Path Params: ueId
type GetNwdafRegistrationParams struct {
	UeId              string
	AnalyticsIds      []string
	SupportedFeatures string
}

func GetNwdafRegistration(cli sbi.ConsumerClient, params GetNwdafRegistrationParams) (rsp *models.NwdafRegistration, err error) {
	if len(params.UeId) == 0 {
		err = fmt.Errorf("ueId is required")
		return
	}
	request := sbi.DefaultRequest()
	var response sbi.Response
	request.Path = fmt.Sprintf("%s/%s/registrations/nwdaf-registrations", PATH_ROOT, params.UeId)
	if response, err = cli.Send(request); err != nil {
		return
	}
	switch response.StatusCode {
	case 200:
		rsp = new(NwdafRegistration)
		response.Body = rsp
		err = cli.DecodeResponse(response)
	case 400, 403, 404, 500, 503:
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
