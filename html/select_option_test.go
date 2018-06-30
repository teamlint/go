package html

import "testing"

func TestSelectOption(t *testing.T) {
	o := Option{}
	t.Logf("select option: %s", o.String())
	o.Text = "选项一"
	o.Value = "1"
	t.Logf("select option: %s", o.String())
	o.IsSelected = true
	t.Logf("select option: %s", o.String())
	o.Attrs = map[string]string{"class": "btn", "data-id": "abc"}
	t.Logf("select option: %s", o.String())
}
