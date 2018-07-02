package templatex

import (
	"github.com/teamlint/gox/html/dom/attr"
	"github.com/teamlint/gox/stringx"
)

// AddFuncs 添加模板方法
func AddFuncs(funcer Funcer) {
	funcer.AddFunc("datetimeFormat", DatetimeFormat)
	funcer.AddFunc("dateFormat", DateFormat)
	funcer.AddFunc("timeFormat", TimeFormat)
	funcer.AddFunc("raw", Raw)
	funcer.AddFunc("pager", URLPager)
	funcer.AddFunc("url", URL)
	funcer.AddFunc("urlDecode", URLDecode)
	funcer.AddFunc("formBegin", FormBegin)
	funcer.AddFunc("formEnd", FormEnd)
	funcer.AddFunc("checked", attr.Checked)
	funcer.AddFunc("selected", attr.Selected)
	funcer.AddFunc("readonly", attr.Readonly)
	funcer.AddFunc("disabled", attr.Disabled)
	funcer.AddFunc("now", Now)
	funcer.AddFunc("year", Year)
	funcer.AddFunc("month", Month)
	funcer.AddFunc("day", Day)
	funcer.AddFunc("text", stringx.Text)
	funcer.AddFunc("select", Select)
	funcer.AddFunc("selectWith", SelectWith)
	funcer.AddFunc("selectOption", SelectOption)
}
