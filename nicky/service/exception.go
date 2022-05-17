package service

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Throw(res *Response) {
	panic(res)
}

func CheckError(err error, res *Response) {
	if err != nil {
		Throw(res)
	}
}
