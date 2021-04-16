// Package dto
package dto

import "EduMall/config"

// ResponseJson ...
type ResponseJson struct {
	St   config.ErrorCode `json:"st"`
	Msg  string           `json:"msg"`
	Data interface{}      `json:"data"`
}

type DataList struct {
	List  interface{} `json:"list,omitempty"`
	Total int64       `json:"total,omitempty"`
}

type EmptyObj struct{}

func NewResponse(data interface{}) *ResponseJson {
	return &ResponseJson{Data: data}
}

func NewResponseWithStatusData(st config.ErrorCode, msg string, data interface{}) *ResponseJson {
	return &ResponseJson{St: st, Msg: msg, Data: data}
}

func NewResponseWithTotal(data interface{}, total int64) *ResponseJson {
	return &ResponseJson{
		Data: DataList{
			List:  data,
			Total: total,
		}}
}
