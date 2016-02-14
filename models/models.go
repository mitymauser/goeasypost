package models


type EasyPostError struct {
	Code    string            `json:"code"`
	Errors  []interface{}     `json:"errors"`
	Message string            `json:"message"`
}

type TrackerList struct {
	ErrorResponse *EasyPostError  `json:"error"`
	Trackers      []interface{}  `json:"trackers"`
	HasMore       bool           `json:"has_more"`
}