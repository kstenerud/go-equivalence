package equivalence

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/kstenerud/go-describe"
)

func assertEquivalent(t *testing.T, a, b interface{}) {
	if !IsEquivalent(a, b) {
		t.Errorf("Expected %v (%v) and %v (%v) to be equivalent", describe.D(a), reflect.TypeOf(a), describe.D(b), reflect.TypeOf(b))
	}
}

func assertNotEquivalent(t *testing.T, a, b interface{}) {
	if IsEquivalent(a, b) {
		t.Errorf("Expected %v (%v) and %v (%v) to not be equivalent", describe.D(a), reflect.TypeOf(a), describe.D(b), reflect.TypeOf(b))
	}
}

type MyStruct struct {
	IntVal    int
	StringVal string
}

type ComplexStruct struct {
	Map     map[interface{}]interface{}
	Struct  MyStruct
	StructP *MyStruct
}

func TestSameTypesEqual(t *testing.T) {
	var nilValue *int
	ptrValue := new(interface{})
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	struct1 := MyStruct{1, "test"}
	struct2 := MyStruct{1, "test"}
	map1 := map[string]int{"a": 1}
	map2 := map[string]int{"a": 1}
	assertEquivalent(t, nilValue, nilValue)
	assertEquivalent(t, true, true)
	assertEquivalent(t, int8(100), int8(100))
	assertEquivalent(t, int16(100), int16(100))
	assertEquivalent(t, int32(100), int32(100))
	assertEquivalent(t, int64(100), int64(100))
	assertEquivalent(t, int(100), int(100))
	assertEquivalent(t, uint8(100), uint8(100))
	assertEquivalent(t, uint16(100), uint16(100))
	assertEquivalent(t, uint32(100), uint32(100))
	assertEquivalent(t, uint64(100), uint64(100))
	assertEquivalent(t, uint(100), uint(100))
	assertEquivalent(t, float32(100), float32(100))
	assertEquivalent(t, float64(100), float64(100))
	assertEquivalent(t, complex(10, 15), complex(10, 15))
	assertEquivalent(t, ptrValue, ptrValue)
	assertEquivalent(t, "test", "test")
	assertEquivalent(t, []byte{1, 2, 3}, []byte{1, 2, 3})
	assertEquivalent(t, []interface{}{1, "blah", 3.8}, []interface{}{1, "blah", 3.8})
	assertEquivalent(t, array1, array2)
	assertEquivalent(t, map1, map2)
	assertEquivalent(t, struct1, struct2)
}

func TestSameTypesNotEqual(t *testing.T) {
	var nilValue *int
	intValue := 1
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 1}
	struct1 := MyStruct{1, "test"}
	struct2 := MyStruct{2, "test"}
	map1 := map[string]int{"a": 1}
	map2 := map[string]int{"a": 2}
	assertNotEquivalent(t, nilValue, &intValue)
	assertNotEquivalent(t, true, false)
	assertNotEquivalent(t, int8(100), int8(101))
	assertNotEquivalent(t, int16(100), int16(101))
	assertNotEquivalent(t, int32(100), int32(101))
	assertNotEquivalent(t, int64(100), int64(101))
	assertNotEquivalent(t, int(100), int(101))
	assertNotEquivalent(t, uint8(100), uint8(101))
	assertNotEquivalent(t, uint16(100), uint16(101))
	assertNotEquivalent(t, uint32(100), uint32(101))
	assertNotEquivalent(t, uint64(100), uint64(101))
	assertNotEquivalent(t, uint(100), uint(101))
	assertNotEquivalent(t, float32(100), float32(101))
	assertNotEquivalent(t, float64(100), float64(101))
	assertNotEquivalent(t, complex(10, 15), complex(11, 15))
	assertNotEquivalent(t, "test", "testing")
	assertNotEquivalent(t, []byte{1, 2, 3}, []byte{1, 2, 5})
	assertNotEquivalent(t, []interface{}{1, "blah", 3.8}, []interface{}{2, "blah", 3.8})
	assertNotEquivalent(t, array1, array2)
	assertNotEquivalent(t, map1, map2)
	assertNotEquivalent(t, struct1, struct2)
}

func TestInt8Equal(t *testing.T) {
	assertEquivalent(t, int8(100), int8(100))
	assertEquivalent(t, int8(100), int16(100))
	assertEquivalent(t, int8(100), int32(100))
	assertEquivalent(t, int8(100), int64(100))
	assertEquivalent(t, int8(100), int(100))
	assertEquivalent(t, int8(100), uint8(100))
	assertEquivalent(t, int8(100), uint16(100))
	assertEquivalent(t, int8(100), uint32(100))
	assertEquivalent(t, int8(100), uint64(100))
	assertEquivalent(t, int8(100), uint(100))
	assertEquivalent(t, int8(100), float32(100))
	assertEquivalent(t, int8(100), float64(100))
}

func TestInt16Equal(t *testing.T) {
	assertEquivalent(t, int16(100), int8(100))
	assertEquivalent(t, int16(100), int16(100))
	assertEquivalent(t, int16(100), int32(100))
	assertEquivalent(t, int16(100), int64(100))
	assertEquivalent(t, int16(100), int(100))
	assertEquivalent(t, int16(100), uint8(100))
	assertEquivalent(t, int16(100), uint16(100))
	assertEquivalent(t, int16(100), uint32(100))
	assertEquivalent(t, int16(100), uint64(100))
	assertEquivalent(t, int16(100), uint(100))
	assertEquivalent(t, int16(100), float32(100))
	assertEquivalent(t, int16(100), float64(100))
}

func TestInt32Equal(t *testing.T) {
	assertEquivalent(t, int32(100), int8(100))
	assertEquivalent(t, int32(100), int16(100))
	assertEquivalent(t, int32(100), int32(100))
	assertEquivalent(t, int32(100), int64(100))
	assertEquivalent(t, int32(100), int(100))
	assertEquivalent(t, int32(100), uint8(100))
	assertEquivalent(t, int32(100), uint16(100))
	assertEquivalent(t, int32(100), uint32(100))
	assertEquivalent(t, int32(100), uint64(100))
	assertEquivalent(t, int32(100), uint(100))
	assertEquivalent(t, int32(100), float32(100))
	assertEquivalent(t, int32(100), float64(100))
}

func TestInt64Equal(t *testing.T) {
	assertEquivalent(t, int64(100), int8(100))
	assertEquivalent(t, int64(100), int16(100))
	assertEquivalent(t, int64(100), int32(100))
	assertEquivalent(t, int64(100), int64(100))
	assertEquivalent(t, int64(100), int(100))
	assertEquivalent(t, int64(100), uint8(100))
	assertEquivalent(t, int64(100), uint16(100))
	assertEquivalent(t, int64(100), uint32(100))
	assertEquivalent(t, int64(100), uint64(100))
	assertEquivalent(t, int64(100), uint(100))
	assertEquivalent(t, int64(100), float32(100))
	assertEquivalent(t, int64(100), float64(100))
}

func TestIntEqual(t *testing.T) {
	assertEquivalent(t, int(100), int8(100))
	assertEquivalent(t, int(100), int16(100))
	assertEquivalent(t, int(100), int32(100))
	assertEquivalent(t, int(100), int64(100))
	assertEquivalent(t, int(100), int(100))
	assertEquivalent(t, int(100), uint8(100))
	assertEquivalent(t, int(100), uint16(100))
	assertEquivalent(t, int(100), uint32(100))
	assertEquivalent(t, int(100), uint64(100))
	assertEquivalent(t, int(100), uint(100))
	assertEquivalent(t, int(100), float32(100))
	assertEquivalent(t, int(100), float64(100))
}

func TestUint8Equal(t *testing.T) {
	assertEquivalent(t, uint8(100), int8(100))
	assertEquivalent(t, uint8(100), int16(100))
	assertEquivalent(t, uint8(100), int32(100))
	assertEquivalent(t, uint8(100), int64(100))
	assertEquivalent(t, uint8(100), int(100))
	assertEquivalent(t, uint8(100), uint8(100))
	assertEquivalent(t, uint8(100), uint16(100))
	assertEquivalent(t, uint8(100), uint32(100))
	assertEquivalent(t, uint8(100), uint64(100))
	assertEquivalent(t, uint8(100), uint(100))
	assertEquivalent(t, uint8(100), float32(100))
	assertEquivalent(t, uint8(100), float64(100))
}

func TestUint16Equal(t *testing.T) {
	assertEquivalent(t, uint16(100), int8(100))
	assertEquivalent(t, uint16(100), int16(100))
	assertEquivalent(t, uint16(100), int32(100))
	assertEquivalent(t, uint16(100), int64(100))
	assertEquivalent(t, uint16(100), int(100))
	assertEquivalent(t, uint16(100), uint8(100))
	assertEquivalent(t, uint16(100), uint16(100))
	assertEquivalent(t, uint16(100), uint32(100))
	assertEquivalent(t, uint16(100), uint64(100))
	assertEquivalent(t, uint16(100), uint(100))
	assertEquivalent(t, uint16(100), float32(100))
	assertEquivalent(t, uint16(100), float64(100))
}

func TestUint32Equal(t *testing.T) {
	assertEquivalent(t, uint32(100), int8(100))
	assertEquivalent(t, uint32(100), int16(100))
	assertEquivalent(t, uint32(100), int32(100))
	assertEquivalent(t, uint32(100), int64(100))
	assertEquivalent(t, uint32(100), int(100))
	assertEquivalent(t, uint32(100), uint8(100))
	assertEquivalent(t, uint32(100), uint16(100))
	assertEquivalent(t, uint32(100), uint32(100))
	assertEquivalent(t, uint32(100), uint64(100))
	assertEquivalent(t, uint32(100), uint(100))
	assertEquivalent(t, uint32(100), float32(100))
	assertEquivalent(t, uint32(100), float64(100))
}

func TestUint64Equal(t *testing.T) {
	assertEquivalent(t, uint64(100), int8(100))
	assertEquivalent(t, uint64(100), int16(100))
	assertEquivalent(t, uint64(100), int32(100))
	assertEquivalent(t, uint64(100), int64(100))
	assertEquivalent(t, uint64(100), int(100))
	assertEquivalent(t, uint64(100), uint8(100))
	assertEquivalent(t, uint64(100), uint16(100))
	assertEquivalent(t, uint64(100), uint32(100))
	assertEquivalent(t, uint64(100), uint64(100))
	assertEquivalent(t, uint64(100), uint(100))
	assertEquivalent(t, uint64(100), float32(100))
	assertEquivalent(t, uint64(100), float64(100))
}

func TestUintEqual(t *testing.T) {
	assertEquivalent(t, uint(100), int8(100))
	assertEquivalent(t, uint(100), int16(100))
	assertEquivalent(t, uint(100), int32(100))
	assertEquivalent(t, uint(100), int64(100))
	assertEquivalent(t, uint(100), int(100))
	assertEquivalent(t, uint(100), uint8(100))
	assertEquivalent(t, uint(100), uint16(100))
	assertEquivalent(t, uint(100), uint32(100))
	assertEquivalent(t, uint(100), uint64(100))
	assertEquivalent(t, uint(100), uint(100))
	assertEquivalent(t, uint(100), float32(100))
	assertEquivalent(t, uint(100), float64(100))
}

func TestFloat32Equal(t *testing.T) {
	assertEquivalent(t, float32(100), int8(100))
	assertEquivalent(t, float32(100), int16(100))
	assertEquivalent(t, float32(100), int32(100))
	assertEquivalent(t, float32(100), int64(100))
	assertEquivalent(t, float32(100), int(100))
	assertEquivalent(t, float32(100), uint8(100))
	assertEquivalent(t, float32(100), uint16(100))
	assertEquivalent(t, float32(100), uint32(100))
	assertEquivalent(t, float32(100), uint64(100))
	assertEquivalent(t, float32(100), uint(100))
	assertEquivalent(t, float32(100), float32(100))
	assertEquivalent(t, float32(100), float64(100))
}

func TestFloat64Equal(t *testing.T) {
	assertEquivalent(t, float64(100), int8(100))
	assertEquivalent(t, float64(100), int16(100))
	assertEquivalent(t, float64(100), int32(100))
	assertEquivalent(t, float64(100), int64(100))
	assertEquivalent(t, float64(100), int(100))
	assertEquivalent(t, float64(100), uint8(100))
	assertEquivalent(t, float64(100), uint16(100))
	assertEquivalent(t, float64(100), uint32(100))
	assertEquivalent(t, float64(100), uint64(100))
	assertEquivalent(t, float64(100), uint(100))
	assertEquivalent(t, float64(100), float32(100))
	assertEquivalent(t, float64(100), float64(100))
}

func TestBigIntEqual(t *testing.T) {
	assertEquivalent(t, new(big.Int), big.NewInt(0))

	assertEquivalent(t, big.NewInt(5), 5)
	assertEquivalent(t, big.NewInt(-70), -70)
	assertEquivalent(t, big.NewInt(100), 1.0e2)
	assertEquivalent(t, big.NewInt(10000000), big.NewFloat(10000000))
	assertEquivalent(t, big.NewInt(10000000), big.NewInt(10000000))

	assertEquivalent(t, 5, big.NewInt(5))
	assertEquivalent(t, -70, big.NewInt(-70))
	assertEquivalent(t, 1.0e2, big.NewInt(100))
	assertEquivalent(t, big.NewFloat(10000000), big.NewInt(10000000))
	assertEquivalent(t, big.NewInt(10000000), big.NewInt(10000000))
}

func TestBigFloatEqual(t *testing.T) {
	assertEquivalent(t, new(big.Float), big.NewFloat(0))

	assertEquivalent(t, big.NewFloat(5), 5)
	assertEquivalent(t, big.NewFloat(-70), -70)
	assertEquivalent(t, big.NewFloat(1.454e100), 1.454e100)
	assertEquivalent(t, big.NewFloat(10000000), big.NewInt(10000000))
	assertEquivalent(t, big.NewFloat(10000000.1234), big.NewFloat(10000000.1234))

	assertEquivalent(t, 5, big.NewFloat(5))
	assertEquivalent(t, -70, big.NewFloat(-70))
	assertEquivalent(t, 1.454e100, big.NewFloat(1.454e100))
	assertEquivalent(t, big.NewInt(10000000), big.NewFloat(10000000))
	assertEquivalent(t, big.NewFloat(10000000.1234), big.NewFloat(10000000.1234))
}

func TestNotEqual(t *testing.T) {
	assertNotEquivalent(t, -1, uint(1))
	assertNotEquivalent(t, uint(1), -1)
	assertNotEquivalent(t, 1.1, 1)
	assertNotEquivalent(t, 1, 1.1)
}

func TestRecursive(t *testing.T) {
	a := make([]interface{}, 1)
	a[0] = a
	b := make([]interface{}, 1)
	b[0] = b

	assertEquivalent(t, a, b)
}

func TestComplex(t *testing.T) {
	a := ComplexStruct{
		Map:     map[interface{}]interface{}{1: "a"},
		Struct:  MyStruct{1, "a"},
		StructP: &MyStruct{100, "test"},
	}

	b := ComplexStruct{
		Map:     map[interface{}]interface{}{int8(1): "a"},
		Struct:  MyStruct{1, "a"},
		StructP: &MyStruct{100, "test"},
	}
	assertEquivalent(t, &a, b)
}

func TestComplexMaps(t *testing.T) {
	a := map[interface{}]interface{}{
		"complex": ComplexStruct{
			Map:     map[interface{}]interface{}{1: "a"},
			Struct:  MyStruct{1, "a"},
			StructP: &MyStruct{100, "test"},
		},
		float32(500): "aaa",
		"x": map[interface{}]interface{}{
			"mystruct": MyStruct{10, "x"},
		},
	}

	b := map[interface{}]interface{}{
		"complex": ComplexStruct{
			Map:     map[interface{}]interface{}{int8(1): "a"},
			Struct:  MyStruct{1, "a"},
			StructP: &MyStruct{100, "test"},
		},
		float32(500): "aaa",
		"x": map[interface{}]interface{}{
			"mystruct": MyStruct{10, "x"},
		},
	}
	assertEquivalent(t, a, &b)
}

func TestManyComparisons(t *testing.T) {
	for i := 0; i < 100000; i++ {
		TestComplexMaps(t)
	}
}

func DemonstrateEquivalence() {
	a := ComplexStruct{
		Map:     map[interface{}]interface{}{1: "a"},
		Struct:  MyStruct{1, "a"},
		StructP: &MyStruct{100, "test"},
	}

	b := &ComplexStruct{
		Map:     map[interface{}]interface{}{int8(1): "a"},
		Struct:  MyStruct{1, "a"},
		StructP: &MyStruct{100, "test"},
	}
	fmt.Printf("Equivalent: %v\n", IsEquivalent(a, b))
}

func TestDemonstrateEquivalence(t *testing.T) {
	DemonstrateEquivalence()
}
