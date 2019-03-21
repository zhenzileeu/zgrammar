package zgrammar

import (
	"testing"
	"fmt"
	"strings"
	"errors"
)

var testMap ZMap = nil

func TestZMap_Set(t *testing.T) {
	testMap = make(ZMap)
	testMap.Set("k1", 123)
	testMap.Set("k2", "123")

	if v,ok := testMap["k1"]; !ok {
		panic("map key set failure, k1 is not exists")
	} else if vi,ok := v.(int); !ok {
		panic("map key set failure, k1 is not integer")
	} else if vi != 123 {
		panic("map key set failure, k1 is not 123")
	} else {
		fmt.Println("map key set success, k1 is ", vi)
	}

	if v,ok := testMap["k2"]; !ok {
		panic("map key set failure, k2 is not exists")
	} else if vs,ok := v.(string); !ok {
		panic("map key set failure, k2 is not string")
	} else if vs != "123" {
		panic("map key set failure, k2 is not 123")
	} else {
		fmt.Println("map key set success, k2 is ", vs)
	}

	testMap.Set("k1", "34.555")
	if v,ok := testMap["k1"]; !ok {
		panic("map key set failure, k1 is not exists")
	} else if vs,ok := v.(string); !ok {
		panic("map key set failure, k1 is not string")
	} else if vs != "34.555" {
		panic("map key set failure, k1 is not 34.555")
	} else {
		fmt.Println("map key set success, k1 is ", vs)
	}

	fmt.Println(testMap)
	reset()
}

func TestZMap_CanSet(t *testing.T) {
	testMap = nil
	if testMap.CanSet() {
		panic("test map can not set")
	} else {
		fmt.Println("test map is nil")
	}

	testMap = make(ZMap)
	if !testMap.CanSet() {
		panic("test map should can set")
	} else {
		fmt.Println("test map can set")
	}

	reset()
}

func TestZMap_IsNil(t *testing.T) {
	testMap = nil
	if !testMap.IsNil() {
		panic("test map is nil, assertion failure")
	} else {
		fmt.Println("test map is nil, assertion success")
	}

	testMap = make(ZMap)
	if testMap.IsNil() {
		panic("test map is not nil, assertion failure")
	} else {
		fmt.Println("test map is not nil, assertion success")
	}

	reset()
}

func TestZMap_Exists(t *testing.T) {
	testMap = nil
	if testMap.Exists("k1") {
		panic("k1 is not exists assertion failure")
	} else {
		fmt.Println("k1 is not exists, assertion success")
	}

	testMap = make(ZMap)
	testMap.Set("k1", 123)
	if !testMap.Exists("k1") {
		panic("k1 exists, assertion failure")
	} else {
		fmt.Println("k1 exists, value is ", testMap["k1"])
	}

	if testMap.Exists("k2") {
		panic("k2 is not exists, assertion failure")
	} else {
		fmt.Println("k2 is not exists, assertion success")
	}

	reset()
}

func TestZMap_Index(t *testing.T) {
	defer func() {
		if r := recover(); r!=nil {
			if ValueOf(r).String() == "k2 cannot be found in current map" {
				fmt.Println("k2 is not in current map, assertion success")
			} else {
				panic(r)
			}
		}
	}()
	testMap = make(ZMap)
	testMap.Set("k1", "123")
	v := testMap.Index("k1")
	if v.String() == "123" {
		fmt.Println("map k1 values ", v.String())
	} else {
		panic("map k1 not indexed, assertion failure")
	}

	testMap.Index("k2")

	reset()
}

func TestZMap_Len(t *testing.T) {
	testMap = nil
	if testMap.Len() != 0 {
		panic(fmt.Sprintf("test map is nil, length should be 0, assertion is %d", testMap.Len()))
	} else {
		fmt.Println("test map is nil, length is ", testMap.Len())
	}

	testMap = make(ZMap)
	testMap.Set("k1", 100)
	testMap.Set("k2", 123)
	testMap.Set("k3", testMap)
	testMap.Set("k2", "123")
	if testMap.Len() != 3 {
		panic(fmt.Sprintf("test map length should be 3, assertion is %d", testMap.Len()))
	} else {
		fmt.Println("test map length is ", testMap.Len())
	}

	testMap.Set("k4", 12345)
	if testMap.Len() != 4 {
		panic(fmt.Sprintf("test map length should be 4, assertion is %d", testMap.Len()))
	} else {
		fmt.Println("test map length is ", testMap.Len())
	}
	reset()
}

func TestZMap_Keys(t *testing.T) {
	testMap = nil
	keys := testMap.Keys()
	if keys == nil {
		fmt.Println("test map is nil, no key exists")
	} else {
		panic("test map is nil, no key exists, assertion failure")
	}

	testMap = make(ZMap)
	testMap.Set("k1", "ddd")
	testMap.Set("k2", nil)
	keys = testMap.Keys()
	if len(keys) != 2 {
		panic(fmt.Sprintf("test map has two keys, assertion %d", len(keys)))
	} else if keys[0]=="k1" && keys[1]=="k2" {
		fmt.Println("test map has two keys, they are ", keys)
	} else if keys[0]=="k2" && keys[2]=="k1" {
		fmt.Println("test map has two keys, they are ", keys)
	} else {
		panic(fmt.Sprintf("test map keys should be k1 and k2, assertion %v", keys))
	}

	fmt.Println("test map keys order, 2nd call results:  ", testMap.Keys())
	fmt.Println("test map keys order, 3rd call results:  ", testMap.Keys())
	fmt.Println("test map keys order, 4th call results:  ", testMap.Keys())

	testMap.Set("k3", "ddd")
	fmt.Println("test map keys, results: ", testMap.Keys())

	reset()
}

func TestZMap_SortedKeys(t *testing.T) {
	defer func() {
		if r:=recover(); r!=nil {
			if ValueOf(r).String() == "sort direction [xxx] is unknown" {
				fmt.Println("sort direction [xxx] is unknown")
			} else {
				panic(r)
			}
		}
	}()

	testMap = nil
	keys := testMap.SortedKeys("asc")
	if keys == nil {
		fmt.Println("test map is nil, keys is also nil")
	} else {
		panic(fmt.Sprintf("test map is nil, keys should be nil, assertion: %v", keys))
	}

	//
	testMap = make(ZMap)
	testMap.Set("k1", "hee")
	testMap.Set("k2", "nnn")
	testMap.Set("1k", "nnn")
	testMap.Set("2k", "ddd")
	testMap.Set("K1", "dxx")

	keys = testMap.SortedKeys("asc")
	if len(keys) != testMap.Len() {
		panic(fmt.Sprintf("test map length is %d, keys length is %d", testMap.Len(), len(keys)))
	} else if keys[0]=="1k" && keys[1]=="2k" && keys[2]=="K1" && keys[3]=="k1" && keys[4]=="k2" {
		fmt.Println("test map sorted keys is [ASC] : ", keys)
	} else {
		panic(fmt.Sprintf("test map sorted keys is not ASC, assertion: %v", keys))
	}

	keys = testMap.SortedKeys("desc")
	if len(keys) != testMap.Len() {
		panic(fmt.Sprintf("test map length is %d, keys length is %d", testMap.Len(), len(keys)))
	} else if keys[0]=="k2" && keys[1]=="k1" && keys[2]=="K1" && keys[3]=="2k" && keys[4]=="1k" {
		fmt.Println("test map sorted keys is [DESC] : ", keys)
	} else {
		panic(fmt.Sprintf("test map sorted keys is not DESC, assertion: %v", keys))
	}

	keys = testMap.SortedKeys("xxx")

	reset()
}

func TestZMap_Values(t *testing.T) {
	testMap = nil
	values := testMap.Values()
	if values == nil {
		fmt.Println("test map is nil, so values is nil")
	} else {
		panic(fmt.Sprintf("test map is nil, values should be nil, %v returned", values))
	}

	testMap = make(ZMap)
	testMap.Set("k1", "dddd")
	testMap.Set("k2", 1234)
	values = testMap.Values()
	if len(values) != testMap.Len() {
		panic(fmt.Sprintf("test map length is %d, values length is %d", testMap.Len(), len(values)))
	}
	valuesStr := fmt.Sprintf("%v", values)
	if strings.Contains(valuesStr, "dddd") && strings.Contains(valuesStr, "1234") {
		fmt.Println("test map values are: ", values)
	} else {
		panic(fmt.Sprintf("test map values are %v", values))
	}

	// values order is not guaranteed
	fmt.Println("test map values are: ", testMap.Values())
	fmt.Println("test map values are: ", testMap.Values())
	fmt.Println("test map values are: ", testMap.Values())
	fmt.Println("test map values are: ", testMap.Values())

	reset()
}

func TestZMap_StableValues(t *testing.T) {
	testMap = nil
	values := testMap.Values()
	if values == nil {
		fmt.Println("test map is nil, so values is nil")
	} else {
		panic(fmt.Sprintf("test map is nil, values should be nil, %v returned", values))
	}

	testMap = make(ZMap)
	testMap.Set("k1", "dddd")
	testMap.Set("k2", 1234)
	values = testMap.DirectionValues("asc")
	if len(values) != testMap.Len() {
		panic(fmt.Sprintf("test map length is %d, values length is %d", testMap.Len(), len(values)))
	}
	if values[0] == "dddd" && values[1] == 1234 {
		fmt.Println("test map stable values are : ", values)
	} else {
		panic(fmt.Sprintf("test map stable values are not key asc sorted: %v", values))
	}


	values = testMap.DirectionValues("desc")
	if len(values) != testMap.Len() {
		panic(fmt.Sprintf("test map length is %d, values length is %d", testMap.Len(), len(values)))
	}
	if values[0] == 1234 && values[1] == "dddd" {
		fmt.Println("test map stable values are : ", values)
	} else {
		panic(fmt.Sprintf("test map stable values are not key desc sorted: %v", values))
	}

	reset()
}

func TestZMap_Delete(t *testing.T) {
	testMap = nil
	testMap.Delete("k1")
	if testMap.Exists("k1") {
		panic("test map is nil, k1 should not exists")
	} else {
		fmt.Println("test map is nil, k1 is not exists")
	}

	testMap = make(ZMap)
	testMap.Set("k1", 111)
	testMap.Delete("k1")
	if testMap.Len() > 0 {
		panic(fmt.Sprintf("test map should empty, assertion %v", testMap))
	} else {
		fmt.Println("test map is empty")
	}

	testMap.Set("k1", 111)
	testMap.Delete("k2")
	if testMap.Len() != 1 {
		panic(fmt.Sprintf("test map should only have one element, assertion %v", testMap))
	} else {
		fmt.Println("test map is ", testMap)
	}

	reset()
}

func TestZMap_Clone(t *testing.T) {
	testMap = nil
	cloneMap := testMap.Clone()
	if !cloneMap.IsNil() {
		panic(fmt.Sprintf("test map is nil, clone should be nil too, assertion : %v", cloneMap))
	} else {
		fmt.Println("clone map is nil")
	}

	testMap = make(ZMap)
	if !cloneMap.IsNil() {
		panic(fmt.Sprintf("test map is made, but clone should be nil, assertion : %v", cloneMap))
	} else {
		fmt.Println("clone map is nil")
	}

	testMap.Set("k1", "123")
	testMap.Set("k2", []int{1,2,3,4})
	testMap.Set("k3", 1234)
	cloneMap = testMap.Clone()
	if !cloneMap.Exists("k1") || cloneMap["k1"] != "123" {
		panic(fmt.Sprintf("k1 should be the key of clone map, assertion %v", cloneMap))
	} else if !cloneMap.Exists("k2") {
		panic(fmt.Sprintf("k2 should be key of clone map, assertion %v", cloneMap))
	} else if kv,ok := cloneMap["k2"].([]int); !ok {
		panic(fmt.Sprintf("value of k2 should be int slice, assertion %v", cloneMap["k2"]))
	} else if len(kv)!=4 || kv[0]!=1 || kv[1]!=2 || kv[2]!=3 || kv[3]!=4 {
		panic(fmt.Sprintf("value of k2 should be [1,2,3,4], assertion %v", cloneMap["k2"]))
	} else if !cloneMap.Exists("k3") || cloneMap["k3"] != 1234 {
		panic(fmt.Sprintf("k3 should be key of clone map, assertion %v", cloneMap))
	}
	fmt.Println("clone map is ", cloneMap)

	testMap["k2"].([]int)[2] = 56
	if !cloneMap.Exists("k2") {
		panic(fmt.Sprintf("k2 should be key of clone map, assertion %v", cloneMap))
	} else if kv,ok := cloneMap["k2"].([]int); !ok {
		panic(fmt.Sprintf("value of k2 should be int slice, assertion %v", cloneMap["k2"]))
	} else if len(kv)!=4 || kv[0]!=1 || kv[1]!=2 || kv[2]!=3 || kv[3]!=4 {
		panic(fmt.Sprintf("value of k2 should be [1,2,3,4], assertion %v", cloneMap["k2"]))
	}
	fmt.Println("clone map is ",cloneMap, ", and test map is ", testMap)

	cloneMap.Set("k2","1234")
	fmt.Println("clone map is ",cloneMap, ", and test map is ", testMap)

	reset()
}

func TestZValue_Bool(t *testing.T) {
	value := ValueOf(true)
	if value.Bool() {
		fmt.Println("value is true")
	} else {
		panic(fmt.Sprintf("value should be boolean true, assertion %v", value.Bool()))
	}

	value = ValueOf(false)
	if value.Bool() {
		panic(fmt.Sprintf("value should be boolean false, assertion %v", value.Bool()))
	} else {
		fmt.Println("value is false")
	}
}

func TestZValue_Empty(t *testing.T) {
	value := ValueOf(nil)
	if !value.Empty() {
		panic(fmt.Sprintf("value is nil, should be empty"))
	} else {
		fmt.Println("value is nil")
	}

	var x ZMap
	value = ValueOf(x)
	if !value.Empty() {
		panic(fmt.Sprintf("value is nil map, should be empty"))
	} else {
		fmt.Println("value is nil map")
	}

	value = ValueOf([]int{})
	if !value.Empty() {
		panic(fmt.Sprintf("value is empty slice, should be empty"))
	} else {
		fmt.Println("value is empty slice")
	}

	value = ValueOf([0]int{})
	if !value.Empty() {
		panic(fmt.Sprintf("vlaue is empty array, should be empty"))
	} else {
		fmt.Println("value is empty array")
	}

	value = ValueOf(0)
	if !value.Empty() {
		panic(fmt.Sprintf("value is 0, should be empty"))
	} else {
		fmt.Println("value is 0")
	}

	value = ValueOf(0.0)
	if !value.Empty() {
		panic(fmt.Sprintf("value is 0.0, should be empty"))
	} else {
		fmt.Println("value is 0.0")
	}

	value = ValueOf("")
	if !value.Empty() {
		panic(fmt.Sprintf("value is empty string, should be empty"))
	} else {
		fmt.Println("value is not empty, its value is empty string")
	}

	value = ValueOf(" ")
	if value.Empty() {
		panic(fmt.Sprintf("value is empty space string, not empty"))
	} else {
		fmt.Println("value is empty space string")
	}

	value = ValueOf(false)
	if !value.Empty() {
		panic(fmt.Sprintf("value is boolean false, should be empty"))
	} else {
		fmt.Println("value is boolean false")
	}

	value = ValueOf([]int{1,2})
	if value.Empty() {
		panic(fmt.Sprintf("value is int slice with 2 elements, not empty"))
	} else {
		fmt.Println("value is not empty, ", value.Value())
	}

	value = ValueOf(make(ZMap))
	if !value.Empty() {
		panic(fmt.Sprintf("value is empty map, should be empty"))
	} else {
		fmt.Println("value is empty map")
	}

	value = ValueOf(0.12)
	if value.Empty() {
		panic(fmt.Sprintf("value is 0.12, not empty"))
	} else {
		fmt.Println("value is 0.12, not empty")
	}
}

func TestZValue_Flatten(t *testing.T) {
	value := ValueOf(nil)
	fv := value.Flatten()
	if fv != nil {
		panic(fmt.Sprintf("value is nil, flatten should be nil too, assertion %v", fv))
	} else {
		fmt.Println("value is nil, so flatten value is nil")
	}

	value = ValueOf([]int{1,2,3})
	fv = value.Flatten()
	if len(fv) != value.Len() {
		panic(fmt.Sprintf("value is int slice with length %v, but length of flatten value is %v", value.Len(), len(fv)))
	} else if fv[0] != 1 || fv[1] != 2 || fv[2] != 3 {
		panic(fmt.Sprintf("flatten value should be [1,2,3], assertion is %v", fv))
	}
	fmt.Println("value is int slice with value ", value.Value(), " and flatten value is ", fv)

	value = ValueOf([2]string{"12x", "456"})
	fv = value.Flatten()
	if len(fv) != value.Len() {
		panic(fmt.Sprintf("value is string array with length %v, but length of flatten value is %v", value.Len(), len(fv)))
	} else if fv[0] != "12x" || fv[1] != "456" {
		panic(fmt.Sprintf("flatten value should be [123, 456], assertion is %v", fv))
	}
	fmt.Println("value is string array with value ", value.Value(), " and flatten value is ", fv)

	value = ValueOf(map[string]string{"k1":"key 1", "k2" : "key 2"})
	fv = value.Flatten()
	if len(fv) != value.Len() {
		panic(fmt.Sprintf("value is map with length %v, but length of flatten value is %v", value.Len(), len(fv)))
	} else if fv[0] == "key 1" && fv[1] == "key 2" {
		fmt.Println("value is map with value ", value.Value(), " and flatten value is ", fv)
	} else if fv[0] == "key 2" && fv[1] == "key 1" {
		fmt.Println("value is map with value ", value.Value(), " and flatten value is ", fv)
	} else {
		panic(fmt.Sprintf("flatten value should be [key 1, key 2] or [key 2, key 1], assertion is %v", fv))
	}

	value = ValueOf([]interface{}{
		"123", 123, []int{1,2,3}, []float32{1.5, 20.5}, map[int]int{1:123, 3:678}, []interface{}{
			"456", 890, [3]int{},
		},
	})
	fv = value.Flatten()
	if len(fv) != 14 {
		panic(fmt.Sprintf("length of flatten value should be 14, assertion is %v", len(fv)))
	}

	if fv[0] != "123" {
		panic(fmt.Sprintf("flatten value 0 should be 123, assertion %v", fv[0]))
	}
	if fv[1] != 123 {
		panic(fmt.Sprintf("flatten value 1 should be 123, assertion %v", fv[1]))
	}
	if fv[2] != 1 || fv[3] != 2 || fv[4] != 3 {
		panic(fmt.Sprintf("flatten value 2 to value 4 should be [1, 2, 3], assertion %v", fv[2:5]))
	}
	if fv[5] != float32(1.5) || fv[6] != float32(20.5) {
		panic(fmt.Sprintf("flatten value 5 to value 6 should be [1.5, 20.5], assertion %v", fv[5:7]))
	}
	if !(fv[7] == 123 && fv[8] == 678) && !(fv[7] == 678 && fv[8] == 123) {
		panic(fmt.Sprintf("flatten value 7 to value 8 should be [123,678] or [678,123], assertion %v", fv[7:9]))
	}
	if fv[9] != "456" {
		panic(fmt.Sprintf("flatten value 9 should be 456, assertion %v", fv[9]))
	}
	if fv[10] != 890 {
		panic(fmt.Sprintf("flatten value 10 should be 890, assertion %v", fv[9]))
	}
	if fv[11] != 0 || fv[12]!=0 || fv[13]!=0 {
		panic(fmt.Sprintf("flatten value 11 to value 13 should be [0, 0, 0], assertion %v", fv[11:]))
	}
	fmt.Println("original value is ", value.Value(), " and flatten value is ", fv)

	value = ValueOf(123)
	fv = value.Flatten()
	if len(fv) != 1 {
		panic(fmt.Sprintf("length of flatten value should be 1, assertion %v", len(fv)))
	}

	if fv[0] != 123 {
		panic(fmt.Sprintf("flatten value should be [123], assertion %v", fv))
	}
	fmt.Println("original value is ", value.Value(), " and flatten value is ", fv)
}

func TestZValue_IsNil(t *testing.T) {
	value := ValueOf(nil)
	if !value.IsNil() {
		panic(fmt.Sprintf("value is nil, assertion: %v", value.IsNil()))
	} else {
		fmt.Println("value is nil, assertion true")
	}

	var a []int = nil
	value = ValueOf(a)
	if !value.IsNil() {
		panic(fmt.Sprintf("value is nil int slice, assertion: %v", value.IsNil()))
	} else {
		fmt.Println("value is nil slice, assertion true")
	}

	var b map[int]string
	value = ValueOf(b)
	if !value.IsNil() {
		panic(fmt.Sprintf("value is nil map, assertion: %v", value.IsNil()))
	} else {
		fmt.Println("value is nil map")
	}

	var c *int = nil
	value = ValueOf(c)
	if !value.IsNil() {
		panic(fmt.Sprintf("value is nil pointer, assertion: %v", value.IsNil()))
	} else {
		fmt.Println("value is nil pointer")
	}

	var myerr= new(d)
	myerr = nil
	var serr error = myerr
	value = ValueOf(serr)
	if !value.IsNil() {
		panic(fmt.Sprintf("value is nil, assertion: %v", value.IsNil()))
	} else {
		fmt.Println("value is nil, and nil equal result is : ", serr == nil)
	}

}

type d struct{}
func(*d) Error() string{
	return "d error"
}

func TestZValue_String(t *testing.T) {
	value := ValueOf(nil)
	fmt.Println(value.String())

	value = ValueOf([]int{1,2,3})
	fmt.Println(value.String())

	value = ValueOf(map[int]string{1:"123", 2:"345"})
	fmt.Println(value.String())

	value = ValueOf("hello world")
	fmt.Println(value.String())
}

type testStruct struct {
	F1   string
	F2   int
	f3   []int
}

func TestZValue_Field(t *testing.T) {
	defer func() {
		if r:=recover(); r!=nil {
			if sr, ok := r.(string); ok {
				if sr == "field f3 cannot export" {
					fmt.Println("field f3 cannot export")
					return
				} else if sr == "field f4 is not exists" {
					fmt.Println("field f4 is not exists")
					return
				}
			}
			panic(r)
		}
	}()

	value := ValueOf(testStruct{
		F1: "field 1", F2: 2, f3: nil,
	})

	fmt.Println("field f1 values ", value.Field("F1").Value())
	fmt.Println("field f2 values ", value.Field("F2").Value())
	fmt.Println("field f4 values ", value.Field("f4").Value())
	fmt.Println("field f3 values ", value.Field("f3").Value())
}

func TestZValue_IsArray(t *testing.T) {
	value := ValueOf(nil)
	fmt.Println("value nil is array, assertion: ", value.IsArray())

	value = ValueOf([]int{1,2,3})
	fmt.Println("value int slice is array, assertion: ", value.IsArray())

	value = ValueOf([2]string{})
	fmt.Println("value string array is array, assertion: ", value.IsArray())

	value = ValueOf(1)
	fmt.Println("value int is array, assertion: ", value.IsArray())

	value = ValueOf(map[string]string{})
	fmt.Println("value map is array, assertion: ", value.IsArray())

	var a []int = nil
	value = ValueOf(a)
	fmt.Println("value of nil slice is array, assertion: ", value.IsArray())
}

func TestZValue_IsNumeric(t *testing.T) {
	value := ValueOf(0)
	fmt.Println("value ", value.Value(), " is numberic, assertion: ", value.IsNumeric())

	value = ValueOf(0.89)
	fmt.Println("value ", value.Value(), " is numberic, assertion: ", value.IsNumeric())

	value = ValueOf("13.4")
	fmt.Println("value ", value.Value(), " is numberic, assertion: ", value.IsNumeric())

	value = ValueOf("xx")
	fmt.Println("value ", value.Value(), " is numberic, assertion: ", value.IsNumeric())

	value = ValueOf(nil)
	fmt.Println("value ", value.Value(), " is numberic, assertion: ", value.IsNumeric())

	value = ValueOf([]int{1})
	fmt.Println("value ", value.Value(), " is numberic, assertion: ", value.IsNumeric())
}

func TestZValue_IsStruct(t *testing.T) {
	value := ValueOf(nil)
	fmt.Println("value ", value.Value(), " is struct, assertion: ", value.IsStruct())

	value = ValueOf(map[string]interface{}{})
	fmt.Println("value ", value.Value(), " is struct, assertion: ", value.IsStruct())

	value = ValueOf(testStruct{})
	fmt.Println("vlaue ", value.Value(), " is struct, assertion: ", value.IsStruct())

	var err error
	value = ValueOf(err)
	fmt.Println("value ", value.Value(), " is struct, assertion: ", value.IsStruct())

	err = errors.New("my error")
	value = ValueOf(err)
	fmt.Println("value ", value.Value(), " is struct, assertion: ", value.IsStruct())

	value = ValueOf(errors.New("my error"))
	fmt.Println("value ", value.Value(), " is struct, assertion: ", value.IsStruct())
}

func TestZValue_Len(t *testing.T) {
	value := ValueOf([]int{1,2,3})
	fmt.Println("length of value ", value.Value(), " is ", value.Len())

	value = ValueOf("12345")
	fmt.Println("length of value ", value.Value(), " is ", value.Len())

	value = ValueOf(map[string]interface{}{"s":"ddd"})
	fmt.Println("length of value ", value.Value(), " is ", value.Len())

	value = ValueOf([4]int{1,2})
	fmt.Println("length of value ", value.Value(), " is ", value.Len())
}

type tstring string

func TestZValue_IsString(t *testing.T) {
	value := ValueOf("23")
	fmt.Println("value ", value.Value(), " is string, assertion: ", value.IsString())

	value = ValueOf("")
	fmt.Println("value ", value.Value(), " is string, assertion: ", value.IsString())

	value = ValueOf(tstring("123"))
	fmt.Println("value ", value.Value(), " is string, assertion: ", value.IsString())

	value = ValueOf(12)
	fmt.Println("vallue ", value.Value(), " is string, assertion: ", value.IsString())

	value = ValueOf([]string{"345"})
	fmt.Println("value ", value.Value(), " is string, assertion: ", value.IsString())
}

func reset() {
	testMap = nil
}