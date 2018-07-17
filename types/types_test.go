package types

import (
	"reflect"
	"testing"
)

func TestTypes(t *testing.T) {
	mapIntString := make(map[int]string)
	t.Logf("map[int]string type: %T,%v", mapIntString, reflect.TypeOf(mapIntString))
	pinfo(mapIntString, t)
	mapStringInterface := make(map[string]interface{})
	pinfo(mapStringInterface, t)

	mapStringBool := make(map[string]bool)
	// mapx := reflect.MakeMap(reflect.TypeOf(mapStringBool))
	toType := reflect.TypeOf(mapStringBool)
	mapx := reflect.MakeMap(toType)
	key := reflect.ValueOf("key")
	val := reflect.ValueOf(true)
	key2 := reflect.ValueOf("中文键")
	val2 := reflect.ValueOf(false)
	// 设置键值
	mapx.SetMapIndex(key, val)
	mapx.SetMapIndex(key2, val2)
	// 遍历map
	keys := mapx.MapKeys()
	for _, k := range keys {
		t.Logf("mapx key:%v, value:%v", k.Interface(), mapx.MapIndex(k).Interface())
	}
}
func pinfo(i interface{}, t *testing.T) {
	ri := reflect.ValueOf(i)
	t.Logf("name: %v\ttype: %v\tkind: %v,", ri.Interface(), ri.Type().String(), ri.Kind().String())
}
