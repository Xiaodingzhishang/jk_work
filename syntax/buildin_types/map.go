package main

func Map() {
	m1 := map[string]int{
		"key1": 123,
	}
	m1["hello"] = 345
	m2 := make(map[string]int, 12)
	m2["key2"] = 12
	val, ok := m1["大名"]
	if ok {
		println(val)
	}
	val = m1["小米"]
	println("小米对应的值：", val)
	delete(m1, "key1")
}
