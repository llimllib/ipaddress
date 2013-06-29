ipaddress
=========

A library with some convenient ip address functions in go

Right now, contains two functions that work:

## func IpToUint32(ip net.IP) (uint32, error)

Converts a given net.IP into a uint32. Assumes that the IP is an ipv4, and returns an error if it's not.

## func IntToIPv4(ip uint32) net.IP

Given a uint32, converts it into a net.IP assuming that it's an ipv4 address.

# TODOs

* Convert to/from integers for ipv6. How to do this when Go doesn't have a uint128?
* Iterate through ranges of IPs
* Count IPs in an IPMask
* indexing of IPMasks?
