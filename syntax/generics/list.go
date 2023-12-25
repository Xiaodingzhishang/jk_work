package main

import "errors"

type List[T any] interface {
	Add(Idx int, t T)
	Addend(t T)
}

func UseList() {
	//var l List[int]
	//l.Addend(12)
	var lany List[any]
	lany.Addend(12.3)
	lany.Addend(123)
}

type LinkedList[T any] struct {
	head *node[T]
	t    T
}
type node[T any] struct {
	val T
}

func Max[T Number](vals ...T) (T, error) {
	if len(vals) == 0 {
		var t T
		return t, errors.New("你的下标不对")
	}
	res := vals[0]
	for i := 0; i < len(vals); i++ {
		if res < vals[i] {
			res = vals[i]
		}
	}
	return res, nil

}
func AddSlice[T any](slice []T, idx int, val T) ([]T, error) {
	if idx < 0 || idx > len(slice) {
		return nil, errors.New("下标出错")
	}
	res := make([]T, 0, len(slice)+1)
	for i := 0; i < idx; i++ {
		res = append(res, slice[i])
	}
	slice[idx] = val
	for i := idx; i < len(slice); i++ {
		res = append(res, slice[i])
	}
	return res, nil
}

func main() {
	//UseList()
	println(Sum[int](1, 2, 3))
	println(Sum[Integer](1, 2, 3))
	println(Sum[float64](1.2, 2.6, 3))
	println(Sum[float32](1.2, 2.6, 3))
}
