package models

type ExceptionTrend string

// Define constant values for ExceptionTrend
const (
	EXCEPTIONTREND_UP     ExceptionTrend = "UP"
	EXCEPTIONTREND_DOWN   ExceptionTrend = "DOWN"
	EXCEPTIONTREND_UNKNOW ExceptionTrend = "UNKNOW"
	EXCEPTIONTREND_STABLE ExceptionTrend = "STABLE"
)
