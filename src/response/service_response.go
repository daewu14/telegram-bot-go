package response

func Service() ServiceResponse {
	return ServiceResponse{}
}

type ServiceResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (s ServiceResponse) Error(error error, data interface{}) ServiceResponse {
	println("error :",error.Error())
	return ServiceResponse{
		Status:  false,
		Message: error.Error(),
		Data:    data,
	}
}

func (ServiceResponse) Success(message string, data interface{}) ServiceResponse {
	return ServiceResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}