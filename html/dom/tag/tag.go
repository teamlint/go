package tag

import (
	"fmt"
	"html/template"
)

// FormBegin 表单开始
func FormBegin(action string, method string, opts ...string) template.HTML {
	form := fmt.Sprintf(`<form action="%s" method="%s"`, action, method)
	if len(opts) > 0 {
		for _, v := range opts {
			form += " " + v
		}
	}
	form += ">"
	return template.HTML(form)
}

// FormEnd 表单结束
func FormEnd() template.HTML {
	return template.HTML("</form>")
}
