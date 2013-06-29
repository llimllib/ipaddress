ipaddress
=========

A library with some convenient ip address functions missing from the go standard library.

Right now, contains three functions:

### func IpToUint32(ip net.IP) (uint32, error)

Converts a given net.IP into a uint32. Assumes that the IP is an ipv4, and returns an error if it's not.

### func Uint32ToIP(ip uint32) net.IP

Given a uint32, converts it into a net.IP assuming that it's an ipv4 address.

### func BroadcastAddress(n \*net.IPNet) net.IP

Given an IPNet, return the [broadcast address](https://en.wikipedia.org/wiki/Broadcast_address) for that network.
Works for both ipv6 and ipv4.

#Example

```go
package main

import (
    "fmt"
    "net"
    "github.com/llimllib/ipaddress"
)

func main() {
	ip := net.ParseIP("1.2.3.4")

	ipn, err := ipaddress.IpToUint32(ip)
	if err != nil {
		fmt.Printf("Error")
	}

	fmt.Printf("ip %s is %d as an integer\n", ip, ipn)

	ips := ipaddress.Uint32ToIP(ipn).String()

	fmt.Printf("And we used Uint32ToIP to return it to %s\n", ips)

	_, netw, _ := net.ParseCIDR("1.2.3.4/24")
	bcast := ipaddress.BroadcastAddress(netw)

	fmt.Printf("The network broadcast address is %s\n", bcast.String())

	//BroadcastAddress works for ipv6 too
	_, net6, _ := net.ParseCIDR("2001:658:22a:cafe::/64")
	bcast = ipaddress.BroadcastAddress(net6)

	fmt.Printf("The ipv6 network broadcast address is %s\n", bcast.String())
}
```

# TODOs

* Convert to/from integers for ipv6. How to do this when Go doesn't have a uint128?
* Count IPs in an IPMask
