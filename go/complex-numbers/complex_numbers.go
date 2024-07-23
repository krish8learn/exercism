package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	// panic("Please implement the Real method")
	return n.a
}

func (n Number) Imaginary() float64 {
	// panic("Please implement the Imaginary method")
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	// panic("Please implement the Add method")
	return Number{
		n1.a + n2.a,
		n1.b + n2.b,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	// panic("Please implement the Subtract method")
	return Number{
		n1.a - n2.a,
		n1.b - n2.b,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	// panic("Please implement the Multiply method")
	return Number{
		(n1.a * n2.a) - (n1.b * n2.b),
		(n1.b * n2.a) + (n1.a * n2.b),
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		a: n.a * factor,
		b: n.b * factor,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	// panic("Please implement the Divide method")
	divider := (n2.a * n2.a) + (n2.b * n2.b)
	realpart := (n1.a*n2.a + n1.b*n2.b) / divider
	imaginePart := (n1.b*n2.a - n1.a*n2.b) / divider

	return Number{
		realpart,
		imaginePart,
	}
}

func (n Number) Conjugate() Number {
	// panic("Please implement the Conjugate method")
	return Number{
		n.a,
		-n.b,
	}
}

func (n Number) Abs() float64 {
	z := math.Sqrt(n.a*n.a + n.b*n.b)
	return z
}

func (n Number) Exp() Number {
	factor := math.Exp(n.a)
	return Number{
		a: factor * math.Cos(n.b),
		b: factor * math.Sin(n.b),
	}
}
