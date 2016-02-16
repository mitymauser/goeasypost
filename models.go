package goeasypost


type EasyPostError struct {
	Code    string            `json:"code"`
	Errors  []interface{}     `json:"errors"`
	Message string            `json:"message"`
}

