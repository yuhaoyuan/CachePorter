package generics

import (
	"testing"
)

func TestFunction(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []string{"a", "b", "c", "d", "e"}
	PrintSlice(a)
	PrintSlice(b)

	c := AnyList[string]{"一", "二", "三", "四", "五"}
	PrintSlice(c)
}
