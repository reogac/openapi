package NudmSDM

func OnGetNSSAI(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetNSSAI()
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
func OnGetTraceConfigData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetTraceConfigData()
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
func OnGetLcsPrivacyData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetLcsPrivacyData()
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
func OnUpuAck(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body := new(models.AcknowledgeInfo)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleUpuAck(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnGetGroupIdentifiers(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetGroupIdentifiers()
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
func OnGetDataSets(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetDataSets()
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
func OnGetUeCtxInAmfData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetUeCtxInAmfData()
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
func OnGetAmData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetAmData()
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
func OnGetSmfSelData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSmfSelData()
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
func OnGetLcsBcaData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetLcsBcaData()
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
func OnCAGAck(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body := new(models.AcknowledgeInfo)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleCAGAck(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnGetIndividualSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetIndividualSharedData()
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
func OnGetMultipleIdentifiers(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetMultipleIdentifiers()
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
func OnGetSmsData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSmsData()
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
func OnGetSmsMngtData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSmsMngtData()
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
func OnGetLcsMoData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetLcsMoData()
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
func OnGetV2xData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetV2xData()
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
func OnSorAckInfo(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body := new(models.AcknowledgeInfo)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleSorAckInfo(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnGetUeCtxInSmsfData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetUeCtxInSmsfData()
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
func OnUpdateSORInfo(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.SorUpdateInfo
	if ctx.HaveRequestBody() {
		body := new(models.SorUpdateInfo)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleUpdateSORInfo(body)
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
func OnUnsubscribeForSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleUnsubscribeForSharedData()
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnUnsubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleUnsubscribe()
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnGetEcrData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetEcrData()
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
func OnGetSmData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSmData()
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
func OnSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.SdmSubscription)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleSubscribe(body)
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
func OnSNSSAIsAck(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	var body *models.AcknowledgeInfo
	if ctx.HaveRequestBody() {
		body := new(models.AcknowledgeInfo)
		err = ctx.DecodeRequest(body)
	}
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		prob := HandleSNSSAIsAck(body)
		if prob != nil {
			response.SetBody(models.GetProblemDetailCode(prob), prob)
			return
		}
	}
	return
}
func OnGetSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSharedData()
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
func OnSubscribeToSharedData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	body := new(models.SdmSubscription)
	err = ctx.DecodeRequest(body)
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleSubscribeToSharedData(body)
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
func OnGetUeCtxInSmfData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetUeCtxInSmfData()
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
func OnGetProseData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetProseData()
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
func OnGetMbsData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetMbsData()
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
func OnGetUcData(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetUcData()
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
func OnGetSupiOrGpsi(ctx sbi.RequestContext, handler any) (response sbi.Response) {
	prod := handler.(Producer)
	var err error
	if err == nil {
		response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
	} else {
		rsp, prob := HandleGetSupiOrGpsi()
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
