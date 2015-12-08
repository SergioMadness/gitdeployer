package models

type Response struct {
	Result        int
	ResultMessage string
	Data          map[string]interface{}
}

func CreateResponse(resultCode int, resultMessage string, data map[string]interface{}) *Response {
	result := new(Response)

	result.Result = resultCode
	result.ResultMessage = resultMessage
	result.Data = data

	return result
}
