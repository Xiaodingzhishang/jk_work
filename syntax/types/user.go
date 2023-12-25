package main

import "fmt"

func NewUser() {
	u := User{}
	fmt.Printf("%v \n", u)
	fmt.Printf("%+v \n", u)
	//是一个指针
	up := &User{}
	fmt.Printf("up %+v \n", up)
	up2 := new(User)
	fmt.Printf("up2 %+v \n", up2)
	u4 := User{
		Name: "Tom",
		Age:  14,
	}
	u5 := User{"Tom", 14}
	u4.Name = "Jerry"
	u5.Age = 18
	var up3 *User
	println(up3)

}

type User struct {
	Name string
	Age  int
}

func (u User) ChangName(name string) {
	fmt.Printf("change name 中U 的地址%p \n", &u)
	u.Name = name

}

//	func ChangName(u User, name string) {
//		fmt.Printf("change name 中U 的地址%p \n", &u)
//		u.Name = name
//
// }
func (u *User) ChangAge(age int) {
	fmt.Printf("change Age 中U 的地址%p \n", u)
	u.Age = age
}

//func ChangAge(u *User, age int) {
//	fmt.Printf("change Age 中U 的地址%p \n", u)
//	u.Age = age
//}

func ChangeUser() {
	u1 := User{
		Name: "Tom",
		Age:  18,
	}
	fmt.Printf("U1 的地址%p \n", &u1)
	u1.ChangAge(35)
	u1.ChangName("Jerry")
	fmt.Printf("%+v \n", u1)
	up1 := &User{}
	up1.ChangName("Jerry")
	up1.ChangAge(35)
	fmt.Printf("%+v \n", up1)
}

type Integer int

func UseInt() {
	i1 := 10
	i2 := Integer(i1)
	var i3 Integer = 11
	println(i2, i3)
}

type Fish struct {
	Name string
}

func (f Fish) Swim() {
	println("fist在游")
}

type FakeFish Fish

func UseFish() {
	f1 := Fish{}
	f2 := FakeFish(f1)
	println(f2)
}
