package models

type Response struct {
	Code      int         `json:"code" example:"200"` // 代码
	Data      interface{} `json:"data"`               // 数据集
	Msg       string      `json:"msg"`                // 消息
	RequestId string      `json:"requestId"`          //请求序列号
}

type Page struct {
	List   interface{} `json:"list"`
	Total  int         `json:"total"`
	Page   int         `json:"page"`
	Offset int         `json:"offset"`
}

// ReturnOK 正常返回
func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

// ReturnError 错误返回
func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}
