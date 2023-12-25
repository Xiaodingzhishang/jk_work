package main

import "fmt"

func Defer() {
	defer func() {
		println("第一个 defer")
	}()
	defer func() {
		println("第二个 defer")
	}()
}

func DeferClosure() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}
func DeferClosurev1() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

func DefrtReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DefrtReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a
}
func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			fmt.Printf("i的地址是 %p i值是%d \n", &i, i)
		}()

	}
}
func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			fmt.Printf("val的地址是 %p val值是%d \n", &val, val)
		}(i)

	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			fmt.Printf("j的地址是 %p j值是%d \n", &j, j)
		}()

	}
}

func DeferClosureLoopV4() {
	var j int
	for i := 0; i < 10; i++ {
		j = i
		defer func() {
			fmt.Printf("j的地址是 %p j值是%d \n", &j, j)
		}()

	}
}
