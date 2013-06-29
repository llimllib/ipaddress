package ipaddress

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

// Convert an ipv4 to a uint32
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

//convert a uint32 to an ipv4
func Uint32ToIP(ip uint32) net.IP {
	return net.IPv4(byte(ip>>24), byte(ip>>16&0xFF), byte(ip>>8&0xFF), byte(ip&0xFF))
}

// Return the broadcast address of a net range. Convert to IPv4 if possible,
// otherwise return an ipv6
func BroadcastAddress(n *net.IPNet) net.IP {
	ip := n.IP.To4()
	if ip == nil {
		ip = n.IP
		return net.IP{
			ip[0] | ^n.Mask[0], ip[1] | ^n.Mask[1], ip[2] | ^n.Mask[2],
			ip[3] | ^n.Mask[3], ip[4] | ^n.Mask[4], ip[5] | ^n.Mask[5],
			ip[6] | ^n.Mask[6], ip[7] | ^n.Mask[7], ip[8] | ^n.Mask[8],
			ip[9] | ^n.Mask[9], ip[10] | ^n.Mask[10], ip[11] | ^n.Mask[11],
			ip[12] | ^n.Mask[12], ip[13] | ^n.Mask[13], ip[14] | ^n.Mask[14],
			ip[15] | ^n.Mask[15]}
	}

	return net.IPv4(
		ip[0]|^n.Mask[0],
		ip[1]|^n.Mask[1],
		ip[2]|^n.Mask[2],
		ip[3]|^n.Mask[3])
}
