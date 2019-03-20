package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var BUFFERSIZE int
var FILESIZE int

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func createBuffer(buf *[]byte, count int) {
	*buf = make([]byte, count)
	if count == 0 {
		return
	}
	for i := 0; i < count; i++ {
		intByte := byte(random(0, 100))
		if len(*buf) > count {
			return
		}
		*buf = append(*buf, intByte)
	}
}

func Create(dst string, b, f int) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fmt.Error("File %s already exists.", dst)
	}
	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, 0)
	for {
		createBuffer(&buff, b)
		buf = buf[:b]
		if _, err := destination.Write(buf); err != nil {
			return err
		}
		if f < 0 {
			break
		}
		f = f - len(buf)
	}
	return err
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Need BUFFERSIZE FILESIZE!")
		return
	}
	output := "/tmp/randomFile"
	BUFFERSIZE, _ = strconv.Atoi(os.Args[1])
	FILESIZE, _ = strconv.Atoi(os.Args[2])
	err := Create(output, BUFFERSIZE, FILESIZE)
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove(output)
	if err != nil {
		fmt.Println(err)
	}
}
