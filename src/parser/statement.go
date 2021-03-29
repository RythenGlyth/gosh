package parser

import "gosh/src/shared"

type ExecutableStatement interface {
	Exec(shared.IGosh)
}
