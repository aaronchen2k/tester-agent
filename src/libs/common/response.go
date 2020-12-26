package common

type Response struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data"`

	PageSize   int `json:"pageSize"`
	PageNo     int `json:"pageNo"`
	TotalPage  int `json:"totalPage"`
	TotalCount int `json:"totalCount"`
}

type Lists struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

func ApiResource(code int64, objects interface{}, msg string) (r *Response) {
	r = &Response{Code: code, Data: objects, Msg: msg}
	return
}
