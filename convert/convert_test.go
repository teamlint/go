package convert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teamlint/gox/timex"
)

func TestRound(t *testing.T) {
	assert := assert.New(t)
	vals := make(map[float64]float64)
	vals[3.1455] = 3.15
	vals[3.1451] = 3.15
	vals[3.1458] = 3.15
	vals[3.1448] = 3.14
	vals[3.1449] = 3.14
	vals[3.1449] = 3.14
	vals[-3.1451] = -3.15
	vals[-3.1449] = -3.14
	vals[-3.0000] = -3.00
	for k, v := range vals {
		act := Round(k, 2)
		fmt.Printf("%v round  %v =  %v: %v\n", k, act, v, act == v)
		assert.Equal(v, act)
	}
	var k, exp float64
	k = 3.5415
	exp = 4
	act := Round(k, 0)
	fmt.Printf("%v round  %v =  %v: %v\n", k, act, exp, act == exp)
	assert.Equal(4.0, act)
	k = 0
	exp = 0
	act = Round(k, 0)
	fmt.Printf("%v round  %v =  %v: %v\n", k, act, exp, act == exp)
	assert.Equal(exp, act)
}
func TestRoundInt(t *testing.T) {
	assert := assert.New(t)
	vals := make(map[float64]int)
	vals[3.1455] = 315
	vals[3.1451] = 315
	vals[3.1458] = 315
	vals[3.1448] = 314
	vals[3.1449] = 314
	vals[3.1449] = 314
	for k, v := range vals {
		act := RoundInt(k, 2)
		fmt.Printf("%v round  %v =  %v: %v\n", k, act, v, act == v)
		assert.Equal(v, act)
	}

}
func TestConvert(t *testing.T) {
	timestr := "2016-11-30T12:15:53"
	jsonTime := ToTime(timestr, "2006-01-02T15:04:05")
	t.Log(jsonTime)
	now := ToTime("2020-06-15 10:43:29.168542123")
	// nanosecond
	ns := 1592189009168542123
	t.Log("nanosecond ", ToTime(ns))
	assert.Equal(t, now.UnixNano(), ToTime(ns).UnixNano())
	// microsecond
	mis := 1592189009168542
	t.Log("microsecond ", ToTime(mis))
	assert.Equal(t, now.Unix(), ToTime(mis).Unix())
	// millisecond
	ms := 1592189009168
	t.Log("millisecond ", ToTime(ms))
	assert.Equal(t, now.Unix(), ToTime(ms).Unix())
	// second
	epoch := 1592189009
	t.Log("second ", ToTime(epoch))
	assert.Equal(t, now.Unix(), ToTime(epoch).Unix())
	jsJsonTime := "2020-06-15T10:43:29.168Z"
	t.Log("jsJSON time", ToTime(jsJsonTime))
	assert.Equal(t, ToInt64(ms), timex.Timestamp(ToTime(jsJsonTime)))

}
