package p3

import "CachePorter/generics/offical/p2"

func CheckIdentity(v interface{}) {
	if vi, ok := v.(p2.HasIdentity2[int]); ok {
		if got := vi.Identity(0); got != 0 {
			panic(got)
		}
	}
}
