package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"solvery/lesson_one/internal"
)

func main() {
	var points []string
	var distance bool
	var perimeter bool
	var radius float64
	var area bool

	pflag.StringArrayVar(&points, "point", []string{}, "list of 2 points")
	pflag.BoolVar(&distance, "distance", false, "distance between 1 and 2 point")
	pflag.BoolVar(&perimeter, "perimeter", false, "polygon perimeter")
	pflag.Float64Var(&radius, "radius", -1, "in radius")
	pflag.BoolVar(&area, "area", false, "polygon area")
	pflag.Parse()

	if len(points) >= 3 {
		if perimeter {
			perimeterCheck(points)
		} else if area {
			areaCheck(points)
		}
	} else if len(points) == 2 {
		if radius > 0 {
			radiusCheck(points, radius)
		} else if distance {
			distanceCheck(points)
		} else {
			fmt.Println("unknown format of the command")
		}
	} else {
		fmt.Println("unknown format of the command")
	}
}

func radiusCheck(points []string, radius float64) {
	p1, err := internal.ParsePoint(points[0])
	if err != nil {
		fmt.Printf("parsing err of p1 %v", err)
		return
	}
	p2, err := internal.ParsePoint(points[1])
	if err != nil {
		fmt.Printf("parsing err of p2 %v", err)
		return
	}

	result, err := p1.InRadius(p2, radius)
	if err != nil {
		fmt.Printf("in radius err %v", err)
		return
	}
	fmt.Println(result)
}

func distanceCheck(points []string) {
	p1, err := internal.ParsePoint(points[0])
	if err != nil {
		fmt.Printf("parsing err of p1 %v", err)
		return
	}
	p2, err := internal.ParsePoint(points[1])
	if err != nil {
		fmt.Printf("parsing err of p2 %v", err)
		return
	}
	result, err := p1.CalculateDistance(p2)
	if err != nil {
		fmt.Printf("calculate distance err %v", err)
		return
	}
	fmt.Println(result)
}

func perimeterCheck(points []string) {
	polygon := internal.Polygon{}.Points
	for i, point := range points {
		p, err := internal.ParsePoint(point)
		if err != nil {
			fmt.Printf("point %d parsing err %v", i, err)
			return
		}
		polygon = append(polygon, p)
	}

	result, err := internal.Polygon{Points: polygon}.Perimeter()
	if err != nil {
		fmt.Printf("perimeter err %v", err)
		return
	}

	fmt.Println(result)
}

func areaCheck(points []string) {
	polygon := internal.Polygon{}.Points
	for i, point := range points {
		p, err := internal.ParsePoint(point)
		if err != nil {
			fmt.Printf("point %d parsing err %v", i, err)
			return
		}
		polygon = append(polygon, p)
	}

	result := internal.Polygon{Points: polygon}.Area()

	fmt.Println(result)
}
