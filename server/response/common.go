package response

type CommonResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponse[T any](code int, message string, data T) CommonResponse[T] {
	return CommonResponse[T]{
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
