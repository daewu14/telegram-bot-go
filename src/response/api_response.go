package response

func Api() ApiResponse {
	return ApiResponse{}
}

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Method  string      `json:"method"`
	Data    interface{} `json:"data"`
}

func (ApiResponse) Error(message string, method string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  false,
		Message: message,
		Method:  method,
		Data:    data,
	}
}

func (ApiResponse) Success(message string, method string, data interface{}) ApiResponse {
	return ApiResponse{
		Status:  true,
		Message: message,
		Method:  method,
		Data:    data,
	}
}
