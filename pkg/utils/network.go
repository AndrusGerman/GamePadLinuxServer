package utils

import "net"

func GetLocalIP() []string {
	listInt, err := net.Interfaces()
	if err != nil {
		return make([]string, 0)
	}
	var listIP = make([]string, 0)
	for _, i2 := range listInt {
		ip := getIPByInterface(i2.Name)
		if ip != "" {
			listIP = append(listIP, ip)
		}
	}
	return listIP
}

func getIPByInterface(interfaceName string) string {
	itf, _ := net.InterfaceByName(interfaceName) //here your interface
	item, _ := itf.Addrs()
	var ip net.IP
	for _, addr := range item {
		switch v := addr.(type) {
		case *net.IPNet:
			if !v.IP.IsLoopback() {
				if v.IP.To4() != nil { //Verify if IP is IPV4
					ip = v.IP
				}
			}
		}
	}
	if ip != nil {
		return ip.String()
	} else {
		return ""
	}
}
