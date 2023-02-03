package generics

import "fmt"

func PrintSlice[T any](s []T) {
	for _, item := range s {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
