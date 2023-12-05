package main

import (
	"unicode/utf8"
)

func String() {
	println(len("wer"))
	println(utf8.RuneCountInString("你好呀"))
	println(utf8.RuneCountInString("wwww"))
}
