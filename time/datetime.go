package time

import "time"

// DatetimeFormat 日期时间格式化 // arg: 可包含三个参数 1、时间格式（time.Time或*time.Time) 2、格式化字符串(string) 3、空值时返回值字符串(string)
func DatetimeFormat(v ...interface{}) string {
	if len(v) > 0 {
		format := "2006-01-02 15:04:05"
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

// DateFormat 格式化日期,第一个参数为日期数据,第二个参数为格式化字符串(如果有),第三个参数为默认值(如果有)
func DateFormat(v ...interface{}) string {
	if len(v) > 0 {
		format := "2006-01-02"
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
		format := "15:04:05"
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
