/*
This file is generated with a SBI APIs generator tool developed by ETRI
Generated at Fri Nov 15 22:12:00 KST 2024 by TungTQ<tqtung@etri.re.kr>
Do not modify
*/

package NsmfPDUSession

import (
	"net/http"
	"sbi"
)

var _routes = []sbi.Route{
	{
		Label:   "ReleasePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/release",
		Handler: "OnReleasePduSession",
	},
	{
		Label:   "RetrieveSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/retrieve",
		Handler: "OnRetrieveSmContext",
	},
	{
		Label:   "UpdateSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/modify",
		Handler: "OnUpdateSmContext",
	},
	{
		Label:   "ReleaseSmContext",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/release",
		Handler: "OnReleaseSmContext",
	},
	{
		Label:   "UpdatePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/modify",
		Handler: "OnUpdatePduSession",
	},
	{
		Label:   "TransferMoData",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/transfer-mo-data",
		Handler: "OnTransferMoData",
	},
	{
		Label:   "PostSmContexts",
		Method:  http.MethodPost,
		Path:    "/sm-contexts",
		Handler: "OnPostSmContexts",
	},
	{
		Label:   "SendMoData",
		Method:  http.MethodPost,
		Path:    "/sm-contexts/:smContextRef/send-mo-data",
		Handler: "OnSendMoData",
	},
	{
		Label:   "PostPduSessions",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions",
		Handler: "OnPostPduSessions",
	},
	{
		Label:   "RetrievePduSession",
		Method:  http.MethodPost,
		Path:    "/pdu-sessions/:pduSessionRef/retrieve",
		Handler: "OnRetrievePduSession",
	},
}

func Service(p Producer) sbi.Service {
	return sbi.Service{
		Group:   PATH_ROOT,
		Routes:  _routes,
		Handler: p,
	}
}
