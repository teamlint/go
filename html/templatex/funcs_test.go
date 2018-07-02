package templatex

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/teamlint/gox/stringx"
	"github.com/teamlint/gox/timex"
)

type EnumMethod = string
type EnumAction int
type Foo struct {
	Name       string
	Age        int
	IsApproved bool
}

const (
	EnumActionCreate  EnumAction = 1
	EnumActionRetrive EnumAction = 2
	EnumActionUpdate  EnumAction = 3
	EnumActionDelete  EnumAction = 4
	EnumActionAll     EnumAction = -1
	EnumActionNone    EnumAction = 0

	EnumMethodGet  EnumMethod = "GET"
	EnumMethodPost EnumMethod = "POST"
)

func (enum EnumAction) Text() string {
	/*
		switch enum {
		case EnumActionCreate:
			return "创建"
		case EnumActionRetrive:
			return "查找"
		case EnumActionUpdate:
			return "更新"
		case EnumActionDelete:
			return "删除"
		}
		return ""
	*/
	if val, ok := actionMap[enum]; ok {
		return val
	}
	return ""
}

var actionMap = map[EnumAction]string{
	EnumActionAll:     "全部",
	EnumActionNone:    "未设置",
	EnumActionCreate:  "创建",
	EnumActionRetrive: "查找",
	EnumActionUpdate:  "更新",
	EnumActionDelete:  "删除",
}

func (enum EnumAction) Map() map[EnumAction]string {
	return actionMap
}
func (enum EnumAction) SelectList() []EnumAction {
	var items = make([]EnumAction, 0)
	items = append(items, EnumActionAll)
	items = append(items, EnumActionCreate)
	items = append(items, EnumActionRetrive)
	items = append(items, EnumActionUpdate)
	items = append(items, EnumActionDelete)
	return items
}

var EnumMethodText = map[EnumMethod]string{
	EnumMethodGet:  "获取",
	EnumMethodPost: "回发",
}

// func (enum EnumMethod) String() string {
// 	switch enum {
// 	case EnumMethodGet:
// 		return "获取"
// 	case EnumMethodPost:
// 		return "回发"
// 	}
// 	return ""
// }
func (f *Foo) Text() string {
	return fmt.Sprintf("foo:{Name:%v,Age:%v,IsApproved:%v}", f.Name, f.Age, f.IsApproved)
}
func (f *Foo) List() []*Foo {
	var items []*Foo
	for i := 0; i < 6; i++ {
		item := new(Foo)
		item.Name = "Item " + strconv.Itoa(i)
		items = append(items, item)
	}
	return items
}

func TestFunc(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	datetimeString := "2018-03-30 16:24:58"
	expectDatetime := "2018-03-30 16:24:58"
	expectDate := "2018-03-30"
	expectTime := "16:24:58"
	dt, _ := time.Parse(layout, datetimeString)
	actual := timex.Format(dt)
	if actual != expectDatetime {
		t.Errorf("expect value:%v,actual value:%v", expectDatetime, actual)
	}
	t.Logf("actual value:%v", actual)
	actual = timex.Format(dt, timex.DateFormat)
	if actual != expectDate {
		t.Errorf("expect value:%v,actual value:%v", expectDate, actual)
	}
	t.Logf("actual value:%v", actual)
	actual = timex.Format(dt, timex.TimeFormat)
	if actual != expectTime {
		t.Errorf("expect value:%v,actual value:%v", expectTime, actual)
	}
	t.Logf("actual value:%v", actual)

	now := time.Now()
	actual = timex.Format(now, "2006/01/02 15H04:05")
	t.Logf("actual value:%v", actual)
	actual = timex.Format(time.Time{}, "2006/01/02 15H04:05", "no date")
	t.Logf("actual value:%v", actual)

	actual = Year()
	expectValue := "2018"
	if actual != expectValue {
		t.Errorf("expect value:%v,actual value:%v", expectValue, actual)
	}
	t.Logf("actual value:%v", actual)

	actual = Month()
	expectValue = "7"
	if actual != expectValue {
		t.Errorf("expect value:%v,actual value:%v", expectValue, actual)
	}
	t.Logf("actual value:%v", actual)

	actual = Day()
	expectValue = "2"
	if actual != expectValue {
		t.Errorf("expect value:%v,actual value:%v", expectValue, actual)
	}
	t.Logf("actual value:%v", actual)
}
func TestEnum(t *testing.T) {
	action := EnumActionNone
	method := EnumMethodPost
	foo := Foo{Name: "Foo结构", Age: 40, IsApproved: true}
	t.Logf("[%v]action update:%s", action, stringx.Text(action))
	t.Logf("[%v]method post:%s", method, stringx.Text(method))
	t.Logf("[%v]struct:%s", foo, stringx.Text(&foo))

}
