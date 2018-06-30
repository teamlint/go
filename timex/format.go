package timex

import "time"

// Format 日期时间格式化
// 可包含三个参数 1、时间对象 2、格式化字符串(string) 3、空值时返回值字符串(string)
func Format(t time.Time, f ...string) string {
	if t.IsZero() {
		if len(f) > 1 {
			return f[1]
		}
		return ""
	}
	format := DatetimeFormat
	if len(f) > 0 {
		format = f[0]
	}

	return t.Format(format)
}
