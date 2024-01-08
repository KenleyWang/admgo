package response

import (
	"context"
	"reflect"
)

type okRsp struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func OKResponse(ctx context.Context, v any) any {
	//判断是否是空结构体，是的话将v置为nil
	t_v := reflect.TypeOf(v)
	if t_v.Kind() == reflect.Ptr {
		if t_v.Elem().Kind() == reflect.Struct {
			if t_v.Elem().NumField() == 0 {
				v = nil
			}
		}
	}
	//判断是否是空接口类型，是的话将v置为nil
	if t_v.Kind() == reflect.Interface {
		v = nil
	}
	return &okRsp{
		Code: 0,
		Msg:  "success",
		Data: v,
	}
}
