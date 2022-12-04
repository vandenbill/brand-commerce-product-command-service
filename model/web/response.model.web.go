package web

type emptObj struct{}

type Response struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	ID      interface{} `json:"data"`
}

func BuildResponse(message string, data interface{}) Response {
	res := Response{
		Message: message,
		Errors:  emptObj{},
		ID:      data,
	}
	return res
}

func BuildErrorResponse(message string, err ...string) Response {
	res := Response{
		Message: message,
		Errors:  err,
		ID:      emptObj{},
	}
	return res
}
