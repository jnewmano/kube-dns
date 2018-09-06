package main

import (
	"context"
	"fmt"
	"net"
)

func main() {
	net.DefaultResolver.Dial = dialer

	_, out, err := net.LookupSRV("dns-tcp", "tcp", "kube-dns.kube-system")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(out)

}

func dialer(ctx context.Context, protocol string, addr string) (net.Conn, error) {

	// uncomment to send request directly to kube-dns
	// and bypass dnsmasq
	// need to set ip to a dns pod id address
	// addr = "172.17.0.2:10053"
	fmt.Println(protocol, addr)

	var d net.Dialer
	return d.DialContext(ctx, protocol, addr)

}
