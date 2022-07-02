package main

import (
	"flag"
)

// TODO: main과 func를 나누는 것에 대한 생각

var startAddrStr, endAddrStr string

func init() {
	flag.StringVar(&startAddrStr, "start", "", "Start IPv4 address. ex) 192.168.0.1")
	flag.StringVar(&endAddrStr, "end", "", "End IPv4 address. ex) 192.168.0.2")
}

func main() {
	// No flag.parse in init function
	// NOTE: https://github.com/golang/go/issues/31859
	flag.Parse()
	err := checkValidIPRange(startAddrStr, endAddrStr)
	if err != nil {
		panic(err)
	}

	// TODO: actual http communication logic
	/*
		var result string[]
		for {
			// Check certain ip addr is
			content, err := connect(ip_addr)
			if err != nil {
				panic(err)
			}
			res, ok, err := getDefaultPasswordAndSend(content)
			if err != nil {
				panic(err)
			}

			result = append(result, res)
		}
	*/
}
