package internal

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var ErrPoint = errors.New("point must contain x and y")
var ErrNaN = errors.New("NaN cannot be used as a value in this example")
var ErrInf = errors.New("number overflow")

type Point struct {
	X float64
	Y float64
}

type Polygon struct {
	Points []Point
}

func (p1 Point) CalculateDistance(p2 Point) (float64, error) {
	result := math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
	if math.IsNaN(result) {
		return 0.0, ErrNaN
	} else if math.IsInf(result, 1) {
		return 0.0, ErrInf
	}

	return result, nil
}

func (p1 Point) InRadius(p2 Point, r float64) (bool, error) {
	distance, err := p1.CalculateDistance(p2)
	if err != nil {
		return false, err
	}

	return distance <= r, nil
}

func ParsePoint(s string) (Point, error) {
	coordinates := strings.Split(s, ",")
	if len(coordinates) != 2 {
		return Point{}, ErrPoint
	}

	x, err := strconv.ParseFloat(coordinates[0], 64)
	if err != nil {
		return Point{}, fmt.Errorf("invalid x: %w", err)
	}

	y, err := strconv.ParseFloat(coordinates[1], 64)
	if err != nil {
		return Point{}, fmt.Errorf("invalid y: %w", err)
	}

	return Point{X: x, Y: y}, nil
}

func (p Polygon) Perimeter() (float64, error) {
	var result float64
	points := p.Points
	for i := 1; i < len(points); i++ {
		distance, err := points[i].CalculateDistance(points[i-1])
		if err != nil {
			return 0.0, err
		}
		result += distance
	}
	distance, err := points[0].CalculateDistance(points[len(points)-1])
	if err != nil {
		return 0.0, err
	}
	result += distance

	return result, nil
}

func (p Polygon) Area() float64 {
	points := p.Points
	var result float64
	var sum float64
	for i := 1; i < len(points); i++ {
		p1 := points[i-1]
		p2 := points[i]
		calculations := p1.X*p2.Y - p1.Y*p2.X
		sum += calculations
	}
	p1 := points[len(points)-1]
	p2 := points[0]
	calculations := p1.X*p2.Y - p1.Y*p2.X
	sum += calculations
	result = math.Abs(sum) / 2

	return result
}
