package models

type PcfUeCallbackInfo struct {
	BindingInfo string `json:"bindingInfo,omitempty"`
	CallbackUri string `json:"callbackUri"`
}
