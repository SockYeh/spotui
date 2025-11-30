package utils

import (
	"os/exec"
	"runtime"
	"strings"
)

func OpenURL(url string) error {
    switch runtime.GOOS {
    case "windows":
        return exec.Command("cmd", "/c", "start", "", strings.ReplaceAll(url, "&","^&")).Start()
    case "darwin":
        return exec.Command("open", url).Start()
    default:
        return exec.Command("xdg-open", url).Start()
    }
}
