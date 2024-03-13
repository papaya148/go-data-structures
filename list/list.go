package list

type List[T any] interface {
	Length() uint
	HeadValue() T
	TailValue() T
	Append(value T)
	Insert(value T, position uint)
	Delete(position uint) T
	ToString() string
}
