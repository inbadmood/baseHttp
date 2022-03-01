package utils

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

const DatetimeFormat = "2006-01-02 15:04:05"

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func GetLogTimeFormat(t time.Time) string {
	return t.Format(DatetimeFormat)
}

func CheckInArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func InterfaceToString(inter interface{}) string {
	var ret string
	switch codeType := inter.(type) {
	case int:
		ret = strconv.Itoa(codeType)
	case string:
		ret = codeType
	case float64:
		ret = strconv.FormatFloat(codeType, 'g', 10, 64)
	default:
		ret = fmt.Sprintf("%v", inter)

	}
	return ret
}
