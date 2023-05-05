// go:build windows

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var windir = strings.ToLower(os.Getenv("windir"))

func main() {
	fmt.Print("# WINGERPRINT\n\n\n")
	printEnv()
	printUsers()
	printPrivileges()
	printAV()
	printUnattended()
}

func printUnattended() {
	paths := []string{
		filepath.Join(windir, "..\\unattend.inf"),
		filepath.Join(windir, "..\\unattend.txt"),
		filepath.Join(windir, "panther\\unattend.xml"),
		filepath.Join(windir, "panther\\unattend\\unattend.xml"),
		filepath.Join(windir, "panther\\unattend\\unattended.xml"),
		filepath.Join(windir, "panther\\unattended.xml"),
		filepath.Join(windir, "sysprep.inf"),
		filepath.Join(windir, "sysprep\\sysprep.inf"),
		filepath.Join(windir, "sysprep\\sysprep.xml"),
		filepath.Join(windir, "system32\\sysprep\\unattend.xml"),
		filepath.Join(windir, "system32\\sysprep\\unattended.xml"),
	}

	var found []string
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			found = append(found, p)
		}
	}

	if len(found) == 0 {
		return
	}
	fmt.Print("## Unattended files\n\n")
	for _, p := range found {
		fmt.Println("-", p)
	}
	fmt.Print("\n\n")
}
