package pkg

import (
	"os/exec"
	"strings"
)

func getWindowsEnvVariable(varName string) string {
	cmd := exec.Command("cmd", "/C", "echo", "%"+varName+"%")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(output))
}
