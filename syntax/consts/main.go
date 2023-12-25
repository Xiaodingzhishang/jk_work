package main

const External = "包外"
const internal = "包内"
const (
	a = 123
)
const (
	StatusA = iota
	StatusB
	StatusC
	StatusD
)
const (
	DayA = iota*12 + 13
	DayB
	DayC
)

func main() {
	const a = 123
}
