package html

import "testing"

func TestSelect(t *testing.T) {
	// enum := EnumActionDelete
	// t.Logf("enum type: %s\n", reflect.TypeOf(enum).String())
	// t.Logf("enum value: %v\n", reflect.ValueOf(enum).Interface())
	// t.Logf("enum string: %v\n", reflect.ValueOf(enum))

	// m := EnumActionCreate.Map()
	// t.Logf("EnumAction Map:%v", m)
	html := Select(EnumActionUpdate, "EnumAction", nil)
	t.Logf("EnumActionUpdate Select:%v", html)
	html = Select(EnumActionAll, "EnumAction", nil)
	t.Logf("EnumActionAll Select:%v", html)
	html = SelectWith(EnumActionRetrive, "abc", nil, "", "")
	t.Logf("EnumActionAll Select:%v", html)
	html = SelectWith(EnumActionCreate, "abc", nil, "", "", "class=\"btn btn-primary\" data-id=\"sdfasdfas\"")
	t.Logf("EnumActionAll Select:%v", html)
	html = Select(&Foo{}, "foo", nil, "", "")
	t.Logf("Foo Select:%v", html)
	mapAttrs := map[string]interface{}{"class": "box", "readonly": "readonly"}
	html = SelectWith(EnumActionDelete, "delete", nil, "", "", mapAttrs, "data-num='12'", "data-pid=\"10\"")
	t.Logf("EnumActionDelete Select:%v", html)
	options := []Option{
		{Text: "选项一", Value: "1"},
		{Text: "选项二", Value: "2", IsSelected: true},
		{Text: "选项三", Value: "3"},
	}
	// html = SelectWith(EnumActionNone, "all", options, "", "", mapAttrs, "data-num='12'", "data-pid=\"10\"")
	html = SelectWith(EnumActionNone, "all", options, "", "", nil, "data-num='12'", "data-pid=\"10\"")
	t.Logf("EnumActionNone with data Select:%v", html)
	html = Select(EnumActionNone, "all", nil, "data-pid=\"100\"")
	t.Logf("EnumActionNone with data Select:%v", html)
}
