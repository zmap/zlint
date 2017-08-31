// iputil.go
// contains helper functions for ip address lints

package util

import (
	"bytes"
	"net"
)

type IPRange struct {
	lower net.IP
	upper net.IP
}

var ReservedRanges4 = []IPRange{
	{net.ParseIP("0.0.0.0"), net.ParseIP("0.255.255.255")},
	{net.ParseIP("10.0.0.0"), net.ParseIP("10.255.255.255")},
	{net.ParseIP("100.64.0.0"), net.ParseIP("10.127.255.255")},
	{net.ParseIP("127.0.0.0"), net.ParseIP("127.255.255.255")},
	{net.ParseIP("169.254.0.0"), net.ParseIP("169.254.255.255")},
	{net.ParseIP("172.16.0.0"), net.ParseIP("172.31.255.255")},
	{net.ParseIP("192.0.0.0"), net.ParseIP("192.0.0.255")},
	{net.ParseIP("192.0.2.0"), net.ParseIP("192.0.2.255")},
	{net.ParseIP("192.88.99.0"), net.ParseIP("192.88.99.255")},
	{net.ParseIP("192.168.0.0"), net.ParseIP("192.168.255.255")},
	{net.ParseIP("198.18.0.0"), net.ParseIP("198.19.255.255")},
	{net.ParseIP("198.51.100.0"), net.ParseIP("198.51.100.255")},
	{net.ParseIP("203.0.113.0"), net.ParseIP("203.0.113.255")},
	{net.ParseIP("224.0.0.0"), net.ParseIP("239.255.255.255")},
	{net.ParseIP("240.0.0.0"), net.ParseIP("255.255.255.255")},
}

var ReservedRanges6 = []IPRange{
	{net.ParseIP("::"), net.ParseIP("::")},
	{net.ParseIP("::1"), net.ParseIP("::1")},
	{net.ParseIP("::ffff:0.0.0.0"), net.ParseIP("::ffff:255.255.255.255")},
	{net.ParseIP("100::"), net.ParseIP("100::ffff:ffff:ffff:ffff")},
	{net.ParseIP("64:ff9b::0.0.0.0"), net.ParseIP("64:ff9b::255.255.255.255")},
	{net.ParseIP("2001::"), net.ParseIP("2001::ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("2001:10::"), net.ParseIP("2001:1f:ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("2001:20::"), net.ParseIP("2001:2f:ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("2001:db8::"), net.ParseIP("2001:db8:ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("2002::"), net.ParseIP("2002:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("fc00::"), net.ParseIP("fdff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("fe80::"), net.ParseIP("febf:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
	{net.ParseIP("ff00::"), net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff")},
}

func isInRange(ip net.IP, ipRange IPRange) bool {
	if bytes.Compare(ip, ipRange.lower) >= 0 && bytes.Compare(ip, ipRange.upper) <= 0 {
		return true
	} else {
		return false
	}
}

func IsReservedIP(ip net.IP) bool {
	if ip.To4() != nil {
		// This is to deal with the case where ip is shrunk because it is IPv4
		if len(ip) == 4 {
			ip = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, ip[0], ip[1], ip[2], ip[3]}
		}
		for _, theRange := range ReservedRanges4 {
			if isInRange(ip, theRange) {
				return true
			}
		}
	} else if ip.To16() != nil {
		for _, theRange := range ReservedRanges6 {
			if isInRange(ip, theRange) {
				return true
			}
		}
	}
	return false
}

func ValidIP(ip net.IP) bool {
	return !ip.Equal(net.IPv4zero)
}
