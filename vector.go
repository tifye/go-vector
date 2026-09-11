package vector

import (
	"fmt"
	"math"
)

const (
	epsilon float64 = 1e-9
)

type Vector struct {
	X, Y float64
}

func (v Vector) Subtract(o Vector) Vector {
	return Vector{
		X: v.X - o.X,
		Y: v.Y - o.Y,
	}
}

func (p Vector) Add(o Vector) Vector {
	return Vector{
		X: p.X + o.X,
		Y: p.Y + o.Y,
	}
}

func (p Vector) Limit(max float64) Vector {
	magSqrd := p.MagnitudeSquared()

	if magSqrd < max*max {
		return p
	}

	scale := max/math.Sqrt(magSqrd) + epsilon
	return p.MultiplyScalar(scale)
}

func (p Vector) MagnitudeSquared() float64 {
	return p.X*p.X + p.Y*p.Y
}

func (p Vector) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Vector) Distance(o Vector) float64 {
	temp := p.Subtract(o)
	return math.Sqrt(temp.X*temp.X + temp.Y*temp.Y)
}

func (p Vector) DistanceSquared(o Vector) float64 {
	temp := p.Subtract(o)
	return temp.X*temp.X + temp.Y*temp.Y
}

func (p Vector) Normalize() Vector {
	sqr := math.Sqrt(p.X*p.X + p.Y*p.Y)
	if sqr == 0 {
		return Vector{}
	}

	return Vector{
		X: p.X / sqr,
		Y: p.Y / sqr,
	}
}

func (p Vector) MultiplyScalar(v float64) Vector {
	return Vector{
		X: p.X * v,
		Y: p.Y * v,
	}
}

func (p Vector) Dot(o Vector) float64 {
	return p.X*o.X + p.Y*o.Y
}

func (p Vector) Cross(o Vector) float64 {
	return p.X*o.Y - p.Y*o.X
}

func (p Vector) Rotate90CounterClockwise() Vector {
	return Vector{
		X: -p.Y,
		Y: p.X,
	}
}

func (p Vector) Rotate90Clockwise() Vector {
	return Vector{
		X: p.Y,
		Y: -p.X,
	}
}

func (p Vector) Rotate(deltaAngle float64) Vector {
	newAngle := math.Atan2(p.Y, p.X) + deltaAngle
	p.X = math.Cos(newAngle)
	p.Y = math.Sin(newAngle)
	return p
}

func (p Vector) RotateAround(o Vector, deltaAngle float64) Vector {
	var t Vector

	p = p.Subtract(o)
	t.X = p.X*math.Cos(deltaAngle) - p.Y*math.Sin(deltaAngle)
	t.Y = p.X*math.Sin(deltaAngle) + p.Y*math.Cos(deltaAngle)

	return t.Add(o)
}

func (p Vector) Follow(o Vector, distance float64) Vector {
	return p.Subtract(o).
		Normalize().
		MultiplyScalar(distance).
		Add(o)
}

func (p Vector) AngleBetween(o Vector) float64 {
	temp := p.Subtract(o)
	return math.Atan2(temp.Y, temp.X) + math.Pi
}

func (p Vector) Angle() float64 {
	return math.Atan2(p.Y, p.X)
}

func (p Vector) String() string {
	return fmt.Sprintf("(%b, %b)", p.X, p.Y)
}
