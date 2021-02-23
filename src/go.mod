module gosh

go 1.15

require (
	github.com/scrouthtv/golang-ipc v0.0.0-20210223075931-56b586a10741 // indirect
	golang.org/x/sys v0.0.0-20210220050731-9a76102bfb43 // indirect
	gosh/debug v0.0.0-00010101000000-000000000000
	gosh/gosh v0.0.0-00010101000000-000000000000
	gosh/util v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	gosh/debug => ./debug
	gosh/gosh => ./gosh
	gosh/lexer => ./lexer
	gosh/runner => ./runner
	gosh/util => ./util
)
