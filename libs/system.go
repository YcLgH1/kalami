package libs

import (
	"fmt"
	"os/user"
)

func GetLoginUser() {
	curUser, err := user.Current()

	if err == nil {
		fmt.Println("")
		fmt.Printf("Gid %s\n", curUser.Gid)
		fmt.Printf("Uid %s\n", curUser.Uid)
		fmt.Printf("Username %s\n", curUser.Username)
		fmt.Printf("Name %s\n", curUser.Name)
		fmt.Printf("HomeDir %s\n", curUser.HomeDir)
		fmt.Println("")
	}
}

func GetHomeDir() string {
	curUser, err := user.Current()
	if err != nil {
		return ""
	}

	return curUser.HomeDir
}
