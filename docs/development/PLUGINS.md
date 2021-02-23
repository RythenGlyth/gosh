# Development - Plugins

To allow every user to customize `gosh` to their needs, we introduced Plugin Support.

The basic idea is that every plugin is individually compiled to a dynamic library that is loaded at runtime by the plugin handler.

For the Go compiler to find the `gosh` source, the plugin has to be written inside the `gosh` tree.

## Creating Plugins

1. Get the `gosh` source
```
 ~ > git clone https://github.com/RythenGlyth/gosh
 ~ > cd gosh
 ~/gosh > mkdir myplugin
 ~/gosh > cd myplugin
```
2. Create a basic plugin (name this file `main.go`):
```
package main

import "gosh/src/gosh"

func OnKey(g *gosh.Gosh, s string) bool {
    return s == "ok"
}
```
3. Compile the plugin:
```
 ~/gosh/myplugin > go build -buildmode=plugin .
```
Now you should end up with an `example-mod.so`. Remember its location and load it into `gosh`.