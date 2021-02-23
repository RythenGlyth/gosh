module github.com/RythenGlyth/gosh/src/gosh

go 1.15

require github.com/scrouthtv/termios v0.0.0-20210223081113-535b04571c2f

replace (
	gosh/debug => ../debug
	gosh/gosh => ../gosh
	gosh/lexer => ../lexer
	gosh/runner => ../runner
	gosh/util => ../util
)