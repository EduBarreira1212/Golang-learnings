package shapes_test

import (
	. "advanced-tests/shape"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rect := Rectangle{10, 12}
		areaExpected := float64(120)
		areaReturned := rect.Area()

		if areaExpected != areaReturned {
			t.Fatalf("The area returned %f is different than expected %f", areaReturned, areaExpected)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circ := Circle{10}
		areaExpected := float64(math.Pi * 100)
		areaReturned := circ.Area()

		if areaExpected != areaReturned {
			t.Fatalf("The area returned %f is different than expected %f", areaReturned, areaExpected)
		}
	})
}
