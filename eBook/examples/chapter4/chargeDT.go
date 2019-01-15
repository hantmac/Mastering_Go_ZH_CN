package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) ==1 {
		fmt.Println("Please provide one text file to process")
		os.Exit(1)
	}
	fileName := arguments[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("error opening file %s ",err)
		os.Exit(1)
	}
	defer f.Close()

	notAmatch := 0
	r := bufio.NewReader(f)
	for{
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil{
			fmt.Printf("error reading file %s,",err)

		}

		r1 := regexp.MustCompile(`.*\[(\d\d\/\w+/\d\d\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r1.MatchString(line) {
			match := r1.FindStringSubmatch(line)
			d1, err := time.Parse("02/Jan/2006:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Printf(strings.Replace(line,match[1],newFormat,1))
			} else {
				notAmatch++
				continue
			}

		}

		r2 := regexp.MustCompile(`.*\[(\w+\-\d\d-\d\d:\d\d:\d\d:\d\d.*)\] .*`)
		if r2.MatchString(line) {
			match := r2.FindStringSubmatch(line)
			d1, err := time.Parse("Jan-02-06:15:04:05 -0700", match[1])
			if err == nil {
				newFormat := d1.Format(time.Stamp)
				fmt.Print(strings.Replace(line, match[1], newFormat, 1))
			} else {
				notAmatch++
			}
			continue
		}
	}
	fmt.Println(notAmatch, "lines did not match!")
}