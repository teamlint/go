package html

import (
	"fmt"
	"html/template"
	"net/url"
	"strconv"
	"time"

	"github.com/teamlint/lib/html/pager"
)

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

// Raw 原始HTML
func Raw(s string) template.HTML {
	return template.HTML(s)
}

// URLPager 分页器
func URLPager(url string, total int, pageIndex int, pageSize ...int) template.HTML {
	urlPager := pager.NewUrlPager(url, total, pageIndex, pageSize...)
	return template.HTML(urlPager.PagerString())
}

// URL 原始URL
func URL(url string) template.URL {
	return template.URL(url)
}

// URLDecode 解码
func URLDecode(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return uri
	}
	return u.String()
}

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

// Year 获取当前年度
func Year() string {
	return strconv.Itoa(time.Now().Year())
}

// Month 获取当前月份
func Month() string {
	return strconv.Itoa(int(time.Now().Month()))
}

// Day 获取当前日
func Day() string {
	return strconv.Itoa(time.Now().Day())
}

// Now 当前时间
func Now() time.Time {
	return time.Now()
}
