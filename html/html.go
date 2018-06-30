package html

import (
	"github.com/teamlint/lib"
)

// AddFuncs 添加模板方法
func AddFuncs(funcer Funcer) {
	funcer.AddFunc("datetimeFormat", lib.DatetimeFormat)
	funcer.AddFunc("dateFormat", lib.DateFormat)
	funcer.AddFunc("timeFormat", lib.TimeFormat)
	funcer.AddFunc("raw", Raw)
	funcer.AddFunc("pager", URLPager)
	funcer.AddFunc("url", URL)
	funcer.AddFunc("urlDecode", URLDecode)
	funcer.AddFunc("formBegin", FormBegin)
	funcer.AddFunc("formEnd", FormEnd)
	funcer.AddFunc("checked", Checked)
	funcer.AddFunc("selected", Selected)
	funcer.AddFunc("readonly", Readonly)
	funcer.AddFunc("disabled", Disabled)
	funcer.AddFunc("now", Now)
	funcer.AddFunc("year", Year)
	funcer.AddFunc("month", Month)
	funcer.AddFunc("day", Day)
	funcer.AddFunc("text", lib.Text)
	funcer.AddFunc("select", Select)
	funcer.AddFunc("selectWith", SelectWith)
	funcer.AddFunc("selectOption", SelectOption)
}
