package common

type APIResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data"` // Gunakan interface{} untuk data yang fleksibel
}

func BuildResponse(statusCode int, data interface{}) (int, *APIResponse) {
	return statusCode, &APIResponse{
		Message:    "Success",
		StatusCode: statusCode,
		Data:       data,
	}
}

// func BuildFailedResponse(statusCode int, data interface{}) *APIResponse {
// 	return &APIResponse{
// 		Message:    "Failed",
// 		StatusCode: statusCode,
// 		Data:       data,
// 	}
// }
