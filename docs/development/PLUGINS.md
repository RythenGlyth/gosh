# Development - Plugins

To allow every user to customize `gosh` to their needs, we introduced Plugin Support.

The basic idea is that every plugin is individually compiled to a dynamic library that is loaded at runtime by the plugin handler.

## Creating Plugins

1. Create a new folder for the plugin:
```
 ~ > mkdir myplugin
 ~ > cd myplugin
 ~/myplugin > go mod init main

```
It is very important that the module name is `main`.
2. Create a basic plugin (name this file `main.go`):
```
package main

import "github.com/RythenGlyth/gosh/src/gosh"

func OnKey(g *gosh.Gosh, s string) bool {
    return s == "ok"
}
```
3. Compile the plugin:
```
 ~/myplugin > go build -buildmode=plugin .
```
Now you should end up with a `main.so`. Remember its location and load it into `gosh`.