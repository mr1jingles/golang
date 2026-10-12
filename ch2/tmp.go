package main

import (
	"fmt"
	"net"
)

func main() {
	addr := "1.2.3.4/32"

	fmt.Println(net.ParseCIDR(addr))

}
