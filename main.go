package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	fHost := flag.String("host", "", "Host/IP to connect to")
	fPort := flag.Uint("port", 443, "Port to connect to")
	fTimeout := flag.Uint("timeout", 5, "Timeout in seconds")
	flag.Parse()
	if *fHost == "" || *fPort >= 0xffff || *fPort == 0 {
		flag.Usage()
		os.Exit(1)
	}
	var timeout uint = 5
	if *fTimeout > 0 {
		timeout = *fTimeout
	}
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", *fHost, *fPort),
		time.Duration(timeout)*time.Second)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	fmt.Println("OK")
}
