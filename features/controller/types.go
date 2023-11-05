package controller

const (
	NEGATIVE_GAIN   int  = 0x00
	POSITIVE_GAIN   int  = 0x01
	RESPONSE_HEADER byte = 0x2B
	SET_REGISTER    byte = 0xFC
	SET_FILTER      byte = 0xFD
	GET_VERSION     byte = 0xFE
	GET_STATUS      byte = 0xFF
)
