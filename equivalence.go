// Equivalence tests whether two objects are considered equivalent, which means
// that an object (and its constituent parts) is equal to another other object
// (and its constituent parts), disregarding actual data types. If one object
// can be converted to the other without data loss and still be equal, it is
// considered equivalent.
//
// Note: This comparison isn't cheap; it's primarily meant for unit test code.
package equivalence

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"

	"github.com/kstenerud/go-duplicates"
)

// Test if two objects are equivalent.
//
// Equivalence means that they are either equal, or one can be converted to the
// other's type without data loss and still be considered equal.
//
// The initial compared objects will be drilled down through pointers and
// interfaces, and the first concrete value of each object will be used for the
// comparison.
//
// The following numeric types will be converted (if an exact conversion is
// possible) and numerically compared: int, uint, float, big.Int, big.Float
//
// For slices, arrays, maps, and structs, it will compare elements. Element
// values will not be drilled down.
//
// NaN values are considered equivalent, regardless of actual payload.
// Empty containers are considered equivalent, regardless of element type.
func IsEquivalent(a, b interface{}) (isEquivalent bool) {
	defer func() {
		// The internal comparison functions just assume that the types are compatible,
		// which causes panics when that's not actually the case. It's simpler
		// to just let them panic since this also means that they cannot be equivalent.
		if r := recover(); r != nil {
			isEquivalent = false
		}
	}()
	if a == nil && b == nil {
		return true
	}
	c := newComparator()
	return c.areObjectsEquivalent(reflect.ValueOf(a), reflect.ValueOf(b))
}

type comparator struct {
	aFinder duplicates.DuplicateFinder
	bFinder duplicates.DuplicateFinder
}

func newComparator() *comparator {
	_this := &comparator{}
	_this.Init()
	return _this
}

func (_this *comparator) Init() {
	_this.aFinder.Init()
	_this.bFinder.Init()
}

func getIntKeyedMapValue(aMap reflect.Value, aKey int64) reflect.Value {
	initialResult := aMap.MapIndex(reflect.ValueOf(aKey))
	if initialResult.IsValid() {
		return initialResult
	}
	asInt := int(aKey)
	if int64(asInt) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asInt)); v.IsValid() {
			return v
		}
	}
	asInt32 := int32(aKey)
	if int64(asInt32) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asInt32)); v.IsValid() {
			return v
		}
	}
	asInt16 := int16(aKey)
	if int64(asInt16) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asInt16)); v.IsValid() {
			return v
		}
	}
	asInt8 := int8(aKey)
	if int64(asInt8) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asInt8)); v.IsValid() {
			return v
		}
	}
	return initialResult
}

func getUintKeyedMapValue(aMap reflect.Value, aKey uint64) reflect.Value {
	initialResult := aMap.MapIndex(reflect.ValueOf(aKey))
	if initialResult.IsValid() {
		return initialResult
	}
	asUint := uint(aKey)
	if uint64(asUint) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asUint)); v.IsValid() {
			return v
		}
	}
	asUint32 := uint32(aKey)
	if uint64(asUint32) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asUint32)); v.IsValid() {
			return v
		}
	}
	asUint16 := uint16(aKey)
	if uint64(asUint16) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asUint16)); v.IsValid() {
			return v
		}
	}
	asUint8 := uint8(aKey)
	if uint64(asUint8) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asUint8)); v.IsValid() {
			return v
		}
	}
	return initialResult
}

func getFloatKeyedMapValue(aMap reflect.Value, aKey float64) reflect.Value {
	initialResult := aMap.MapIndex(reflect.ValueOf(aKey))
	if initialResult.IsValid() {
		return initialResult
	}
	asFloat32 := float32(aKey)
	if float64(asFloat32) == aKey {
		if v := aMap.MapIndex(reflect.ValueOf(asFloat32)); v.IsValid() {
			return v
		}
	}
	return initialResult
}

func getMapValue(aMap reflect.Value, aKey reflect.Value) reflect.Value {
	if aKey.Kind() == reflect.Interface {
		aKey = aKey.Elem()
	}

	initialResult := aMap.MapIndex(aKey)
	if initialResult.IsValid() {
		return initialResult
	}

	switch aKey.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		asInt := aKey.Int()
		if v := getIntKeyedMapValue(aMap, asInt); v.IsValid() {
			return v
		}

		if asInt >= 0 {
			if v := getUintKeyedMapValue(aMap, uint64(asInt)); v.IsValid() {
				return v
			}
		}
		return getFloatKeyedMapValue(aMap, float64(asInt))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		asUint := aKey.Uint()
		if v := getUintKeyedMapValue(aMap, asUint); v.IsValid() {
			return v
		}
		if asUint <= math.MaxInt64 {
			if v := getIntKeyedMapValue(aMap, int64(asUint)); v.IsValid() {
				return v
			}
		}
		return getFloatKeyedMapValue(aMap, float64(asUint))
	case reflect.Float32, reflect.Float64:
		asFloat := aKey.Float()
		if asInt := int64(asFloat); float64(asInt) == asFloat {
			if v := getIntKeyedMapValue(aMap, asInt); v.IsValid() {
				return v
			}
		}
		if asUint := uint64(asFloat); float64(asUint) == asFloat {
			if v := getUintKeyedMapValue(aMap, asUint); v.IsValid() {
				return v
			}
		}
	default:
	}
	return initialResult
}

func (_this *comparator) areArraysOrSlicesEquivalent(a, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if !_this.areObjectsEquivalent(a.Index(i), b.Index(i)) {
			return false
		}
	}
	return true
}

func (_this *comparator) areMapsEquivalent(a, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}
	iter := mapRange(a)
	for iter.Next() {
		k := iter.Key()
		av := iter.Value()
		bv := getMapValue(b, k)
		if !_this.areObjectsEquivalent(av, bv) {
			return false
		}
	}
	return true
}

var bigIntType = reflect.TypeOf(big.Int{})
var bigFloatType = reflect.TypeOf(big.Float{})
var bitsToDecimalDigitsTable = []int{0, 1, 1, 1, 1, 2, 2, 2, 3, 3}

func bitsToDecimalDigits(bitCount int) int {
	return (bitCount/10)*3 + bitsToDecimalDigitsTable[bitCount%10]
}

func bigFloatToString(v big.Float) string {
	return v.Text('g', bitsToDecimalDigits(int(v.Prec())))
}

func floatToString(v float64) string {
	return strconv.FormatFloat(v, 'g', -1, 64)
}

func areBigFloatsEquivalent(a, b reflect.Value) bool {
	return bigFloatToString(a.Interface().(big.Float)) == bigFloatToString(b.Interface().(big.Float))
}

func (_this *comparator) areStructsEquivalent(a, b reflect.Value) bool {
	switch a.Type() {
	case bigIntType:
		return isEquivalentToBigInt(a.Interface().(big.Int), b)
	case bigFloatType:
		return isEquivalentToBigFloat(a.Interface().(big.Float), b)
	}

	if a.NumField() != b.NumField() {
		return false
	}
	for i := 0; i < a.NumField(); i++ {
		if !_this.areObjectsEquivalent(a.Field(i), b.Field(i)) {
			return false
		}
	}
	return true
}

func numericToString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return floatToString(v.Float())
	case reflect.Struct:
		switch v.Type() {
		case bigIntType:
			val := v.Interface().(big.Int)
			return val.String()
		case bigFloatType:
			val := v.Interface().(big.Float)
			return bigFloatToString(val)
		}
	}
	return fmt.Sprintf("NOT NUMERIC: %v", v)
}

func isEquivalentToBigFloat(a big.Float, b reflect.Value) bool {
	return bigFloatToString(a) == numericToString(b)
}

func isEquivalentToBigInt(a big.Int, b reflect.Value) bool {
	return a.String() == numericToString(b)
}

func isEquivalentToInt(a int64, b reflect.Value) bool {
	switch b.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return b.Int() == a
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ub := b.Uint()
		if ub&0x1000000000000000 != 0 {
			return false
		}
		return a == int64(ub)
	case reflect.Float32, reflect.Float64:
		fb := b.Float()
		return a == int64(fb) && float64(a) == fb
	case reflect.Struct:
		switch b.Type() {
		case bigIntType:
			bi := b.Interface().(big.Int)
			return strconv.FormatInt(a, 10) == bi.String()
		case bigFloatType:
			return strconv.FormatInt(a, 10) == bigFloatToString(b.Interface().(big.Float))
		}
		return false
	default:
		return false
	}
}

func isEquivalentToUint(a uint64, b reflect.Value) bool {
	switch b.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ib := b.Int()
		if ib < 0 {
			return false
		}
		return a == uint64(ib)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return a == b.Uint()
	case reflect.Float32, reflect.Float64:
		fb := b.Float()
		if fb < 0 {
			return false
		}
		return a == uint64(fb) && float64(a) == fb
	case reflect.Struct:
		switch b.Type() {
		case bigIntType:
			bi := b.Interface().(big.Int)
			return strconv.FormatUint(a, 10) == bi.String()
		case bigFloatType:
			return strconv.FormatUint(a, 10) == bigFloatToString(b.Interface().(big.Float))
		}
		return false
	default:
		return false
	}
}

func isEquivalentToFloat(a float64, b reflect.Value) bool {
	switch b.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ib := b.Int()
		return a == float64(ib) && int64(a) == ib
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		ub := b.Uint()
		return a == float64(ub) && uint64(a) == ub
	case reflect.Float32, reflect.Float64:
		fb := b.Float()
		if math.IsNaN(a) && math.IsNaN(fb) {
			return true
		}
		return a == fb
	case reflect.Struct:
		switch b.Type() {
		case bigIntType:
			bi := b.Interface().(big.Int)
			return floatToString(a) == bi.String()
		case bigFloatType:
			return floatToString(a) == bigFloatToString(b.Interface().(big.Float))
		}
		return false
	default:
		return false
	}
}

func (_this *comparator) areObjectsEquivalent(a, b reflect.Value) bool {
	var aHasDuplicate, bHasDuplicate bool
	a, aHasDuplicate = drillDown(&_this.aFinder, a)
	b, bHasDuplicate = drillDown(&_this.bFinder, b)

	if aHasDuplicate || bHasDuplicate {
		return true
	}

	if !a.IsValid() || !b.IsValid() {
		// Special case: zero value
		return !a.IsValid() && !b.IsValid()
	}

	switch a.Kind() {
	case reflect.Bool:
		return a.Bool() == b.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return isEquivalentToInt(a.Int(), b)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return isEquivalentToUint(a.Uint(), b)
	case reflect.Float32, reflect.Float64:
		return isEquivalentToFloat(a.Float(), b)
	case reflect.Complex64, reflect.Complex128:
		return a.Complex() == b.Complex()
	case reflect.String:
		return a.Type() == b.Type() && a.String() == b.String()
	case reflect.Array:
		return _this.areArraysOrSlicesEquivalent(a, b)
	case reflect.Slice:
		if hasDuplicate := _this.aFinder.RegisterPointer(a); hasDuplicate {
			return true
		}
		return _this.areArraysOrSlicesEquivalent(a, b)
	case reflect.Map:
		if hasDuplicate := _this.aFinder.RegisterPointer(a); hasDuplicate {
			return true
		}
		return _this.areMapsEquivalent(a, b)
	case reflect.Struct:
		return _this.areStructsEquivalent(a, b)
	case reflect.Uintptr:
		return a.Pointer() == b.Pointer()
	case reflect.UnsafePointer:
		return a.UnsafeAddr() == b.UnsafeAddr()
	case reflect.Chan:
		return a.Type() == b.Type() && a.Type().Elem() == b.Type().Elem() && a.Type().ChanDir() == b.Type().ChanDir()
	case reflect.Func:
		return a.Type() == b.Type()
	default:
		return false
	}
}

func drillDown(finder *duplicates.DuplicateFinder, v reflect.Value) (value reflect.Value, hasDuplicate bool) {
	for v.IsValid() && (v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface) {
		if v.Kind() == reflect.Ptr {
			if hasDuplicate = finder.RegisterPointer(v); hasDuplicate {
				return v, true
			}
		}
		v = v.Elem()
	}
	return v, false
}
