package main

import (
	"fmt"
	"proxy"
)

func main() {
	fmt.Println("Start the proxy server")
	proxy.StartProxy()
	fmt.Println("End the proxy server")
}
