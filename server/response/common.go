package response

type CommonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func CreateResponse(code int, message string, data interface{}) CommonResponse {
	return CommonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

type CommonResponseWithoutData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func CreateResponseWithoutData(code int, message string) CommonResponseWithoutData {
	return CommonResponseWithoutData{
		Code:    code,
		Message: message,
	}
}
