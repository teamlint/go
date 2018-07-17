package convert

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
