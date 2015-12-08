package models

type Response struct {
	Result        int
	ResultMessage string
	Data          map[string]interface{}
}
