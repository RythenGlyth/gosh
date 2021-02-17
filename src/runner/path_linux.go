// +build linux freebsd netbsd openbsd dragonfly darwin

package runner

import "os"

const (
	path_sep = ":"
	//            rwxrwxrwx
	exec_mask = 0b001001001
)

func isExecutable(info os.FileInfo) bool {
	var m os.FileMode = info.Mode()
	// only symlinks and regular files are executable
	if (m & os.ModeType) != os.ModeSymlink && !info.Mode().IsRegular() {
		return false
	}
	// return whether any of the executable bits is set
	return (m & os.ModePerm & exec_mask) != 0
}
