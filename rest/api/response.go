package api

type IResponse interface {
	IsSuccess() bool
	GetCode() string
	GetMessage() string
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) IsSuccess() bool {
	return r.Code == "0"
}

func (r *Response) GetCode() string {
	return r.Code
}

func (r *Response) GetMessage() string {
	return r.Message
}
