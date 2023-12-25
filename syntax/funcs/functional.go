package main

func Functional4() string {
	println("hello functional4")
	return "hello"
}
func Functional5(name string) string {
	println("hello functional5" + name)
	return "hello"
}
func UseFunctional4() {
	myFunc := Functional4
	myFunc()
	myFunc5 := Functional5
	myFunc5("heee")
}

func Functional6() {
	fn := func() string {
		return "hello"
	}
	fn()
}
func Functional7() func() string {
	return func() string {
		return "hello worf"
	}
}

func Functional8() {
	fn := func() string {
		return "hello"
	}()
	println(fn)
}
