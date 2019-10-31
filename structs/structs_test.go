package structs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	assert.Equal(t, want, got, want)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{6.0, 12.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36},
	}

	for _, tt := range areaTests {
		//checkArea(t, tt.shape, tt.want)
		got := tt.shape.Area()
		assert.Equal(t, got, tt.want)
	}
}
