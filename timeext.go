package gotsrpc

import (
	"time"
)

type TimeExt struct{}

var timeExt = &TimeExt{}

func (x *TimeExt) ConvertExt(v interface{}) interface{} {
	return v.(*time.Time).UnixMilli()
}

func (x *TimeExt) UpdateExt(dest interface{}, src interface{}) {
	tt := dest.(*time.Time)
	*tt = time.Unix(0, src.(int64)*int64(time.Millisecond)).Local()
}
