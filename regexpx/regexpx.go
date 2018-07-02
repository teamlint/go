package regexpx

import "regexp"

// Compile 根据pattern生成对应的regexp正则对象
func Compile(pattern string) (*regexp.Regexp, error) {
	return regexp.Compile(pattern)
}

// Validate 校验所给定的正则表达式是否符合规范
func Validate(pattern string) error {
	_, err := Compile(pattern)
	return err
}

// IsMatch 正则表达式是否匹配
func IsMatch(pattern string, src []byte) bool {
	if r, err := Compile(pattern); err == nil {
		return r.Match(src)
	}
	return false
}

// IsMatchString 正则表达式是否匹配
func IsMatchString(pattern string, src string) bool {
	return IsMatch(pattern, []byte(src))
}

// FindMatchString 正则匹配，并返回匹配的列表
func FindMatchString(pattern string, src string) ([]string, error) {
	r, err := Compile(pattern)
	if err == nil {
		return r.FindStringSubmatch(src), nil
	}
	return nil, err
}

// FindAllMatchString 正则匹配，并返回所有匹配的列表
func FindAllMatchString(pattern string, src string) ([][]string, error) {
	r, err := Compile(pattern)
	if err == nil {
		return r.FindAllStringSubmatch(src, -1), nil
	}
	return nil, err
}

// Replace 正则替换(全部替换)
func Replace(pattern string, src []byte, replace []byte) ([]byte, error) {
	r, err := Compile(pattern)
	if err == nil {
		return r.ReplaceAll(src, replace), nil
	}
	return nil, err
}

// ReplaceString 正则替换(全部替换)，字符串
func ReplaceString(pattern, src string, replace string) (string, error) {
	r, e := Replace(pattern, []byte(src), []byte(replace))
	return string(r), e
}
