package component

type Inner struct {
}

func (i Inner) DoSomething() {

}

type Outer struct {
	Inner
}
type OuterPtr struct {
	*Inner
}
type OOOOuter struct {
	Outer
}

func UserInner() {
	var o Outer
	o.DoSomething()
	var op *OuterPtr
	op.DoSomething()
}
