package main

import (
	"bytes"
	"fmt"
	"net/netip"
)

type Printer int32

const (
	None Printer = iota
	SamsungSyncThru
)

func checkAndGetValidIPRange(start, end string) (netip.Addr, netip.Addr, error) {
	startAddr, err := netip.ParseAddr(start)
	if err != nil {
		return netip.Addr{}, netip.Addr{}, fmt.Errorf("%v: %v", "Start address", err)
	}
	endAddr, err := netip.ParseAddr(end)
	if err != nil {
		return netip.Addr{}, netip.Addr{}, fmt.Errorf("%v: %v", "End address", err)
	}
	if !(startAddr.Is4() && endAddr.Is4()) {
		return netip.Addr{}, netip.Addr{}, fmt.Errorf("not IPv4 address")
	}
	// TODO: Check zero Address or sth like 0.0.0.0?

	// endAddr should be bigger address than startAddr
	if startAddr.Compare(endAddr) > 0 {
		return netip.Addr{}, netip.Addr{}, fmt.Errorf("start address is bigger than end address")
	}
	return startAddr, endAddr, nil
}

// XXX: Using hacky methods
func isPrinterPage(body []byte) (Printer, bool) {
	if bytes.Contains(body, []byte("RedirectToSWS();")) &&
		bytes.Contains(body, []byte("sws/index.html")) {
		return SamsungSyncThru, true
	}
	return None, false
}

func getDefaultPWAndRequest(url string, kind Printer) (string, string, bool, error) {
	switch kind {
	case None:
		return "", "", false, fmt.Errorf("none kind printer cannot be exist")
	case SamsungSyncThru:
		id, password := "admin", "sec00000"
		url = url + "/sws/app/gnb/login/login.jsp"
		//resp, http.PostForm()
	}
}
