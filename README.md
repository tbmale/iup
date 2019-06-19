# IUP
GoLang wrapper for IUP GUI library (3.18)

This is a work in progress, but since it's a simple wrapper, I think it is ready for reasonably serious work.

It is only tested and developped for windows (actually on a 64bits Windows 7) since I have absolutly no experience with linux.

# Installation
To install and use this wrapper, you first need to download it

```
go get github.com/tbmale/iup
```

Then, you must create a directory inside the iup folder called "lib" and put in it the static IUP libraries. You also can if you want/need to use IM or CD create the folders "im" and "cd" inside the "lib" folder and put in it the IM and CD static libraries.

Download links:
- IUP: https://sourceforge.net/projects/iup/files/3.18/Windows%20Libraries/
- CD: https://sourceforge.net/projects/canvasdraw/files/5.9/Windows%20Libraries/
- IM: https://sourceforge.net/projects/imtoolkit/files/3.10/Windows%20Libraries/

Please download Mingw GCC static or dynamic libraries (Win32_mingw4, Win64_mingw4, Win32_dllw4, Win64_dllw4).

Then you can test this hello world example:

```go
package main

import "github.com/tbmale/iup"

func main() {
    iup.Open()
    defer iup.Close()

    dlg := iup.Dialog(iup.Vbox(
        iup.Label("Hello, world!"),
        iup.Button("Click me").SetCallback("ACTION", func(ih iup.Ihandle) int {
            return iup.CLOSE
        }),
    ).SetAttributes(`MARGIN=10x10,GAP=10`))

    iup.Show(dlg)
    iup.MainLoop()
}
```

# Additional libraries
The sub folders in the IUP folder are additional libraries of IUP. They all depend on the main iup lib.

They must be subfolders in order to be able to find the static libs.

# Known issues
- iupweb doesn't work because it needs MS Visual C++ compiler (OLE...).
- I don't think it's really an issue when using Go, but Lua stuff is not (and will not be) implemented.

# Todo
- Find a way to compile the IUP lib from source (as I said, absolutly no experience with unix tools (make, configure...)) to include directly the source code of IUP to this wrapper.
- Make wrappers for CD and IM (other TECGRAF's libraries) to use them with IUP.
