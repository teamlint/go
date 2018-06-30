package attr

// Checked 根据bool参数值返回html属性checked
func Checked(v bool) string {
	if v {
		return "checked"
	}
	return ""
}

// Selected 根据参数值返回html属性seleted
func Selected(v int, checkValue int) string {
	if v == checkValue {
		return "selected"
	}
	return ""
}

// Readonly 根据bool参数值返回html属性readonly
func Readonly(v bool) string {
	if v {
		return "readonly"
	}
	return ""
}

// Disabled 根据bool参数值返回html属性disabled
func Disabled(v bool) string {
	if v {
		return "disabled"
	}
	return ""
}
