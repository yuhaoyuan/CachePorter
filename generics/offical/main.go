package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// Officials think how generics should be used

/*
Generics adds three new big things to the language:
1-Type parameters for function and types.
2-Defining interface types as sets of types, including types that don’t have methods.
3-Type inference, which permits omitting type arguments in many cases when calling a function.
*/

func Min(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

/*
type Tree[T interface{}] struct {
	left, right *Tree[T]
	value       T
}

// Lookup :Generic types can have methods
func (t *Tree[T]) Lookup(x T) *Tree[T] {
	return nil
}
*/

// 思考: 如何套娃?
type MyConstraints[T any] interface {
	~[]T
}

//type MyConstraints []any

func ForSlice1[T MyConstraints[E], E any](myList T) {
	for _, item := range myList {
		fmt.Print(item)
	}
	fmt.Println(myList)
}

// ForSlice2
/*
func ForSlice[T interface{ ~[]E }, E interface{}] (myList T) {}
func ForSlice[T ~[]E, E interface{}] (myList T){}
func ForSlice[T ~[]E, E any] (myList T) {}
*/

func ForSlice2[T ~[]E, E any](myList T) {
	for _, item := range myList {
		fmt.Print(item)
	}
	fmt.Println()
}

// Scale returns a copy of s with each element multiplied by c.
func Scale[E constraints.Integer](s []E, c E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

// Scale2 returns a copy of s with each element multiplied by c.
func Scale2[S ~[]E, E constraints.Integer](s S, c E) S {
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v * c
	}
	return r
}

type Point []int32

func (p Point) String() string {
	// Details not important.
	return "Point.String"
}

// ScaleAndPrint doubles a Point and prints it.
func ScaleAndPrint(p Point) {
	//r := Scale(p, 2)
	//fmt.Println(r.String()) // DOES NOT COMPILE

	r1 := Scale2[Point, int32](p, 2)
	r2 := Scale2(p, 2) // p:constraint type inference.
	fmt.Println(r1.String())
	fmt.Println(r2.String())
}

type myListType []int

func main() {
	fmt.Println("/* ******************** P1 - Type parameters *********************/")

	/* ********************************************************************************************/

	fmt.Println("/* ******************** P2 - Type sets *********************/")

	myList := []int{1, 2, 3, 4, 5}

	//ForSlice1[[]int](myList)
	ForSlice1[[]int, int](myList)

	ForSlice2[[]int](myList)
	/* ********************************************************************************************/

	ScaleAndPrint([]int32{1, 2})
	/* ********************************************************************************************/

}
