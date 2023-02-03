package p2

//type HasIdentity interface {
//	Identity[T any](T) T
//}

type HasIdentity2[T any] interface {
	Identity(T) T
}
