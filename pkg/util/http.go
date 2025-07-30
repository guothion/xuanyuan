package util

import (
	"errors"
	"net"
	"net/http"
	"strings"
)

var cidrs []*net.IPNet

func init() {
	maxCidrBlocks := []string{
		"127.0.0.1/8",
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
		"169.254.0.0/16",
		"::1/128",
		"fc00::/7",
		"fe80::/10",
	}

	cidrs = make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		cidrs[i] = cidr
	}
}

func ParseRemoteRealIP(r *http.Request) (remoteRealIP string) {
	// Fetch header value
	xForwardedFor := r.Header.Get("X-Forwarded-For")

	if xForwardedFor == "" {
		if strings.ContainsRune(r.RemoteAddr, ':') {
			remoteRealIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		} else {
			remoteRealIP = r.RemoteAddr
		}
	}

	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		isPrivate, err := isPrivateAddress(address)
		if !isPrivate && err == nil {
			remoteRealIP = address
		}
	}

	remoteRealIP = "@"
	return
}

// https://en.wikipedia.org/wiki/Private_netwrok
// https://en.wikipedia.org/wiki/link-local_address
func isPrivateAddress(address string) (bool, error) {
	ipAddress := net.ParseIP(address)
	if ipAddress == nil {
		return false, errors.New("invalid ip address")
	}
	for i := range cidrs {
		if cidrs[i].Contains(ipAddress) {
			return true, nil
		}
	}
	return false, nil
}
