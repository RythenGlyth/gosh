// +build windows

package runner

import "os"

const (
	pathSep = ";"
)

func isExecutable(info os.FileInfo) bool {
	return info.Mode().IsRegular()
}
