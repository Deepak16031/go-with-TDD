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

	areaTests := []struct {
		Shape        Shapes
		expectedArea float64
	}{
		{Rectangle{
			length:  5.0,
			breadth: 6.0,
		}, 30.0},
		{Circle{Radius: 5.0}, 78.53981633974483},
		{Triangle{5.0, 6.0}, 15.0},
	}

	for _, tt := range areaTests {
		actual := tt.Shape.Area()
		if actual != tt.expectedArea {
			t.Errorf("%#v expected %g actual %g", tt.Shape, tt.expectedArea, actual)
		}
	}

}
