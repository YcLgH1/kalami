package main

import (
	"fmt"
	"syscall"
	"time"
)

// 开机时间
func GetStartTime() string {
	GetTickCount := kernel.NewProc("GetTickCount")
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return ""
	}
	ms := time.Duration(r * 1000 * 1000)
	return ms.String()
}

// 时区, 与UTC时间差
func GetTimezone() string {
	name, offset := time.Now().Zone()
	return fmt.Sprintf("Local Zone: %s, offset to UTC(seconds): %v", name, offset)
}

// 当前用户名
func GetUserName() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user, _ := syscall.UTF16PtrFromString("USERNAME")
	domain, _ := syscall.UTF16PtrFromString("USERDOMAIN")
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
		return ""
	}

	buffer[r] = '@'
	old := r + 1
	if old >= size {
		return syscall.UTF16ToString(buffer[:r])
	}

	_, err = syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
	if err != nil {
		return ""
	}

	return syscall.UTF16ToString(buffer[:old+r])
}

// 系统版本
func GetSystemVersion() string {
	version, err := syscall.GetVersion()
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%d.%d (%d)", byte(version), uint8(version>>8), version>>16)
}
