package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountdown(t *testing.T) {
	buffer := new(bytes.Buffer)

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
GO!`

	assert.Equal(t, want, got)
}
