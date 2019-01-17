package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: use selectColumn column [file1] [file2] [...]\n")
		os.Exit(1)
	}

	temp , err := strconv.Atoi(arguments[i])
	if err != nil  {
		fmt.Printf("column value is not an integer", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Invalid column value")
		os.Exit(1)
	}

	for _, fileName := range arguments[2:]{
		fmt.Println("\t\t",fileName)
		f, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("error opening file %s\n",err)
			continue
		}
		f.Close()

		r := bufio.NewReader(f)
		for {
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil{
				fmt.Printf("error reading file %s\n",err)
			}

			data := strings.Fields(line)
			if len(data) > column {
				fmt.Println(data[column-1])
			}
		}

	}

}