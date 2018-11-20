package templatex

import (
	"github.com/teamlint/gox/html/dom/attr"
	"github.com/teamlint/gox/stringx"
)

// AddFuncs 添加模板方法
func AddFuncs(funcer Funcer) {
	funcer.AddFunc("datetimePretty", DatetimePretty)
	funcer.AddFunc("datetimeFormat", DatetimeFormat)
	funcer.AddFunc("dateFormat", DateFormat)
	funcer.AddFunc("timeFormat", TimeFormat)
	funcer.AddFunc("now", Now)
	funcer.AddFunc("year", Year)
	funcer.AddFunc("month", Month)
	funcer.AddFunc("day", Day)
	// iif
	funcer.AddFunc("iif", iif)
	// tags
	funcer.AddFunc("raw", Raw)
	funcer.AddFunc("pager", URLPager)
	funcer.AddFunc("url", URL)
	funcer.AddFunc("urlDecode", URLDecode)
	funcer.AddFunc("urlEncode", URLEncode)
	funcer.AddFunc("formBegin", FormBegin)
	funcer.AddFunc("formEnd", FormEnd)
	funcer.AddFunc("checked", attr.Checked)
	funcer.AddFunc("selected", attr.Selected)
	funcer.AddFunc("readonly", attr.Readonly)
	funcer.AddFunc("disabled", attr.Disabled)
	funcer.AddFunc("text", stringx.Text)
	funcer.AddFunc("select", Select)
	funcer.AddFunc("selectWith", SelectWith)
	funcer.AddFunc("selectOption", SelectOption)
	// string
	funcer.AddFunc("json", toJSON)
	funcer.AddFunc("length", Lengh)
	funcer.AddFunc("truncate", truncate)
	funcer.AddFunc("striptags", striptags)
	funcer.AddFunc("join", join)
	// data
	funcer.AddFunc("default", dfault)
	funcer.AddFunc("empty", empty)
	// dict
	funcer.AddFunc("dict", dict)
	funcer.AddFunc("set", set)
	funcer.AddFunc("unset", unset)
	funcer.AddFunc("hasKey", hasKey)
	funcer.AddFunc("pluck", pluck)
	funcer.AddFunc("keys", keys)
	funcer.AddFunc("pick", pick)
	funcer.AddFunc("omit", omit)
	funcer.AddFunc("values", values)
	// list
	funcer.AddFunc("list", list)
	funcer.AddFunc("append", push)
	funcer.AddFunc("prepend", prepend)
	funcer.AddFunc("push", push)
	funcer.AddFunc("first", first)
	funcer.AddFunc("rest", rest)
	funcer.AddFunc("last", last)
	funcer.AddFunc("initial", initial)
	funcer.AddFunc("reverse", reverse)
	funcer.AddFunc("uniq", uniq)
	funcer.AddFunc("without", without)
	funcer.AddFunc("has", has)
	funcer.AddFunc("slice", slice)
	funcer.AddFunc("random", random)
	// env
	funcer.AddFunc("env", env)
	funcer.AddFunc("expandenv", expandEnv)
	// reflect
	funcer.AddFunc("typeOf", typeOf)
	funcer.AddFunc("typeIs", typeIs)
	funcer.AddFunc("typeIsLike", typeIsLike)
	funcer.AddFunc("kindOf", kindOf)
	funcer.AddFunc("kindIs", kindIs)

}
