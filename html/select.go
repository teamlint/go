package html

import (
	"bytes"
	"fmt"
	"html/template"
	"reflect"
)

// SelectWith 添加顶级选择项的下拉列表
func SelectWith(v interface{}, name string, selectOptions []Option, selectText string, selectValue string, attrs ...interface{}) template.HTML {
	return selectList(v, name, selectOptions, selectText, selectValue, attrs...)
}
func Select(v interface{}, name string, selectOptions []Option, attrs ...interface{}) template.HTML {
	return selectList(v, name, selectOptions, "", "", attrs...)
}

// SelectList 生成下拉列表选择框
// name: 元素name,selectOptions 下拉列表项列表数据源,selectText: 自定义最顶层项文本,selectValue: 自定义最顶层项值,attrs: html元素属性,
func selectList(v interface{}, name string, selectOptions []Option, selectText string, selectValue string, attrs ...interface{}) template.HTML {
	// func SelectList(v interface{}, name string, selectText string, selectValue string, attrs ...string) template.HTML {
	buf := bytes.Buffer{}
	// select begin
	buf.WriteString(`<select name="` + name + `" id="` + name + `"`)
	if len(attrs) > 0 {
		for _, attr := range attrs {
			av := reflect.ValueOf(attr)
			at := av.Kind()
			// fmt.Printf("attrs type: %v\n", at)
			switch at {
			// 如: "class='btn'" "data-id='abc'"
			case reflect.String:
				buf.WriteString(" ")
				buf.WriteString(attr.(string))
			// 如: {"class":"btn","data-id":"abc"}
			case reflect.Map:
				attrKeys := av.MapKeys()
				for _, key := range attrKeys {
					buf.WriteString(" ")
					buf.WriteString(key.Interface().(string))
					buf.WriteString(`="`)
					buf.WriteString(av.MapIndex(key).Interface().(string))
					buf.WriteString(`"`)
				}
			}
		}
	}
	buf.WriteString(">")
	// options
	// select text
	if selectText != "" {
		buf.WriteString(`<option value="` + selectValue + `">`)
		buf.WriteString(selectText)
		buf.WriteString(`</option>`)
	}
	// select items
	val := reflect.ValueOf(v)
	// 如果提供数据源
	if len(selectOptions) > 0 {
		for _, o := range selectOptions {
			// selected
			itemValue := fmt.Sprintf("%v", val)
			if itemValue == o.Value {
				o.IsSelected = true
			}
			buf.WriteString(o.String())
		}

	} else { // 自动查找列表方法
		t := val.Type()
		nums := val.NumMethod()
		for i := 0; i < nums; i++ {
			method := t.Method(i).Name
			// fmt.Printf("method: %v\n", method)
			if method == "SelectList" {
				result := val.Method(i).Call(nil)[0].Interface()
				// fmt.Printf("result type: %v\n", reflect.ValueOf(result).Kind())
				// fmt.Printf("result value : %v\n", reflect.ValueOf(result))
				if !(reflect.ValueOf(result).Kind() == reflect.Array || reflect.ValueOf(result).Kind() == reflect.Slice) {
					break
				}
				resultValue := reflect.ValueOf(result)
				len := resultValue.Len()
				for j := 0; j < len; j++ {
					option := resultValue.Index(j)
					// fmt.Printf("option item: %v\n", option)
					optionText := ""
					methodValue := option.MethodByName("Text")
					if !methodValue.IsValid() {
						optionText = fmt.Sprintf("%v", option)
						// fmt.Printf("option text1: %v\n", optionText)
					} else {
						optionText = methodValue.Call(nil)[0].Interface().(string)
						// fmt.Printf("option text2: %v\n", optionText)
					}
					optionValue := fmt.Sprintf("%v", option)
					// fmt.Printf("option type: %v\n", option.Type())
					// fmt.Printf("option value: %v\n", reflect.ValueOf(option).Interface())
					buf.WriteString(`<option value="`)
					buf.WriteString(optionValue)
					buf.WriteString(`"`)
					// selected
					itemValue := fmt.Sprintf("%v", val)
					if itemValue == optionValue {
						buf.WriteString(` selected="selected"`)
					}
					buf.WriteString(`>`)
					buf.WriteString(optionText)
					buf.WriteString(`</option>`)
				}
			}
		}
	}
	// select end
	buf.WriteString(`</select>`)

	return template.HTML(buf.String())
}
