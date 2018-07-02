package mapx

import (
	"reflect"

	"github.com/teamlint/gox/stringx"
)

// ToStruct 将map键值对映射到对应的struct对象属性上，需要注意：
// 1、第二个参数为struct对象指针；
// 2、struct对象的公开属性才能被映射赋值；
// 3、map中的键名可以为小写，映射转换时会自动将键名首字母转为大写做匹配映射，如果无法匹配则忽略；
func ToStruct(m map[string]interface{}, o interface{}) error {
	for k, v := range m {
		mapToStructSetField(o, k, v)
	}
	return nil
}
func mapToStructSetField(obj interface{}, name string, value interface{}) {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByName(stringx.UpperFirst(name))
	// 键名与对象属性匹配检测
	if !structFieldValue.IsValid() {
		//return fmt.Errorf("No such field: %s in obj", name)
		return
	}
	// CanSet的属性必须为公开属性(首字母大写)
	if !structFieldValue.CanSet() {
		//return fmt.Errorf("Cannot set %s field value", name)
		return
	}
	// 必须将value转换为struct属性的数据类型
	structFieldValue.Set(reflect.ValueOf(Convert(value, structFieldValue.Type().String())))
}
