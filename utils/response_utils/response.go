package response_utils

type Response struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponseWithEmptyData(code int, message string) Response {
	return Response{
		Code:    code,
		Type:    "success",
		Error:   false,
		Message: message,
	}
}

func NewSuccessResponse(code int, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Type:    "success",
		Error:   false,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(code int, messageType string, message string) Response {
	return Response{
		Code:    code,
		Type:    messageType,
		Error:   true,
		Message: message,
	}
}
