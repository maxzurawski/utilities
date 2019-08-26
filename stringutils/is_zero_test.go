package stringutils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. check reflect.TypeOf for float64 if is it working (negative/positive case)
func TestFloat64Type(t *testing.T) {

	// Arrange.

	var f float64
	f = 19.99

	// Act.

	result := reflect.TypeOf(f).String()

	// Assert.

	assert.Equal(t, "float64", result)
}

// 2. check reflect.Zero function for int and float64
func TestFloat64Zero(t *testing.T) {

	// Arrange.

	var f float64
	f = 9.99

	// Act.

	result := reflect.Zero(reflect.TypeOf(f)).Interface()

	// Assert.

	assert.Equal(t, float64(0), result)
}

func TestIntZero(t *testing.T) {

	// Arrange.

	var i int
	i = 8

	// Act.

	result := reflect.Zero(reflect.TypeOf(i)).Interface()

	// Assert.

	assert.Equal(t, int(0), result)
}

// 3. check IsZero func with string variable: "Hello utilities", ""
func TestStringZeroForNotEmptyString(t *testing.T) {

	// Arrange.

	s := "Hello utilities"

	// Act.

	result := IsZero(s)

	// Assert.

	assert.False(t, result)
}

func TestStringZeroForEmptyString(t *testing.T) {

	// Arrange.

	s := ""

	// Act.

	result := IsZero(s)

	// Assert.

	assert.True(t, result)
}

type person struct {
	firstname string
	lastname  string
}

func TestPersonTypeOf(t *testing.T) {

	// Arrange.

	p := person{firstname: "Max"}

	// Act.

	result := reflect.TypeOf(p).String()

	// Assert.

	assert.Equal(t, "stringutils.person", result)
}

func TestPersonZero(t *testing.T) {

	// Arrange.

	p := person{firstname: "Max"}

	// Act.

	result := reflect.Zero(reflect.TypeOf(&p)).Interface()

	// Assert.

	assert.Equal(t, "<nil>", fmt.Sprintf("%v", result))
}
