package xtime

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestXtime_Scan(t *testing.T) {
	c := NewXtime()
	err := c.Scan(time.Now())
	assert.Nil(t, err)
	assert.Equal(t, c.ToDateTimeString(), Now().ToDateTimeString())
}

func TestError_Scan(t *testing.T) {
	c, v := NewXtime(), "xxx"
	err := c.Scan(v)
	assert.Equal(t, err, fmt.Errorf("can not convert %v to xtime", v))
}

func TestXtime_Value(t *testing.T) {
	c := Now()
	v, err := c.Value()
	assert.Nil(t, err)
	assert.Equal(t, v, c.time)
}

func TestError_Value(t *testing.T) {
	c := Parse("")
	v, err := c.Value()
	assert.Nil(t, err)
	assert.Equal(t, v, nil)
}
