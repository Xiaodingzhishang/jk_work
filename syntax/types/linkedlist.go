package main

type node struct {
	prev *node
	next *node
}

type LinkedList struct {
	head *node
	tail *node
	Len  int
}

//
//func (l LinkedList) Add() {
//
//}
//
//// 方法接收器
//func (l *LinkedList) Addv1() {
//
//}

func (l *LinkedList) Add(idx int, val interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) toSlice() ([]any, error) {
	//TODO implement me
	panic("implement me")
}
