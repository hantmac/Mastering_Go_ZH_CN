package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myElement struct {
	Name string
	SurName string
	Id string
}

var DATA = make(map[string]myElement)

func ADD(k string,n myElement) bool {
	if k == "" {
		return false
	}
	if LOOKUP(k) == nil {
		DATA[k] = n
		return true
	}
	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}

func LOOKUP(k string) *myElement  {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	} else {
		return nil
	}
}

func CHANGE(k string, n myElement) bool {
	DATA[k] = n
	return true
}

func PRINT()  {
	for k, v := range DATA {
		fmt.Printf("key: %s value: %v",k,v)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimSpace(text)
		tokens := strings.Fields(text)

		switch len(tokens) {
		case 0:
			continue
		case 1:
			tokens = append(tokens,"")
			tokens = append(tokens,"")
			tokens = append(tokens,"")
			tokens = append(tokens,"")
		case 2:
			tokens = append(tokens,"")
			tokens = append(tokens,"")
			tokens = append(tokens,"")
		case 3:
			tokens = append(tokens,"")
			tokens = append(tokens,"")
		case 4:
			tokens = append(tokens,"")
		
		}

		switch tokens[0] {
		case "PRINT":
			PRINT()
		case "STOP":
			return
		case "DELETE":
			if !DELETE(tokens[1]) {
				fmt.Println("Delete operations failed")
			}
		case "ADD":
			n := myElement{tokens[2],tokens[3],tokens[4]}
			if !ADD(tokens[1],n) {
				fmt.Println("Add operation failed")
			}
		case "LOOKUP":
			n := LOOKUP(tokens[1])
			if n != nil {
				fmt.Printf("%v\n",n)
			}

		case "CHANGE":
			n := myElement{tokens[2],tokens[3],tokens[4]}
			if !CHANGE(tokens[1],n) {
				fmt.Println("Update operation failed")
			}

		default:
			fmt.Println("Unknown command - please try again!")

		}
	}
}