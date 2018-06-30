package timex

import (
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
