package main

type Vector struct {
	X, Y float64
}

func (v *Vector) add(v2 Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector) sub(v2 Vector) Vector {
	return Vector{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v *Vector) distance(v2 Vector) float64 {
	dx := v.X - v2.X
	dy := v.Y - v2.Y
	return dx*dx + dy*dy
}
