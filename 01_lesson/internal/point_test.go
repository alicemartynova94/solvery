package internal_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"solvery/01_lesson/internal"
	"testing"
)

func TestPoint_CalculateDistance(t *testing.T) {
	tests := []struct {
		name        string
		p1, p2      internal.Point
		expectedVal float64
		expectErr   bool
	}{
		{"same location", internal.Point{
			X: 0,
			Y: 0,
		}, internal.Point{
			X: 0,
			Y: 0,
		}, 0, false},
		{"positive", internal.Point{
			X: 2,
			Y: 0,
		}, internal.Point{
			X: 0,
			Y: 0,
		}, 2, false},
		{"negative", internal.Point{
			X: -2,
			Y: -4,
		}, internal.Point{
			X: -6,
			Y: -9,
		}, 6.40312423743284, false},
		{"NaN", internal.Point{
			X: math.NaN(),
			Y: -4,
		}, internal.Point{
			X: -6,
			Y: -9,
		}, 0.0, true},
		{"negative", internal.Point{
			X: math.MaxFloat64,
			Y: -4,
		}, internal.Point{
			X: -6,
			Y: -9,
		}, 0.0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.p1.CalculateDistance(test.p2)
			assert.InDelta(t, test.expectedVal, result, 1e-9)
			assert.Equal(t, test.expectErr, err != nil)
		})
	}
}

func TestPoint_InRadius(t *testing.T) {
	tests := []struct {
		name     string
		p1, p2   internal.Point
		radius   float64
		expected bool
	}{
		{"expected true: in radius", internal.Point{
			X: 0,
			Y: 0,
		}, internal.Point{
			X: 0,
			Y: 0,
		}, 0, true},
		{"expected true: radius=distance", internal.Point{
			X: 0,
			Y: 0,
		}, internal.Point{
			X: 3,
			Y: 4,
		}, 5, true},
		{"expected false: no in radius", internal.Point{
			X: 0,
			Y: 0,
		}, internal.Point{
			X: 1,
			Y: 1,
		}, 0, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := test.p1.InRadius(test.p2, test.radius)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestPoint_InRadius_Err(t *testing.T) {
	radius := 5.0
	p1 := internal.Point{
		X: math.NaN(),
		Y: 0}
	p2 := internal.Point{
		X: 0,
		Y: 0,
	}

	_, err := p1.InRadius(p2, radius)
	assert.Error(t, err)
}

func TestPoint_ParsePoints(t *testing.T) {
	tests := []struct {
		name    string
		point   string
		wantErr bool
	}{
		{"expect success", "2,0", false},
		{"expect parse float err", "s,4", true},
		{"expect parse float err", "4,s", true},
		{"expect err space", "5   ,4", true},
		{"expect slice len err", "4", true},
		{"expect slice len err: too much elements", "4,6,8", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := internal.ParsePoint(test.point)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestPolygon_Perimeter(t *testing.T) {
	tests := []struct {
		name      string
		p         internal.Polygon
		perimeter float64
		wantErr   bool
	}{
		{"expect success: triangle", internal.Polygon{Points: []internal.Point{
			{X: 0, Y: 0},
			{X: 3, Y: 0},
			{X: 3, Y: 4},
		}}, 12.0, false},
		{"expect success: triangle", internal.Polygon{Points: []internal.Point{
			{X: -2, Y: -1},
			{X: 3, Y: -1},
			{X: 3, Y: 2},
		}}, 13.8309518948453, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.p.Perimeter()
			assert.Equal(t, test.perimeter, result)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestPolygon_Area(t *testing.T) {
	tests := []struct {
		name      string
		p         internal.Polygon
		perimeter float64
	}{
		{"expect success: triangle", internal.Polygon{Points: []internal.Point{
			{X: 0, Y: 0},
			{X: 3, Y: 0},
			{X: 3, Y: 4},
		}}, 6.0},
		{"expect success with negative numbers: triangle", internal.Polygon{Points: []internal.Point{
			{X: -2, Y: -1},
			{X: 3, Y: -1},
			{X: 3, Y: 2},
		}}, 7.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.p.Area()
			assert.Equal(t, test.perimeter, result)
		})
	}
}
