package templatex

import (
	"html/template"
	"net/url"
	"strconv"
	"time"

	"github.com/teamlint/gox/html/dom/tag"
	"github.com/teamlint/gox/html/dom/tag/pager"
	"github.com/teamlint/gox/timex"
)

// DatetimeFormat 日期时间格式化 // arg: 可包含三个参数 1、时间格式（time.Time或*time.Time) 2、格式化字符串(string) 3、空值时返回值字符串(string)
func DatetimeFormat(v ...interface{}) string {
	if len(v) > 0 {
		format := timex.DatetimeFormat // "2006-01-02 15:04:05"
		if len(v) > 1 {
			if v[1] != "" {
				format = v[1].(string)
			}
		}
		switch val := v[0].(type) {
		case time.Time:
			return val.Format(format)
		case *time.Time:
			if val != nil {
				return val.Format(format)
			}
		}
		if len(v) > 2 {
			return v[2].(string)
		}
	}
	return "-"
}

// DatetimePretty 日期时间友好格式化
// arg: 可包含三个参数 1、时间格式（time.Time或*time.Time) 2、超过友好格式化范围的格式化字符串(string) 3、空值时返回值字符串(string)
func DatetimePretty(v ...interface{}) string {
	if len(v) > 0 {
		format := timex.DatetimeFormat // "2006-01-02 15:04:05"
		if len(v) > 1 {
			if v[1] != "" {
				format = v[1].(string)
			}
		}
		switch val := v[0].(type) {
		case time.Time:
			return timex.Pretty(val, format)
		case *time.Time:
			if val != nil {
				return timex.Pretty(*val, format)
			}
		}
		if len(v) > 2 {
			return v[2].(string)
		}
	}
	return "-"
}

// DateFormat 格式化日期,第一个参数为日期数据,第二个参数为格式化字符串(如果有),第三个参数为默认值(如果有)
func DateFormat(v ...interface{}) string {
	if len(v) > 0 {
		format := timex.DateFormat // "2006-01-02"
		if len(v) > 1 {
			if v[1] != "" {
				format = v[1].(string)
			}
		}
		switch val := v[0].(type) {
		case time.Time:
			if !val.IsZero() {
				return val.Format(format)
			}
		case *time.Time:
			if val != nil && !val.IsZero() {
				return val.Format(format)
			}
		}
		if len(v) > 2 {
			return v[2].(string)
		}
	}
	return "-"
}

// TimeFormat 日期时间格式化 // arg: 可包含三个参数 1、时间格式（time.Time或*time.Time) 2、格式化字符串(string) 3、空值时返回值字符串(string)
func TimeFormat(v ...interface{}) string {
	if len(v) > 0 {
		format := timex.TimeFormat // "15:04:05"
		if len(v) > 1 {
			if v[1] != "" {
				format = v[1].(string)
			}
		}
		switch val := v[0].(type) {
		case time.Time:
			return val.Format(format)
		case *time.Time:
			if val != nil {
				return val.Format(format)
			}
		}
		if len(v) > 2 {
			return v[2].(string)
		}
	}
	return "-"
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
