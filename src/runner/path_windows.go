// +build windows

package runner

import "os"

const (
	path_sep = ";"
)

func isExecutable(info *os.FileInfo) bool {
	return info.Mode().IsRegular()
}
