package controllers

const FAIL = 0
const SUCCESS = 1

type ApiResponseJson struct {
	Status int `json:"status"`
	Msg    string `json:"msg"`
	Data   interface{} `json:"data"`
}

//api 统一响应
func ApiResponse(status int, msg string, data interface{}) (response *ApiResponseJson) {

	response = &ApiResponseJson{Status: status, Msg: msg, Data: data}
	return
}
