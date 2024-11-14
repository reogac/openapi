/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Thu Nov 14 22:22:55 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NudmUECM

func OnGetLocationInfo(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetLocationInfo()
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
func OnGetNon3GppRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetNon3GppRegistration()
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
func OnGet3GppRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGet3GppRegistration()
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
func OnRetrieveSmfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleRetrieveSmfRegistration()
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
func OnTriggerPCSCFRestoration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.TriggerRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleTriggerPCSCFRestoration(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnNwdafRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.NwdafRegistration)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleNwdafRegistration(body)
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
func OnGetNon3GppSmsfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetNon3GppSmsfRegistration()
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
func OnGetRegistrations(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetRegistrations()
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
func OnSendRoutingInfoSm(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.RoutingInfoSmRequest)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleSendRoutingInfoSm(body)
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
func OnDeregAMF(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.AmfDeregInfo)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleDeregAMF(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnPeiUpdate(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.PeiUpdateInfo)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandlePeiUpdate(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnUpdateRoamingInformation(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.RoamingInfoUpdate)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleUpdateRoamingInformation(body)
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
func OnGetSmfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSmfRegistration()
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
func OnGet3GppSmsfRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGet3GppSmsfRegistration()
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
func OnGetIpSmGwRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetIpSmGwRegistration()
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
func OnGetNwdafRegistration(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetNwdafRegistration()
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
