package ipaddress

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

// Assume it's an IPv4
func IpToUint32(ip net.IP) (uint32, error) {
	var n uint32
	ip = ip.To4()
	if ip == nil {
		return 0, errors.New("not able to convert ip to ipv4")
	}
	buf := bytes.NewBuffer([]byte(ip))
	binary.Read(buf, binary.BigEndian, &n)
	return n, nil
}

func IntToIPv4(ip uint32) net.IP {
	return net.IPv4(byte(ip>>24), byte(ip>>16&0xFF), byte(ip>>8&0xFF), byte(ip&0xFF))
}
