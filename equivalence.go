// Equivalence tests whether two objects are considered equivalent, which means
// that an object (and its constituent parts) is equal to another other object
// (and its constituent parts), disregarding actual data types. If one object
// can be converted to the other without data loss and still be equal, it is
// considered equivalent.
//
// Note: This comparison isn't cheap; it's primarily meant for unit test code.
package equivalence

import (
	"math"
	"reflect"
)

var floatToleranceFactor = 0.0000000001

func asInt(value reflect.Value) (result int64, ok bool) {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int(), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue := value.Uint()
		if uintValue <= math.MaxInt64 {
			return int64(uintValue), true
		}
	case reflect.Float32, reflect.Float64:
		floatValue := value.Float()
		intValue := int64(floatValue)
		if float64(intValue) == floatValue {
			return intValue, true
		}
	}
	return 0, false
}

func asUint(value reflect.Value) (result uint64, ok bool) {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := value.Int()
		if intValue >= 0 {
			return uint64(intValue), true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint(), true
	case reflect.Float32, reflect.Float64:
		floatValue := value.Float()
		uintValue := uint64(floatValue)
		if float64(uintValue) == floatValue {
			return uintValue, true
		}
	}
	return 0, false
}

func asFloat(value reflect.Value) (result float64, ok bool) {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intValue := value.Int()
		floatValue := float64(intValue)
		if int64(floatValue) == intValue {
			return floatValue, true
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		uintValue := value.Uint()
		floatValue := float64(uintValue)
		if uint64(floatValue) == uintValue {
			return floatValue, true
		}
	case reflect.Float32, reflect.Float64:
		return value.Float(), true
	}
	return 0, false
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

func areArraysOrSlicesEquivalent(a, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if !areObjectsEquivalent(a.Index(i), b.Index(i)) {
			return false
		}
	}
	return true
}

func areMapsEquivalent(a, b reflect.Value) bool {
	if a.Len() != b.Len() {
		return false
	}
	iter := a.MapRange()
	for iter.Next() {
		k := iter.Key()
		av := iter.Value()
		bv := getMapValue(b, k)
		if !areObjectsEquivalent(av, bv) {
			return false
		}
	}
	return true
}

func areStructsEquivalent(a, b reflect.Value) bool {
	if a.NumField() != b.NumField() {
		return false
	}
	for i := 0; i < a.NumField(); i++ {
		if !areObjectsEquivalent(a.Field(i), b.Field(i)) {
			return false
		}
	}
	return true
}

func areObjectsEquivalent(a, b reflect.Value) bool {
	if !a.IsValid() && !b.IsValid() {
		// Special case: untyped nil
		return true
	}

	switch a.Kind() {
	case reflect.Interface, reflect.Ptr:
		return areObjectsEquivalent(a.Elem(), b.Elem())
	case reflect.Bool:
		return a.Bool() == b.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		bValue, ok := asInt(b)
		return ok && a.Int() == bValue
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		bValue, ok := asUint(b)
		return ok && a.Uint() == bValue
	case reflect.Float32, reflect.Float64:
		bValue, ok := asFloat(b)
		aValue := a.Float()
		if !ok {
			return false
		}
		if math.IsNaN(aValue) && math.IsNaN(bValue) {
			return true
		}
		if aValue == bValue {
			return true
		}
		tolerance := math.Abs(aValue * floatToleranceFactor)
		return math.Abs(aValue-bValue) <= tolerance
	case reflect.Complex64, reflect.Complex128:
		return a.Complex() == b.Complex()
	case reflect.String:
		return a.Type() == b.Type() && a.String() == b.String()
	case reflect.Array, reflect.Slice:
		return areArraysOrSlicesEquivalent(a, b)
	case reflect.Map:
		return areMapsEquivalent(a, b)
	case reflect.Struct:
		return areStructsEquivalent(a, b)
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

func drillDown(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	return v
}

// Test if two objects are equivalent.
//
// Equivalence means that they are either equal, or one can be converted to the
// other's type without data loss and still be equal.
//
// The initial compared objects will be drilled down through pointers and
// interfaces, and the first concrete value of each object will be used for the
// comparison.
//
// For slices, arrays, maps, and structs, it will compare elements. Element
// values will not be drilled down.
//
// NaN values are considered equal.
// Empty containers are considered equal, regardless of element type.
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
	return areObjectsEquivalent(drillDown(reflect.ValueOf(a)), drillDown(reflect.ValueOf(b)))
}
