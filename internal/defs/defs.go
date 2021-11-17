package defs

import (
	"errors"
	"net/http"
	"syscall"
)

// NewError 전달받은 텍스트 형식에 해당하는 에러 반환
func NewError(arg interface{}) error {
	var str string

	// _type_assertion_(타입 선언) interface{} → 타입 변환
	switch arg.(type) { // _type_switch
	case syscall.Errno:
		str = arg.(syscall.Errno).Error()
	case string:
		str = arg.(string)
	case int:
		str = http.StatusText(arg.(int))
	default:
		panic(syscall.EINVAL.Error())
	}

	return errors.New(str)
}
