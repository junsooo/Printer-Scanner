package main

import (
	"fmt"
	"net/netip"
)

func checkValidIPRange(start, end string) error {
	startAddr, err := netip.ParseAddr(start)
	if err != nil {
		return fmt.Errorf("%v: %v", "Start address", err)
	}
	endAddr, err := netip.ParseAddr(end)
	if err != nil {
		return fmt.Errorf("%v: %v", "End address", err)
	}
	// endAddr should be bigger address than startAddr
	if startAddr.Compare(endAddr) > 0 {
		return fmt.Errorf("start address is bigger than end address")
	}
	return nil
}
