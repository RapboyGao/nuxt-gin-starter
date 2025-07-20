package utils

import (
	"fmt"
	"net"

	"github.com/fatih/color"
)

// GetIPs returns a slice of IP addresses of the local machine.
// If includeLocalhost is true, it also includes "localhost" in the result.
// It filters out loopback addresses and only includes IPv4 addresses.
// Note: This function does not handle errors from net.InterfaceAddrs().
// It is assumed that the caller will handle any potential errors.
// The function is useful for logging or displaying server addresses.
func GetIPs(includeLocalhost bool) []string {
	results := make([]string, 0)

	addrs, _ := net.InterfaceAddrs()

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				results = append(results, ipnet.IP.String())
			}
		}
	}
	if includeLocalhost {
		results = append(results, "localhost")
	}
	return results
}

// LogServer prints the server address in a formatted way.
// It takes a boolean indicating whether to use HTTPS and an integer for the port.
func LogServer(https bool, port int) {
	fmt.Println("Server available:")
	portPart := fmt.Sprint(port)
	for _, ip := range GetIPs(true) {
		protocol := "http://"
		if https {
			protocol = "https://"
		}
		href := protocol + ip + ":" + portPart
		color.New(color.FgGreen, color.BgBlue).Println(href)
	}
}
