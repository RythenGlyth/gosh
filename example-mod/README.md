# Example module

Here is a basic example module that showcases how modules work.

If loaded, it only allows you to type text at even minutes, but blocks you at odd minutes.

## Compilation

Download the `gosh` source tree, cd into the module folder and compile it using `go build -buildmode=plugin .`. Then load it from a running `gosh`. 