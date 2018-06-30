package templatex

import (
	"html/template"
	"net/url"
	"strconv"
	"time"

	"github.com/teamlint/gox/html/dom/tag"
	"github.com/teamlint/gox/html/dom/tag/pager"
)

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

// FormBegin 表单开始
func FormBegin(action string, method string, opts ...string) template.HTML {
	return tag.FormBegin(action, method, opts...)
}

// FormEnd 表单结束
func FormEnd() template.HTML {
	return tag.FormEnd()
}

// SelectWith 下拉列表,附带选择文本
func SelectWith(v interface{}, name string, selectOptions []tag.Option, selectText string, selectValue string, attrs ...interface{}) template.HTML {
	return tag.SelectWith(v, name, selectOptions, selectText, selectValue, attrs...)
}

// Select 下拉列表
func Select(v interface{}, name string, selectOptions []tag.Option, attrs ...interface{}) template.HTML {
	return tag.Select(v, name, selectOptions, attrs...)
}

// SelectOption 下拉列表项HTML文本
func SelectOption(opt tag.Option) template.HTML {
	return tag.SelectOption(opt)
}
