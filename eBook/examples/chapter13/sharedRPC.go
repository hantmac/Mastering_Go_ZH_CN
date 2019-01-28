package sharedRPC

type MyFloats struct {
	A1, A2 float64
}

type MyInterface interface {
	Multiply(arguments *MyFloats, reply *float64) error
	Power(arguments *MyFloats, reply *float64) error
}
