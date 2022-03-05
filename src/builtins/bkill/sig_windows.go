package bkill

import "syscall"

var allsigs = []syscall.Signal{
	syscall.SIGKILL
}

func byname(sig string) syscall.Signal {
	switch sig {
	case "SIGKILL":
		return syscall.SIGKILL
	default:
		return syscall.Signal(0)
	}
}

func bynumber(sig int) string {
	switch sig {
	case 1:
		return syscall.SIGKILL
	default:
		return ""
	}
}
