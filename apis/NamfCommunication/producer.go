package NamfCommunication
func OnCreateUEContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.CreateUEContextRequest)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,ersp,prob := HandleCreateUEContext(body)
if rsp != nil {
response.SetBody(201, rsp)
return
}
if ersp != nil {
response.SetBody(models.StatusFromUeContextCreateError(ersp), ersp)
return
}
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnNonUeN2MessageTransfer(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.NonUeN2MessageTransferRequest)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,ersp,prob := HandleNonUeN2MessageTransfer(body)
if rsp != nil {
response.SetBody(200, rsp)
return
}
if ersp != nil {
response.SetBody(models.StatusFromN2InformationTransferError(ersp), ersp)
return
}
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnNonUeN2InfoSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.NonUeN2InfoSubscriptionCreateData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleNonUeN2InfoSubscribe(body)
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
func OnAMFStatusChangeSubscribeModfy(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.SubscriptionData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleAMFStatusChangeSubscribeModfy(body)
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
func OnReleaseUEContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.UEContextRelease)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
prob := HandleReleaseUEContext(body)
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnUEContextTransfer(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.UEContextTransferRequest)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleUEContextTransfer(body)
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
func OnRelocateUEContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.RelocateUEContextRequest)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleRelocateUEContext(body)
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
func OnN1N2MessageUnSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
prob := HandleN1N2MessageUnSubscribe()
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnEBIAssignment(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.AssignEbiData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,ersp,prob := HandleEBIAssignment(body)
if rsp != nil {
response.SetBody(200, rsp)
return
}
if ersp != nil {
response.SetBody(models.StatusFromAssignEbiError(ersp), ersp)
return
}
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnCancelRelocateUEContext(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.CancelRelocateUEContextRequest)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
prob := HandleCancelRelocateUEContext(body)
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnN1N2MessageTransfer(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.N1N2MessageTransferRequest)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,ersp,prob := HandleN1N2MessageTransfer(body)
if rsp != nil {
response.SetBody(200, rsp)
return
}
if ersp != nil {
response.SetBody(models.StatusFromN1N2MessageTransferError(ersp), ersp)
return
}
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnNonUeN2InfoUnSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
prob := HandleNonUeN2InfoUnSubscribe()
if prob != nil {
response.SetBody(models.GetProblemDetailCode(prob), prob)
return
}
}
return
}
func OnAMFStatusChangeSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.SubscriptionData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleAMFStatusChangeSubscribe(body)
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
func OnRegistrationStatusUpdate(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.UeRegStatusUpdateReqData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleRegistrationStatusUpdate(body)
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
func OnN1N2MessageSubscribe(ctx sbi.RequestContext, handler any) (response sbi.Response) {
prod := handler.(Producer)
var err error
body := new(models.UeN1N2InfoSubscriptionCreateData)
err = ctx.DecodeRequest(body)
if err == nil {
response.SetBody(400, models.NewSimpleProblem(400, err.Error()))
}else{
rsp,prob := HandleN1N2MessageSubscribe(body)
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
