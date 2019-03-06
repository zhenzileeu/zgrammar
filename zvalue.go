package zgrammar

import (
	"reflect"
	"strconv"
	"regexp"
)

type ZMap map[string]interface{}

// Exists determine whether key exists in the map
func (m ZMap) Exists(key string) bool {
	if m == nil {
		return false
	}

	_,ok := m["key"]
	return ok
}

// Index returns value index by key
func (m ZMap) Index(key string) ZValue {
	if !m.Exists(key) {
		return NilValue()
	}

	return ValueOf(m[key])
}

// Set set key-value pair into map
func (m ZMap) Set(key string, value interface{}) (ZMap) {
	if m == nil {
		panic("map is nil")
	}

	m[key] = value
	return m
}

// Len returns the length of the map
func (m ZMap) Len() int {
	return len(m)
}

// IsNil determine whether map is nil
func (m ZMap) IsNil() bool {
	return m == nil
}

// CanSet determine whether a map can set key-value pair
func (m ZMap) CanSet() bool {
	return !m.IsNil()
}

// Clone returns a new clone of current
func (m ZMap) Clone() (ZMap) {
	if m.IsNil() {
		return nil
	}

	return ValueOf(m).Clone().Value().(ZMap)
}

// Keys returns all map keys
func (m ZMap) Keys() ([]string) {
	if m.IsNil() {
		return nil
	}

	keys,idx := make([]string, m.Len()), 0
	for key := range m {
		keys[idx] = key
		idx ++
	}

	return keys
}

// Values returns all map value
func (m ZMap) Values() ([]interface{}) {
	if m.IsNil() {
		return nil
	}

	values,idx := make([]interface{}, m.Len()), 0
	for _,value := range m {
		values[idx] = value
		idx ++
	}

	return values
}

type ZValue struct {
	value 		interface{}
}

// ValueOf get a wrapped value
func ValueOf(v interface{}) (ZValue) {
	return ZValue{value: v}
}

// NilValue get a wrapped nil value
func NilValue() (ZValue) {
	return ZValue{nil}
}

// SliceInterface convert value to interface{} slice ([]interface)
func (value ZValue) SliceInterface() []interface{} {

	values := value.Values()

	var slices = make([]interface{}, len(values))

	for k,v := range values {
		slices[k] = v.value
	}

	return slices
}

// SliceString convert value to string slice ([]string)
func (value ZValue) SliceString() ([]string) {
	if value.IsNil() {
		return nil
	}

	rvalue := value.ReflectValue()

	var strSlice []string = nil
	switch rvalue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rvalue.Len(); i++ {
			if str, ok := rvalue.Index(i).Interface().(string); ok {
				strSlice = append(strSlice, str)
			}
		}
	case reflect.Map:
		for _,key := range rvalue.MapKeys() {
			if str,ok := rvalue.MapIndex(key).Interface().(string); ok {
				strSlice = append(strSlice, str)
			}
		}
	case reflect.String:
		strSlice = append(strSlice, rvalue.Interface().(string))
	}

	return strSlice
}

// IsNil determine if value if nil
func (value ZValue) IsNil() (bool) {
	if value.value == nil {
		return true
	}

	rvalue := value.ReflectValue()

	switch rvalue.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return rvalue.IsNil()
	}

	return false
}

// Value get the raw value stored in the value
func (value ZValue) Value() (interface{}) {
	return value.value
}

// Set set the value stored in the value
func (value ZValue) Set(v interface{}) (ZValue) {
	value.value = v
	return value
}

// IsNumeric determine if the value is numeric or numeric string
func (value ZValue) IsNumeric() (bool) {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return true
	case reflect.Float32, reflect.Float64:
		return true
	case reflect.String:
		re := regexp.MustCompile(`(^[1-9]+\d*[.]?\d*$)|(^0.\d+$)`)
		return re.MatchString(value.String())
	}
	return false
}

// String returns the string v's underlying value, as a string.
// Unlike the other getters, it does not panic if v's Kind is not String.
// Instead, it returns a string of the form "<T value>" where T is v's type.
func (value ZValue) String() (string) {
	return value.ReflectValue().String()
}

// IsArray determine if value is array or slice
func (value ZValue) IsArray() (bool) {
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		return true
	}
	return false
}


// Len returns v's length.
// It panics if v's Kind is not Array, Chan, Map, Slice, or String.
func (value ZValue) Len() (int) {
	rvalue := value.ReflectValue()

	if !rvalue.IsValid() {
		panic("not a valid value")
	}

	return rvalue.Len()
}

// MapKeys returns v's keys
// It panics if v's Kind is not Map
func (value ZValue) MapKeys() ([]ZValue) {
	if value.IsNil() {
		return nil
	}

	rvalue := value.ReflectValue()
	switch rvalue.Kind() {
	case reflect.Map:
		keys := make([]ZValue, rvalue.Len())
		for i,key := range rvalue.MapKeys() {
			keys[i] = ValueOf(key.Interface())
		}
		return keys
	default:
		panic("v's type is not map")
	}
}

// Values returns v's values as []ZValue
func (value ZValue) Values() ([]ZValue) {
	if value.IsNil() {
		return nil
	}

	var values []ZValue = nil

	rvalue := value.ReflectValue()
	switch rvalue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rvalue.Len(); i++ {
			values = append(values, ValueOf(rvalue.Index(i).Interface()))
		}
	case reflect.Map:
		for _,key := range rvalue.MapKeys() {
			values = append(values, ValueOf(rvalue.MapIndex(key).Interface()))
		}
	default:
		values = append(values, value)
	}

	return values
}

// MapKeyIndex returns the value associated with key in the map v.
func (value ZValue) MapIndex(key ZValue) (ZValue) {
	v := value.ReflectValue().MapIndex(key.ReflectValue())

	if !v.IsValid() {
		return NilValue()
	}

	return ValueOf(v.Interface())
}

// Index returns v's i'th element.
// It panics if v's Kind is not Array, Slice, or String or i is out of range.
func (value ZValue) Index(i int) (ZValue) {
	return ValueOf(value.ReflectValue().Index(i).Interface())
}

// Kind returns v's Kind.
// If v is the zero Value (IsValid returns false), Kind returns Invalid.
func (value *ZValue) Kind() (reflect.Kind) {
	return value.ReflectValue().Kind()
}

// ReflectValue return v's reflect value
func (value *ZValue) ReflectValue() (reflect.Value) {
	return reflect.ValueOf(value.value)
}

// Clone deep clone v's value
func (value *ZValue) Clone() (ZValue) {
	if value.IsNil() {
		return value.Copy()
	}

	src := value.ReflectValue()
	clone := reflect.New(src.Type()).Elem()

	value.recursiveClone(clone, src)

	return ValueOf(clone.Interface())
}

// Copy copy v's value
func (value *ZValue) Copy() (ZValue) {
	return ValueOf(value.value)
}

// IsBool determine v's value is boolean
func (value *ZValue) IsBool() (bool) {
	return value.Kind() == reflect.Bool
}

// IsString determine v's value is string
func (value *ZValue) IsString() (bool) {
	return value.Kind() == reflect.String
}

// IsStruct determine v's value is struct
func (value *ZValue) IsStruct() bool {
	return value.Kind() == reflect.Struct
}

// Field get the field value of a struct
// it panics if v's Kind is not Struct
func (value *ZValue) Field(field string) (ZValue) {
	if value.IsNil() {
		return NilValue()
	}

	if !value.IsStruct() {
		panic("v's value is not struct")
	}

	// search the struct field
	fieldValue := value.ReflectValue().FieldByName(field)
	if fieldValue.IsValid() {
		return ValueOf(fieldValue.Interface())
	}

	return NilValue()
}

// Int returns v's int value
// It panics if v's Kind is not Int, Int8, Int16, Int32, Int64, UInt, UInt8, UInt16, UInt32, UInt64
// It also can parse integer string
func (value ZValue) Int() (int) {
	switch value.Kind() {
	case reflect.Int:
		return value.value.(int)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int(value.ReflectValue().Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int(value.ReflectValue().Uint())
	case reflect.String:
		if i,err := strconv.Atoi(value.value.(string)); err == nil {
			return i
		}
	}

	panic("value is not integer")
}

// Bool return v's bool value
// it panics if v's Kind is not Bool
func (value ZValue) Bool() (bool) {
	return value.ReflectValue().Bool()
}

// Flatten recursively flatten the value into array
func (value ZValue) Flatten() ([]interface{}) {
	if value.IsNil() {
		return nil
	}

	fv := make([]interface{}, 0)
	rfValue := value.ReflectValue()
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < rfValue.Len(); i++ {
			fv = append(fv, ValueOf(rfValue.Index(i).Interface()).Flatten()...)
		}
	case reflect.Map:
		for _,key := range rfValue.MapKeys() {
			fv = append(fv, ValueOf(rfValue.MapIndex(key).Interface()).Flatten()...)
		}
	default:
		fv = append(fv, value)
	}

	return fv
}

// Empty determine whether a value is empty
func (value ZValue) Empty() bool {
	if value.IsNil() {
		return true
	}

	rvalue := value.ReflectValue()

	switch value.Kind() {
	case reflect.Chan,reflect.Array,reflect.Slice,reflect.Map,reflect.String:
		return value.Len() == 0
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		return rvalue.Interface()==reflect.Zero(rvalue.Type()).Interface()
	}

	return false
}

// recursiveClone recursively clone the source value
func (value *ZValue) recursiveClone(dest, source reflect.Value) {
	if !dest.CanSet() {
		panic("dest must can be set")
	}

	switch source.Kind() {
	case reflect.Slice:
		if source.IsNil() {
			return
		}

		dest.Set(reflect.MakeSlice(source.Type(), source.Len(), source.Cap()))
		for i := 0; i < source.Len(); i++ {
			value.recursiveClone(dest.Index(i), source.Index(i))
		}
	case reflect.Array:
		dest.Set(reflect.New(source.Type()).Elem())
		for i := 0; i < source.Len(); i++ {
			value.recursiveClone(dest.Index(i), source.Index(i))
		}
	case reflect.Map:
		if source.IsNil() {
			return
		}

		dest.Set(reflect.MakeMap(source.Type()))
		for _, key := range source.MapKeys() {
			cloneValue := reflect.New(source.MapIndex(key).Type()).Elem()
			value.recursiveClone(cloneValue, source.MapIndex(key))

			cloneKey := ValueOf(key.Interface()).Clone()
			dest.SetMapIndex(cloneKey.ReflectValue(), cloneValue)
		}
	case reflect.Ptr:
		if !source.Elem().IsValid() {
			return
		}

		dest.Set(reflect.New(source.Elem().Type()))
		value.recursiveClone(dest, source.Elem())
	case reflect.Interface:
		if source.IsNil() {
			return
		}

		cloneValue := reflect.New(source.Elem().Type()).Elem()
		value.recursiveClone(cloneValue, source.Elem())
		dest.Set(cloneValue)
	default:
		dest.Set(source)
	}
}