# Development - Terminal

To be able to read arbitrary keypresses from the Terminal (both letter but also special keys like Backspace and arrow keys), platform-specific reading & parsing is required.

Terminal I/O all happens inside the termios implementation which currently resides in another tree at [scrouthtv/termios](https://github.com/scrouthtv/termios).

The termios implementation currently works on Windows, Linux and BSD.
It reads a keypress with OS-specific function calls and parses the result to an abstract keypress of type [`Key`](https://pkg.go.dev/github.com/scrouthtv/termios#Key).

This Key is then in turn consumed by the `gosh` implementation.
If any input or output sequence does not work as expected, first look at the [Known issues](https://github.com/scrouthtv/termios#known-issues) page of termios.
