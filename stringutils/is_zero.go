package stringutils

import (
	"reflect"
)

func IsZero(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

/*

1. check reflect.TypeOf for float64 if is it working (negative/positive case)
2. check reflect.Zero function for int and float64
3. check IsZero func with string variable: "Hello utilities", ""
4*. check reflect.type for person struct (containing two fields firstname and lastname)
5*. check reflect.zero for person object created from type person struct

*/
