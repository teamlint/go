package html

import (
	"bytes"
	"html/template"
)

// Option 下拉列表选择项
type Option struct {
	Text       string            // 文本
	Value      string            // 值
	IsSelected bool              // 是否选中
	Attrs      map[string]string // 属性键值对,如: {"class":"btn","data-id"="a1"]
}

// String 下拉列表项字符串
func (o Option) String() string {
	buf := bytes.Buffer{}
	buf.WriteString(`<option value="`)
	buf.WriteString(o.Value)
	buf.WriteString(`"`)
	// selected
	if o.IsSelected {
		buf.WriteString(" ")
		buf.WriteString(`selected="selected"`)
	}
	// html attributes
	for key, val := range o.Attrs {
		buf.WriteString(" ")
		buf.WriteString(key)
		buf.WriteString(`="`)
		buf.WriteString(val)
		buf.WriteString(`"`)
	}
	buf.WriteString(`>`)
	buf.WriteString(o.Text)
	buf.WriteString(`</option>`)
	return buf.String()
}

// SelectOption 下拉列表项HTML文本
func SelectOption(opt Option) template.HTML {
	return template.HTML(opt.String())
}
