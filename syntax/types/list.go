package main

type List interface {
	Add(idx int, val interface{}) error
	Append(val any)
	Delete(index int)
	toSlice() ([]any, error) //私有方法
}
