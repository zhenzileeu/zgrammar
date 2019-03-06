package zgrammar

import (
	"testing"
	"fmt"
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

func reset() {
	testMap = nil
}