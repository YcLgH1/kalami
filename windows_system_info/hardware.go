package main

import (
	"fmt"
	"net"
	"strings"
	"syscall"
	"unsafe"

	"github.com/StackExchange/wmi"
)

type diskusage struct {
	Path  string `json:"path"`
	Total uint64 `json:"total"`
	Free  uint64 `json:"free"`
}

func usage(getDiskFreeSpaceExW *syscall.LazyProc, path string) (diskusage, error) {
	lpFreeBytesAvailable := int64(0)
	var info = diskusage{Path: path}
	_path, _ := syscall.UTF16PtrFromString(info.Path)
	diskRet, _, err := getDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(_path)),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&(info.Total))),
		uintptr(unsafe.Pointer(&(info.Free))))
	if diskRet != 0 {
		err = nil
	}
	return info, err
}

// 硬盘信息
func GetDiskInfo() (infos []diskusage) {
	GetLogicalDriveStringsW := kernel.NewProc("GetLogicalDriveStringsW")
	GetDiskFreeSpaceExW := kernel.NewProc("GetDiskFreeSpaceExW")
	lpBuffer := make([]byte, 254)
	diskRet, _, _ := GetLogicalDriveStringsW.Call(
		uintptr(len(lpBuffer)),
		uintptr(unsafe.Pointer(&lpBuffer[0])))
	if diskRet == 0 {
		return
	}
	for _, v := range lpBuffer {
		if v >= 65 && v <= 90 {
			path := string(v) + ":"
			if path == "A:" || path == "B:" {
				continue
			}
			info, err := usage(GetDiskFreeSpaceExW, string(v)+":")
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
	}
	return infos
}

// CPU信息
// 简单的获取方法fmt.Sprintf("Num:%d Arch:%s\n", runtime.NumCPU(), runtime.GOARCH)
func GetCpuInfo() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	_num, _ := syscall.UTF16FromString("Num:")
	var index = uint32(copy(buffer, _num) - 1)
	nums, _ := syscall.UTF16PtrFromString("NUMBER_OF_PROCESSORS")
	arch, _ := syscall.UTF16PtrFromString("PROCESSOR_ARCHITECTURE")
	r, err := syscall.GetEnvironmentVariable(nums, &buffer[index], size-index)
	if err != nil {
		return ""
	}
	index += r
	_arch, _ := syscall.UTF16FromString(" Arch:")
	index += uint32(copy(buffer[index:], _arch) - 1)
	r, err = syscall.GetEnvironmentVariable(arch, &buffer[index], size-index)
	if err != nil {
		return syscall.UTF16ToString(buffer[:index])
	}
	index += r
	return syscall.UTF16ToString(buffer[:index+r])
}

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

// 内存信息
func GetMemory() string {
	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return ""
	}
	return fmt.Sprint(memInfo.ullTotalPhys / (1024 * 1024))
}

type InterfaceInfo struct {
	Name string
	Ipv4 []string
	Ipv6 []string
}

// 网卡信息
func GetInterfaceInfo() []InterfaceInfo {
	interfaces, err := net.Interfaces()
	if err != nil {
		return []InterfaceInfo{}
	}
	var is = make([]InterfaceInfo, len(interfaces))
	for i, v := range interfaces {
		ips, err := v.Addrs()
		if err != nil {
			continue
		}
		is[i].Name = v.Name
		for _, ip := range ips {
			if strings.Contains(ip.String(), ":") {
				is[i].Ipv6 = append(is[i].Ipv6, ip.String())
			} else {
				is[i].Ipv4 = append(is[i].Ipv4, ip.String())
			}
		}
	}
	return is
}

// 主板信息
func GetMotherboardInfo() string {
	var s = []struct {
		Product string
	}{}
	err := wmi.Query("SELECT Product FROM Win32_BaseBoard WHERE (Product IS NOT NULL)", &s)
	if err != nil {
		return ""
	}
	return s[0].Product
}

// BIOS信息
func GetBiosInfo() string {
	var s = []struct {
		Name string
	}{}
	err := wmi.Query("SELECT Name FROM Win32_BIOS WHERE (Name IS NOT NULL)", &s) // WHERE (BIOSVersion IS NOT NULL)
	if err != nil {
		return ""
	}
	return s[0].Name
}
