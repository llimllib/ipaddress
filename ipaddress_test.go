package ipaddress

import (
	"log"
	"net"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, got interface{}) {
	if got != expected {
		log.Printf("\nExpected %#v\nbut got  %#v", expected, got)
		t.FailNow()
	}
}

func TestIpToUint32(t *testing.T) {
	tests := map[string]uint32{
		"0.0.0.0":   0,
		"0.0.0.255": 255,
		"0.0.1.0":   256,
		"0.0.1.100": 356,
		"0.1.0.0":   65536,
		"0.1.1.1":   65793,
		"1.0.0.0":   16777216,
	}

	for test, expected := range tests {
		i, _ := IpToUint32(net.ParseIP(test))
		assertEqual(t, i, expected)
	}
}

func TestIpToUint32Errors(t *testing.T) {
	_, err := IpToUint32(net.ParseIP("2001:0db8:85a3:0042:1000:8a2e:0370:7334"))
	if err == nil {
		t.FailNow()
	}
}
