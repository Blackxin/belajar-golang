package main

import "fmt"

// types
type Number interface {
	int64 | float64
}

type UserModel[T int | float64] struct {
	Name   string
	Scores []T
}

func (m *UserModel[int]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

// main function

func main() {
	/* total1 := sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total:", total1)

	total2 := sum([]float32{2.5, 7.2})
	fmt.Println("total:", total2)

	total3 := sum([]float64{1.23, 6.33, 12.6})
	fmt.Println("total:", total3)

	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic sums with constraint: %v and %v\n", sumNumberS1(ints), sumNumberS2(floats)) */

  var m1 UserModel[int]
  m1.Name = "Taufik"
  m1.Scores = []int{1, 2, 3}
  fmt.Println("scores:", m1.Scores)


  var m2 UserModel[float64]
  m2.Name = "Fauzi"
  m2.Scores = []float64{10, 11}
  fmt.Println("scores:", m2.Scores)
}

// functions

func sum[T int | float32 | float64](numbers []T) T {
	var total T
	for _, e := range numbers {
		total += e
	}
	return total
}

func sumNumberS1(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

func sumNumberS2[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func sumNumberS3[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
