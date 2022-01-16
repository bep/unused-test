package firstpackage

var (
	UsedVar   = "UsedVar"
	UnusedVar = "UnusedVar"
)

const (
	UsedConst   = "UsedConst"
	UnusedConst = "UnusedConst"
)

func UsedFunction() {
	println("UsedFunction")
}

func UnusedFunction() {
	println("UnusedFunction")
}

type MyType struct {
	UsedField   string
	UnusedField string
}

func (t MyType) UsedMethod() {
	println("UsedMethod")
}

func (t MyType) UnusedMethod() {
	println("UnusedMethod")
}

type UnusedInterfaceWithUsedAndUnusedMethod interface {
	UsedInterfaceMethodReturningInt() int
	UnusedInterfaceMethodReturningInt() int
}

type UnusedInterface interface {
	UnusedInterfaceReturningInt() int
}

type UsedInterface interface {
	UsedInterfaceReturningInt() int
}

type UsedInterface2 interface {
	UsedInterface2ReturningInt() int
}
