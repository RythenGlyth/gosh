# Development

`gosh` uses a modular structure. Every aspect of the shell is implemented in its own module inside the tree.

All loaded modules are stored as members of the `gosh` instance.

## Installing gosh

`gosh` can currently only be built from source:

```bash
 ~ git clone https://github.com/RythenGlyth/gosh
 ~ cd gosh/src
 ~ go build -o goshell
```

The `goshell` executable can then be installed to the OS.