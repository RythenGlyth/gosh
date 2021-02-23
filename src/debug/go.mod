module github.com/RythenGlyth/gosh/src/debug

go 1.15

require github.com/scrouthtv/golang-ipc v0.0.0-20210223075931-56b586a10741

replace (
	gosh/debug => ../debug
	gosh/gosh => ../gosh
	gosh/lexer => ../lexer
	gosh/runner => ../runner
	gosh/util => ../util
)