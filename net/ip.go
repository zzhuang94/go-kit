package net

import (
	"net"
	"net/http"
	"strings"
)

func ClientIp(req *http.Request) string {
	ip := req.Header.Get("X-real-ip")
	if ip == "" || strings.ToUpper(ip) == "UNKNOWN" {
		ip = req.Header.Get("x-forwarded-for")
	}
	if ip == "" || strings.ToUpper(ip) == "UNKNOWN" {
		ip = req.Header.Get("Proxy-Client-IP")
	}
	if ip == "" || strings.ToUpper(ip) == "UNKNOWN" {
		ip = req.Header.Get("WL-Proxy-Client-IP")
	}
	if ip == "" || strings.ToUpper(ip) == "UNKNOWN" {
		addr, _, _ := net.SplitHostPort(req.RemoteAddr)
		ip = net.ParseIP(addr).String()
	}
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	return ip
}

func LocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
