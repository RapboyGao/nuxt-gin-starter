package utils

import (
	"fmt"
	"net"
)

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

func LogServer(https bool, port int) {
	fmt.Println("Server available:")
	portPart := fmt.Sprint(port)
	for _, ip := range GetIPs(true) {
		protocol := "http://"
		if https {
			protocol = "https://"
		}
		fmt.Println("\033[36m", protocol+ip+":"+portPart)
	}
}
