package unit

import (
	"bytes"
	"reflect"
)

type CompareType = compareResult
type compareResult int

const (
	CompareLess compareResult = iota - 1
	CompareEqual
	CompareGreater
	CompareError
)

func Compare(a interface{}, b interface{}, kind reflect.Kind) compareResult {
	switch kind {
	case reflect.Int:
		return compareInts(a, b)
	case reflect.Float32, reflect.Float64:
		return compareFloats(a, b)
	case reflect.Slice:
		return compareSlices(a, b)
	default:
		return CompareEqual
	}
}

func compareInts(a, b interface{}) compareResult {
	intObjA, ok := a.(int)
	if !ok {
		return CompareError
	}
	intObjB, ok := b.(int)
	if !ok {
		return CompareError
	}
	switch {
	case intObjA > intObjB:
		return CompareGreater
	case intObjA < intObjB:
		return CompareLess
	default:
		return CompareEqual
	}
}

func compareFloats(a, b interface{}) compareResult {
	floatObjA, ok := a.(float64)
	if !ok {
		return CompareError
	}
	floatObjB, ok := b.(float64)
	if !ok {
		return CompareError
	}
	switch {
	case floatObjA > floatObjB:
		return CompareGreater
	case floatObjA < floatObjB:
		return CompareLess
	default:
		return CompareEqual
	}
}

func compareSlices(a, b interface{}) compareResult {
	if !reflect.ValueOf(a).CanConvert(reflect.TypeOf([]byte{})) {
		return CompareEqual
	}
	bytesObjA, ok := a.([]byte)
	if !ok {
		return CompareError
	}
	bytesObjB, ok := b.([]byte)
	if !ok {
		return CompareError
	}
	return compareResult(bytes.Compare(bytesObjA, bytesObjB))
}

func Nil(param interface{}) bool {
	return param == nil
}

func IsEmpty(object interface{}) bool {
	if object == nil {
		return true
	}
	objValue := reflect.ValueOf(object)
	switch objValue.Kind() {
	case reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	default:
		return false
	}
}

func CheckFunctionResult(fn interface{}, params []interface{}, expectedResult interface{}) bool {
	fnValue := reflect.ValueOf(fn)
	if fnValue.Kind() != reflect.Func {
		return false
	}

	in := make([]reflect.Value, len(params))
	for i, param := range params {
		in[i] = reflect.ValueOf(param)
	}

	results := fnValue.Call(in)
	if len(results) != 1 {
		return false
	}

	return reflect.DeepEqual(results[0].Interface(), expectedResult)
}
