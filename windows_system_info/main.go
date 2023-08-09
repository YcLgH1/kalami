package main

import (
	"fmt"
	"runtime"
	"strings"
	"syscall"
)

var (
	// advapi = syscall.NewLazyDLL("Advapi32.dll")
	kernel = syscall.NewLazyDLL("Kernel32.dll")
)

func main() {
	// GetBiosStat()

	fmt.Printf("系统时区: %s\n", GetTimezone())
	fmt.Printf("系统语言: %s\n", GetSystemLanguage())

	fmt.Printf("开机时长: %s\n", GetStartTime())
	fmt.Printf("当前用户: %s\n", GetUserName())
	fmt.Printf("当前系统: %s\n", runtime.GOOS)
	fmt.Printf("系统版本: %s\n", GetSystemVersion())
	fmt.Printf("Bios:\t %s\n", GetBiosInfo())
	fmt.Printf("主板:\t %s\n", GetMotherboardInfo())

	fmt.Printf("CPU:\t %s\n", GetCpuInfo())
	fmt.Printf("内存:\t %s MB\n", GetMemory())
	fmt.Printf("磁盘:\t %v\n", GetDiskInfo())

	inets := GetInterfaceInfo()
	fmt.Printf("网卡:\t %v\n", inets)

	fmt.Println("网卡信息:")
	for _, inet := range inets {
		if len(inet.Ipv4) > 0 {
			for _, ip := range inet.Ipv4 {
				if strings.Index(ip, "192.") == 0 ||
					strings.Index(ip, "172.") == 0 ||
					strings.Index(ip, "10.") == 0 {
					fmt.Println("==== ", inet)
				}
			}
		}
	}
}
