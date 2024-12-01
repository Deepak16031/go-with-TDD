package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{5.0, 6.0}
	actual := Perimeter(rectangle)
	expected := 22.0

	if actual != expected {
		t.Errorf("expected %.2f but actual %.2f", expected, actual)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shapes, expected float64) {
		t.Helper()
		actual := shape.Area()
		if actual != expected {
			t.Errorf("expected %g but actual %g", expected, actual)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 6.0}
		checkArea(t, rectangle, 30.0)
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})

}
