package timex

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	a := time.Date(2018, 6, 30, 16, 39, 45, 123, time.Local)
	e := "2018-06-30 16:39:45"
	assert.Equal(t, e, Format(a), "they should be equal")
	assert.Equal(t, "2018-06-30", Format(a, DateFormat), "they should be equal")
	assert.Equal(t, "-", Format(time.Time{}, DateFormat, "-"), "they should be equal")
	// assert.Nil(t, object)
}

func TestPretty(t *testing.T) {
	tests := []struct {
		name string
		t    time.Time
		want string
	}{
		{name: "Just now", t: time.Now(), want: "刚刚"},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(2)),
			want: "2秒后",
		},
		{name: "SecondAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-1)),
			want: "1秒前"},
		{name: "Minutes", t: time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59)), want: "60分钟后"},
		{name: "Tomorrow", t: time.Now().AddDate(0, 0, 1), want: "明天"},
		{name: "Yesterday", t: time.Now().AddDate(0, 0, -1), want: "昨天"},
		{name: "2day", t: time.Now().AddDate(0, 0, 2), want: "2天后"},
		{name: "2dayAgo", t: time.Now().AddDate(0, 0, -2), want: "2天前"},
		{name: "Week", t: time.Now().AddDate(0, 0, 7), want: Format(time.Now().AddDate(0, 0, 7))},
		{name: "WeekAgo", t: time.Now().AddDate(0, 0, -7), want: Format(time.Now().AddDate(0, 0, -7))},
		{name: "Month", t: time.Now().AddDate(0, 1, 0), want: Format(time.Now().AddDate(0, 1, 0))},
		{name: "MonthAgo", t: time.Now().AddDate(0, -1, 0), want: Format(time.Now().AddDate(0, -1, 0))},
		{name: "Year", t: time.Now().AddDate(50, 0, 0), want: Format(time.Now().AddDate(50, 0, 0))},
		{name: "YearAgo", t: time.Now().AddDate(-2, 0, 0), want: Format(time.Now().AddDate(-2, 0, 0))},
		{name: "OvertimeFormatter", t: time.Now().AddDate(0, 10, 0), want: Format(time.Now().AddDate(0, 10, 0), "2006/01/02")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "OvertimeFormatter" {
				gotTimeSince := Pretty(tt.t, "2006/01/02")
				t.Logf("PrettyWithFormatter() = %v, want %v", gotTimeSince, tt.want)
				if gotTimeSince != tt.want {
					t.Errorf("Pretty() = %v, want %v", gotTimeSince, tt.want)
				}
			} else {
				gotTimeSince := Pretty(tt.t)
				t.Logf("Pretty() = %v, want %v", gotTimeSince, tt.want)
				if gotTimeSince != tt.want {
					t.Errorf("Pretty() = %v, want %v", gotTimeSince, tt.want)
				}
			}
		})
	}
}

func ExamplePretty() {
	timeSlots := []struct {
		name string
		t    time.Time
	}{
		{name: "Just now", t: time.Now()},
		{name: "Second", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(1)),
		},
		{name: "SecondAgo", t: time.Now().Add(
			time.Hour*time.Duration(0) +
				time.Minute*time.Duration(0) +
				time.Second*time.Duration(-1)),
		},
		{name: "Minutes", t: time.Now().Add(time.Hour*time.Duration(0) +
			time.Minute*time.Duration(59) +
			time.Second*time.Duration(59))},
		{name: "Tomorrow", t: time.Now().AddDate(0, 0, 1)},
		{name: "Yesterday", t: time.Now().AddDate(0, 0, -1)},
		{name: "Week", t: time.Now().AddDate(0, 0, 7)},
		{name: "WeekAgo", t: time.Now().AddDate(0, 0, -7)},
		{name: "Month", t: time.Now().AddDate(0, 1, 0)},
		{name: "MonthAgo", t: time.Now().AddDate(0, -1, 0)},
		{name: "Year", t: time.Now().AddDate(2, 0, 0)},
		{name: "YearAgo", t: time.Now().AddDate(-2, 0, 0)},
	}

	for _, timeSlot := range timeSlots {
		fmt.Printf("%s = %v\n", timeSlot.name, Format(timeSlot.t))
	}
}
