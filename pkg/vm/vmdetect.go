package vm

import (
	"net"
	"strings"
)

/**
 * @Author Lockly
 * @Description
 * @Date 2023/9/25
 **/

// Modified from https://github.com/ShellCode33/VM-Detection
var blacklistedMacAddressPrefixes = []string{
	"00:1C:42", // Parallels
	"08:00:27", // VirtualBox
	"00:05:69", // |
	"00:0C:29", // | > VMWare
	"00:1C:14", // |
	"00:50:56", // |
	"00:16:E3", // Xen
}

func Check() bool {
	interfaces, err := net.Interfaces()
	if err != nil {
		return false
	}

	for _, iface := range interfaces {
		macAddr := iface.HardwareAddr.String()
		if strings.HasPrefix(iface.Name, "Ethernet") ||
			strings.HasPrefix(iface.Name, "以太网") ||
			strings.HasPrefix(iface.Name, "本地连接") {
			if macAddr != "" {
				for _, prefix := range blacklistedMacAddressPrefixes {
					if strings.HasPrefix(macAddr, prefix) {
						return true
					}
				}
			}
		}
	}

	return false
}
