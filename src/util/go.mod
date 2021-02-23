module github.com/RythenGlyth/gosh/src/util

go 1.15

replace (
	gosh/debug => ../debug
	gosh/gosh => ../gosh
	gosh/lexer => ../lexer
	gosh/runner => ../runner
	gosh/util => ../util
)