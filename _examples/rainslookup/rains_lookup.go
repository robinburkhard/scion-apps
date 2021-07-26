package main

import (
	"fmt"
	"github.com/netsec-ethz/scion-apps/pkg/appnet"
)

func main() {
	test_name := "ch."
	test_port := ":80"
	raddr, err := appnet.ResolveUDPAddr(test_name+test_port)
	if err != nil {
		fmt.Println("Error resolving name: ", test_name, ":" ,err)
		return
	}
	fmt.Println("Address returned: %q", raddr.String())
}
