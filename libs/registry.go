package libs

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func GetUserPath(prefix string) []string {
	keys := []string{}

	fullPrefix, err := registry.OpenKey(registry.USERS, prefix, registry.READ)
	if err != nil {
		return keys
	}
	defer fullPrefix.Close()

	folder, _ := fullPrefix.ReadSubKeyNames(1000)
	if len(folder) == 0 {
		// 没有子文件夹
		return []string{prefix}
	}

	for _, k := range folder {
		subFolder := k
		if prefix != "" {
			subFolder = fmt.Sprintf("%s\\%s", prefix, k)
		}

		subKeys := GetUserPath(subFolder)
		keys = append(keys, subKeys...)
	}

	return keys
}
