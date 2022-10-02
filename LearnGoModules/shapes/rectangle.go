package shapes

type Rectangle struct {
	Base   float64
	Height float64
}

func (r *Rectangle) calculateArea() float64 {
	return r.Base * r.Height
}

func (r *Rectangle) calculatePerimeter() float64 {
	return (2 * r.Base) + (2 * r.Height)
}
