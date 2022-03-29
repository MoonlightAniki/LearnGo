package main

import (
	"LearnGo/basics/basic/generics/set"
	"fmt"
)

func main() {
	//test01()
	test02()
}

func test02() {
	s := set.NewSet(1, 2, 3, 4, 5, 4, 3, 2, 1)
	println(s.String())

	println("Len:", s.Len())
	println("IsEmpty:", s.IsEmpty())
	//s.Add(6.5) // error
	s.Add(6)
	println("s:", s.String())
	s.Remove(1, 2)
	println("s:", s.String())
	println("Contains:", s.Contains(1), s.Contains(6))
	println("ToSlice:", s.ToSlice())
}

func test01() {
	ints := map[string]int64{
		"first":  20,
		"second": 40,
	}
	floats := map[string]float64{
		"first":  1.1,
		"second": 2.33,
	}
	fmt.Printf("No Generics, sumInts=%v, sumFloats=%v\n", sumInts(ints), sumFloats(floats))
	fmt.Printf("Generics, sumIntsOrFloats=%v, sumIntsOrFloats=%v\n", sumIntsOrFloats(ints), sumIntsOrFloats(floats))
	fmt.Printf("Generics, sumNumbers=%v, sumNumbers=%v\n", sumNumbers(ints), sumNumbers(floats))
}

// no generics
func sumInts(nums map[string]int64) int64 {
	var sum int64
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumFloats(nums map[string]float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum
}

// generics
func sumIntsOrFloats[K comparable, V int64 | float64](nums map[K]V) V {
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}

type Number interface {
	int64 | float64
}

func sumNumbers[K comparable, V Number](nums map[K]V) V {
	var sum V
	for _, num := range nums {
		sum += num
	}
	return sum
}
