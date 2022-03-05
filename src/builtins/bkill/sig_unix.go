//go:build linux || freebsd || netbsd || openbsd || darwin || dragonfly || solaris

package bkill

import (
	"syscall"
)

var allsigs = []struct {
	name string
	sig  syscall.Signal
}{
	{"HUP", syscall.SIGHUP},
	{"INT", syscall.SIGINT},
	{"QUIT", syscall.SIGQUIT},
	{"ILL", syscall.SIGILL},
	{"TRAP", syscall.SIGTRAP},
	{"ABRT", syscall.SIGABRT},
	{"IOT", syscall.SIGIOT},
	{"BUS", syscall.SIGBUS},
	{"FPE", syscall.SIGFPE},
	{"KILL", syscall.SIGKILL},
	{"USR1", syscall.SIGUSR1},
	{"SEGV", syscall.SIGSEGV},
	{"USR2", syscall.SIGUSR2},
	{"PIPE", syscall.SIGPIPE},
	{"ALRM", syscall.SIGALRM},
	{"TERM", syscall.SIGTERM},
	{"STKFLT", syscall.SIGSTKFLT},
	{"CHLD", syscall.SIGCHLD},
	{"CLD", syscall.SIGCLD},
	{"CONT", syscall.SIGCONT},
	{"STOP", syscall.SIGSTOP},
	{"TSTP", syscall.SIGTSTP},
	{"TTIN", syscall.SIGTTIN},
	{"TTOU", syscall.SIGTTOU},
	{"URG", syscall.SIGURG},
	{"XCPU", syscall.SIGXCPU},
	{"XFSZ", syscall.SIGXFSZ},
	{"VTALRM", syscall.SIGVTALRM},
	{"PROF", syscall.SIGPROF},
	{"WINCH", syscall.SIGWINCH},
	{"IO", syscall.SIGIO},
	{"POLL", syscall.SIGPOLL},
	{"PWR", syscall.SIGPWR},
	{"SYS", syscall.SIGSYS},
}

func byname(sig string) syscall.Signal {
	for _, t := range allsigs {
		if t.name == sig {
			return t.sig
		}
	}

	return syscall.Signal(0)
}

func bynumber(sig int) string {
	for _, t := range allsigs {
		if int(t.sig) == sig {
			return t.name
		}
	}

	return ""
}
