package main

import (
	"fmt"

	"golang.org/x/sys/windows/registry"
)

func printPrivileges() {
	var privs []string

	if ok, _ := privilegeAlwaysInstallElevated(); ok {
		privs = append(privs, "[!] AlwaysInstallElevated")
	}

	if len(privs) == 0 {
		return
	}
	fmt.Print("## Privileges\n\n")
	for _, priv := range privs {
		fmt.Println("-", priv)
	}
	fmt.Print("\n\n")
}

func privilegeAlwaysInstallElevated() (bool, error) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Policies\Microsoft\Windows\Installer`, registry.QUERY_VALUE)
	if err != nil {
		return false, err
	}
	defer k.Close()

	v, _, err := k.GetIntegerValue("AlwaysInstallElevated")
	if v != 1 {
		if err == registry.ErrNotExist {
			err = nil
		}
		return false, err
	}

	k, err = registry.OpenKey(registry.LOCAL_MACHINE, `Software\Policies\Microsoft\Windows\Installer`, registry.QUERY_VALUE)
	if err != nil {
		return false, err
	}
	defer k.Close()

	v, _, err = k.GetIntegerValue("AlwaysInstallElevated")
	if v != 1 {
		if err == registry.ErrNotExist {
			err = nil
		}
		return false, err
	}
	return true, nil
}
