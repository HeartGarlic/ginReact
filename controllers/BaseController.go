package controllers

const SUCCESS_CODE = 0
const ERROR_CODE = 1

type BaseController struct {
}

type Response struct {
	Code int `json:"code"`
    Msg string `json:"msg"`
    Data interface{} `json:"data"`
}

// Success 返回成功的方法
func (b *BaseController) Success(data interface{}) Response {
	return Response{
        Code: SUCCESS_CODE,
        Msg: "success",
        Data: data,
    }
}

// Error 返回错误的方法
func (b *BaseController) Error(msg string) Response{
	return Response{
		Code: ERROR_CODE,
		Msg:  msg,
		Data: nil,
	}
}

