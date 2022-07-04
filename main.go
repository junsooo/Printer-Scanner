package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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
	startAddr, endAddr, err := checkAndGetValidIPRange(startAddrStr, endAddrStr)
	if err != nil {
		panic(err)
	}
	// Implement actual HTTP communication logic
	currentAddr := startAddr
	for currentAddr.Compare(endAddr) < 1 {
		currentUrl := "http://" + currentAddr.String()
		line := []string{currentAddr.String()}
		// NOTE(junsoo): Put address.next before first continue statement,
		// but after recording current address
		// 여기서 currentAddr를 올려줘서 생기는 혼란 vs 각 에러에 currentAddr next를 모두 넣어주는 불편함?
		currentAddr = currentAddr.Next()
		resp, err := http.Get(currentUrl)
		if err != nil {
			log.Printf("http.Get: %v failed, %v", currentUrl, err)
			continue
		}
		// TODO: Check resp.StatusCode?
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("ReadAll: response.body: %v failed", currentUrl)
			continue
		}
		err = resp.Body.Close()
		if err != nil {
			log.Printf("Close: response.body: %v failed", currentUrl)
			continue
		}
		kind, ok := isPrinterPage(bytes)
		if !ok {
			continue
		}
		id, password, success, err := getDefaultPWAndRequest(currentUrl, kind)
		if err != nil {
			log.Printf("getDefaultPWAndCheck: %v failed, kind: %v", currentUrl, kind)
			continue
		}
		if !success {
			line = append(line, "X")
		} else {
			line = append(line, "O", id, password)
		}
		println(strings.Join(line, ","))
	}
}
