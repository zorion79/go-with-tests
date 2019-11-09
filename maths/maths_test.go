package maths

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{150, 150 - 90}
	got := SecondHand(tm)

	assert.Equal(t, want, got)
}
