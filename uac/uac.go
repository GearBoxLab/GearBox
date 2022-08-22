//go:build windows

package uac

import (
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

// Prompt triggers Windows UAC elevation prompt.
func Prompt() error {
	return PromptWithExtraArguments([]string{})
}

// PromptWithExtraArguments triggers Windows UAC elevation prompt.
// The extraArguments will append to the original command arguments.
func PromptWithExtraArguments(extraArguments []string) error {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(append(os.Args[1:], extraArguments...), " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 // SW_NORMAL

	return windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
}

func IsAdmin() bool {
	systemRoot := os.Getenv("SYSTEMROOT")
	cmd := exec.Command(systemRoot+`\system32\cacls.exe`, systemRoot+`\system32\config\system`)

	if err := cmd.Run(); nil != err {
		return false
	}

	return true
}
