package main

import (
	"fmt"
	"kalami/libs"
	"os/exec"
	"strings"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println()
	fmt.Println("Searching wlan password...")
	fmt.Println()

	// 配置列表
	cfgs := GetWifiCfg()

	for i, cfg := range cfgs {
		name := strings.TrimSpace(strings.Split(strings.TrimSpace(cfg), ":")[1])
		fmt.Printf("== %d %s ", i, name)

		// 密码
		password := GetWifiDetail(name)
		if password != "" {
			fmt.Printf("Password: %s\n", password)
		}
	}

	fmt.Println()
}

// 获取wifi配置列表
func GetWifiCfg() []string {
	listCmd := exec.Command("cmd", "/C", "netsh wlan show profile")
	listOut, err := listCmd.Output()
	if err != nil {
		fmt.Println("Searching password failed.")
		return []string{}
	}

	listOutStr := libs.Byte2String(listOut, libs.GB18030)
	// fmt.Println(listOutStr)

	listIdx := strings.Index(listOutStr, "All User Profile")
	if listIdx == -1 {
		listIdx = strings.Index(listOutStr, "所有用户配置文件")
	}
	// TODO support more lang

	cfg := listOutStr[listIdx:]
	cfg = strings.TrimSpace(cfg)

	cfgs := strings.Split(cfg, "\n")
	// fmt.Println(cfgs)

	return cfgs
}

// 获取wifi详细配置
func GetWifiDetail(name string) string {
	searchCmd := exec.Command("cmd", "/C", "netsh wlan show profile name="+name+" key=clear")
	searchOut, err := searchCmd.Output()
	if err != nil {
		fmt.Printf("Search Failed\n")
		return ""
	}

	searchOutStr := libs.Byte2String(searchOut, libs.GB18030)
	// fmt.Println(searchOutStr)

	searchIdx := strings.Index(searchOutStr, "Key Content")
	if searchIdx == -1 {
		searchIdx = strings.Index(searchOutStr, "关键内容")
	}
	// TODO support more lang

	if searchIdx < 0 {
		fmt.Printf("No Password Found\n")
		return ""
	}

	password := strings.Split(strings.Split(searchOutStr[searchIdx:], "\n")[0], ": ")[1]

	return password
}
