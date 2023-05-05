package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func printEnv() {
	var envInfo []string

	cmd := exec.Command("powershell", "-NoProfile", "-Command", "Get-WmiObject -Class Win32_ComputerSystem | Select-Object -ExpandProperty Model")
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Fatalf("Error: running PowerShell command: %v", err)
	}

	model := strings.TrimSpace(out.String())
	modelLC := strings.ToLower(model)
	if strings.Contains(modelLC, "hyper-v") ||
		strings.Contains(modelLC, "hyperv") ||
		strings.Contains(modelLC, "kvm") ||
		strings.Contains(modelLC, "parallels") ||
		strings.Contains(modelLC, "qemu") ||
		strings.Contains(modelLC, "virtual") ||
		strings.Contains(modelLC, "virtualbox") ||
		strings.Contains(modelLC, "vmware") ||
		strings.Contains(modelLC, "xen") {
		envInfo = append(envInfo, fmt.Sprintf("VM detected (%s)", model))

	}

	if len(envInfo) == 0 {
		return
	}
	fmt.Print("# Environment\n\n")
	for _, ei := range envInfo {
		fmt.Println("-", ei)
	}
	fmt.Print("\n\n")
}
