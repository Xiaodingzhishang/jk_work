package main

import "fmt"

func Byte() {
	var a byte = 'a'
	fmt.Println(a)
	println(fmt.Sprintf("%c", a))
	var str string = "this is string"
	var bs []byte = []byte(str) //[]byte和string之间可以互相装换
	println(str, bs)
}
