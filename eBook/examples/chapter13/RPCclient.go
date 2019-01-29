package main

import (
	"fmt"
	"net/rpc"
	"os"
	"sharedRPC"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port sting!")
		return
	}

	CONNECT := arguments[1]
	c, err := rpc.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	args := sharedRPC.MyFloats(16, -0.5)
	var reply float64

	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MyInterface.Power", argus, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
