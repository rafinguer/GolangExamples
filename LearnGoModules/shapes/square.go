package shapes

type Square struct {
	Side float64
}

func (s *Square) calculateArea() float64 {
	return s.Side * s.Side
}

func (s *Square) calculatePerimeter() float64 {
	return (2 * s.Side) + (2 * s.Side)
}
