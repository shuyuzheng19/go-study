package common

type ResultResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Other   interface{} `json:"other_info,omitempty"`
}

var (
	success = ResultResponse{Code: 200, Message: "成功"}
	failure = ResultResponse{Code: 10001, Message: "失败"}
	error   = ResultResponse{Code: 500, Message: "服务器错误"}
)

func BuildSuccess(data interface{}) ResultResponse {
	var result = success
	result.Data = data
	return result
}

func BuildDefaultFailure() ResultResponse {
	return failure
}

func BuildFailure(message string) ResultResponse {
	var result = failure
	result.Message = message
	return result
}

func BuildDefaultSuccess() ResultResponse {
	return success
}

func BuildDefaultError(message string) ResultResponse {
	return error
}

func BuildError(message string) ResultResponse {
	var result = error
	result.Message = message
	return result
}
