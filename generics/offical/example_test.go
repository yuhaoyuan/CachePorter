package main

import (
	"CachePorter/generics/offical/p1"
	"CachePorter/generics/offical/p3"
	"fmt"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func ThePrint[T any](s []T) { // Just an example, not the suggested syntax.
	for _, v := range s {
		fmt.Println(v)
	}
}
func TestTypeParams(t *testing.T) {
	ThePrint[int]([]int{1, 2, 3})
	ThePrint([]int{1, 2, 3})
}

type Stringer interface {
	String() string
}

type Plusser interface {
	Plus(string) string
}

func ConcatTo[S Stringer, P Plusser](s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}
	return r
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func (m MyString) Plus(s string) string {
	return string(m) + s
}

func TestMultipleTypeParameters(t *testing.T) {
	a := []MyString{"This is MyString a."}
	b := []MyString{"This is MyString b."}
	c := ConcatTo[MyString, MyString](a, b)
	fmt.Println(c)
}

type Vector[T any] []T // just like a function’s type parameters.

func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}

func TestGenericTypes(t *testing.T) {
	var v Vector[int]
	v.Push(1)
	fmt.Println(v)
}

// List is a linked list of values of type T.
type List[T any] struct {
	Next *List[T] // this reference to List[T] is OK
	Val  T
}

// This type is INVALID.
type InvalidP[T1, T2 any] struct {
	F *InvalidP[T2, T1] // INVALID; must be [T1, T2]
}

type ListHead[T any] struct {
	head *ListElement[T]
}

type ListElement[T any] struct {
	next *ListElement[T]
	val  T
	// Using ListHead[T] here is OK.
	// ListHead[T] refers to ListElement[T] refers to ListHead[T].
	// Using ListHead[int] would not be OK, as ListHead[T]
	// would have an indirect reference to ListHead[int].
	head *ListHead[int]
}

func TestGenericTypeReferToItself(t *testing.T) {
	a := List[int]{
		Next: &List[int]{nil, 2},
		Val:  1,
	}
	fmt.Println(a)

	b := InvalidP[int8, int32]{
		F: &InvalidP[int32, int8]{},
	}
	fmt.Println(b.F.F)

	c := ListHead[string]{}
	d := ListElement[string]{}
	fmt.Println(c, d)
}

func Double[E constraints.Integer](l []E) []E {
	r := make([]E, len(l))
	for i, v := range l {
		r[i] = v + v
	}
	return r
}

func DoubleDefined[S ~[]E, E constraints.Integer](l S) S {
	r := make(S, len(l))
	for i, v := range l {
		r[i] = v + v
	}
	return r
}

type DoubleDefinedConstraints[T any] interface {
	~[]T
}

func DoubleDefined2[S DoubleDefinedConstraints[E], E constraints.Integer](l S) S {
	r := make(S, len(l))
	for i, v := range l {
		r[i] = v + v
	}
	return r
}

type MyList []int

func TestElementConstraintExample(t *testing.T) {
	listA := Double(MyList{1, 2, 3})
	fmt.Println(reflect.TypeOf(listA)) // the result will not be that defined type.

	listB := DoubleDefined(MyList{1, 2, 3})
	fmt.Println(reflect.TypeOf(listB))

	listC := DoubleDefined2(MyList{1, 2, 3})
	fmt.Println(reflect.TypeOf(listC))
}

type Setter interface {
	Set(string)
}

func FromStrings[T Setter](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i].Set(v)
	}
	return result
}

type Settable int

// Set sets the value of *p from a string.
func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s) // real code should not ignore the error
	*p = Settable(i)
}

// Setter2 is a type constraint that requires that the type
// 1 implement a Set method that sets the value from a string,
// 2 and also requires that the type be a pointer to its type parameter.
type Setter2[T any] interface {
	Set(string)
	*T
}

// We use two different type parameters so that we can return
// a slice of type T but call methods on *T aka PT.
// The Setter2 constraint ensures that PT is a pointer to T.
func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set of Setter2, so we can convert it to PT.
		p := PT(&result[i])
		// PT has a Set method.
		p.Set(v)
	}
	return result
}

func TestPointerMethodExample(t *testing.T) {
	//nums := FromStrings[Settable]([]string{"1", "2"})
	//fmt.Println(nums)

	/*
		INVALID
		Here we want nums to be []Settable{1, 2}.
		Settable does not have a Set method. The type that has a Set method is *Settable.
	*/

	//nums := FromStrings[*Settable]([]string{"1", "2"})
	//fmt.Println(nums)

	/*
			will panic
			debug on the line 68, p is nil

		When instantiated with *Settable, that means a slice of type []*Settable.
		When FromStrings calls result[i].Set(v), that invokes the Set method on the pointer stored in result[i].
		That pointer is nil.
		The Settable.Set method will be invoked with a nil receiver, and will raise a panic due to a nil dereference error.
	*/

	/*
			we can‘t use Settable because it doesn’t have a Set method,
		    and we can‘t use *Settable because then we can’t create a slice of type Settable.
	*/
	//nums := FromStrings2([]string{"1", "2"})
	nums := FromStrings2[Settable, *Settable]([]string{"1", "2"})
	fmt.Println(nums)

	// 类型推断
	nums = FromStrings2[Settable]([]string{"1", "2"})
	fmt.Println(nums)
}

type Equaler[T any] interface {
	Equal(T) bool
}

// also you can write like this
// func Index[T Equaler[E], E any](s []E, e T) int {
func Index[T interface{ Equal(T) bool }](s []T, e T) int {
	for i, v := range s {
		if e.Equal(v) {
			return i
		}
	}
	return -1
}

type equalInt int

func (a equalInt) Equal(b equalInt) bool {
	return a == b
}

func TestUsingItselfTypesInConstraintsExample(t *testing.T) {
	s := []equalInt{1, 2, 3, 4}
	e := equalInt(3)
	fmt.Println(Index[equalInt](s, e))
}

// TP=type parameters
func TestValuesOfTPNotBoxed(t *testing.T) {
	// 什么叫装箱Boxed, 可以理解为内存逃逸
	// interface 的实现是一个指针形式，它保存的是指针，当把一个非指针的值放入interface中，则会逃逸到堆栈上。

	nums := FromStrings2[Settable]([]string{"1", "2"})
	fmt.Println(nums)

	// 但是当我们用Settable实例化 FromStrings2 时，得到的nums并没有被装箱。
}

/* ------------------- Map/Reduce/Filter ------------------- */

// Map turns a []T1 to a []T2 using a mapping function.
// This function has two type parameters, T1 and T2.
// This works with slices of any type.
func SlicesMap[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func SlicesReduce[T1, T2 any](s []T1, initializer T2, f func(T2, T1) T2) T2 {
	r := initializer
	for _, v := range s {
		r = f(r, v)
	}
	return r
}

// Filter filters values from a slice using a filter function.
// It returns a new slice with only the elements of s
// for which f returned true.
func SlicesFilter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
func TestMapReduceFilter(t *testing.T) {
	s := []int{1, 2, 3}

	floats := SlicesMap(s, func(i int) float64 { return float64(i) })
	// Now floats is []float64{1.0, 2.0, 3.0}.

	sum := SlicesReduce(s, 0, func(i, j int) int { return i + j })
	// Now sum is 6.

	evens := SlicesFilter(s, func(i int) bool { return i%2 == 0 })
	// Now evens is []int{2}.
	fmt.Println(floats, sum, evens)
}

/* ------------------- Map/Reduce/Filter ------------------- */

// Keys returns the keys of the map m in a slice.
// The keys will be returned in an unpredictable order.
// This function has two type parameters, K and V.
// Map keys must be comparable, so key has the predeclared
// constraint comparable. Map values can be any type.
func Keys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func TestMapKeys(t *testing.T) {
	k := maps.Keys(map[int]int{1: 2, 2: 4})
	// Now k is either []int{1, 2} or []int{2, 1}.
	fmt.Println(k)
}

//type Ordered interface {
//	~int | ~int8 | ~int16 | ~int32 | ~int64 |
//		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
//		~float32 | ~float64 |
//		~string
//}

// orderedSlice is an internal type that implements sort.Interface.
// The Less method uses the < operator. The Ordered type constraint
// ensures that T has a < operator.
type orderedSlice[T constraints.Ordered] []T

func (s orderedSlice[T]) Len() int           { return len(s) }
func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s orderedSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func OrderedSlice[T constraints.Ordered](s []T) {
	// Convert s to the type orderedSlice[T].
	// As s is []T, and orderedSlice[T] is defined as []T,
	// this conversion is permitted.
	// orderedSlice[T] implements sort.Interface,
	// so can pass the result to sort.Sort.
	// The elements will be sorted using the < operator.
	sort.Sort(orderedSlice[T](s))
}

func TestSort(t *testing.T) {
	s1 := []int32{3, 5, 2}
	OrderedSlice(s1)
	// Now s1 is []int32{2, 3, 5}

	s2 := []string{"a", "c", "b"}
	OrderedSlice(s2)
	// Now s2 is []string{"a", "b", "c"}
}

func TestNoParameterizedMethods(t *testing.T) {
	p3.CheckIdentity(p1.S2[int]{})
}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

//func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
//	r := make(S, len(s))
//	for i, v := range s {
//		r[i] = v * c
//	}
//	return r
//}
//
//type Point []int32
//
//func (p Point) String() string {
//	// Details not important.
//	return fmt.Sprintf("point string %d %d %d", p[0], p[1], p[2])
//}
//
//// ScaleAndPrint doubles a Point and prints it.
//func ScaleAndPrint(p Point) {
//	r := Scale(p, 2)
//	fmt.Println(r.String()) // DOES NOT COMPILE
//}

func TestTypeInference(t *testing.T) {
	// Function argument type inference
	fmt.Println(GMin[float64](1.0, 2.0))
	fmt.Println(GMin(3.0, 4.0))

	// Constraint type inference
	p := []int32{1, 2, 3}
	ScaleAndPrint(p)

	r1 := Scale2[Point, int32](p, 2)
	fmt.Println(r1.String())

	//r2 := Scale2(p, 2)
	//fmt.Println(r2.String())

	p2 := Point(p)
	r3 := Scale2(p2, 2)
	fmt.Println(r3.String())
}

func TestOperation(t *testing.T) {
	//constraints.Ordered
}
