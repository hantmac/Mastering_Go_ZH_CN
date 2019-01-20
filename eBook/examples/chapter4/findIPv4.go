package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIp(input string) string {
	partIp := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammer := partIp+"\\."+partIp+"\\."+partIp+"\\."+partIp
	matchMe := regexp.MustCompile(grammer)
	return matchMe.FindString(input)
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Printf("usage: %s logfile\n",filepath.Base(arguments[0]))
		os.Exit(1)
	}
	for _, filename := range arguments[1:] {
		f,err := os.Open(filename)
		if err != nil {
			fmt.Printf("error openning file %s\n",err)
			os.Exit(-1)
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			line,err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else  if err != nil {
				fmt.Printf("error openning file %s\n",err)
				break
			}

			ip := findIp(line)
			trail := net.ParseIP(ip)
			if trail.To4() == nil {
				continue
			} else {
				fmt.Println(ip)
			}
		}
	}
}