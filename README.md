![logo](https://git.lunr.sh/UnrealXR/raylib-go/raw/branch/master/assets/raylib-go_256x256.png)
## raylib-go
[![Build Status](https://git.lunr.sh/UnrealXR/raylib-go/actions/workflows/build.yml/badge.svg)](https://git.lunr.sh/UnrealXR/raylib-go/actions)
[![GoDoc](https://godoc.org/git.lunr.sh/UnrealXR/raylib-go/raylib?status.svg)](https://godoc.org/git.lunr.sh/UnrealXR/raylib-go/raylib)
[![Go Report Card](https://goreportcard.com/badge/git.lunr.sh/UnrealXR/raylib-go/raylib)](https://goreportcard.com/report/git.lunr.sh/UnrealXR/raylib-go/raylib)
[![Examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg?style=flat-square)](https://git.lunr.sh/UnrealXR/raylib-go/src/branch/master/examples)

Golang bindings for [raylib](http://www.raylib.com/), a simple and easy-to-use library to enjoy videogames programming.

This is a fork of [gen2brain's `raylib-go` bindings](https://github.com/gen2brain/raylib-go) to add [DRM leasing](https://wayland.app/protocols/drm-lease-v1) support.

raylib C source code is included and compiled together with bindings. Note that the first build can take a few minutes.

It is also possible to use raylib-go without cgo (Windows only; see requirements below).

### Requirements

##### Ubuntu

    apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libwayland-dev libxkbcommon-dev

##### Fedora

    dnf install mesa-libGL-devel libXi-devel libXcursor-devel libXrandr-devel libXinerama-devel wayland-devel libxkbcommon-devel

##### macOS

On macOS, you need Xcode or Command Line Tools for Xcode (if you have `brew` installed, you already have this).

##### Windows

###### cgo

On Windows you need C compiler, like [Mingw-w64](https://mingw-w64.org) or [TDM-GCC](http://tdm-gcc.tdragon.net/).
You can also build binary in [MSYS2](https://msys2.github.io/) shell.

To remove console window, build with `-ldflags "-H=windowsgui"`.

###### purego (without cgo, i.e. CGO_ENABLED=0)

Download the raylib.dll from the assets on the [releases page](https://github.com/raysan5/raylib/releases). It is contained in the `raylib-*_win64_msvc*.zip`.
Put the raylib.dll into the root folder of your project or copy it into `C:\Windows\System32` for a system-wide installation.

It is also possible to build the DLL yourself. You can find more info at [raylib's wiki](https://github.com/raysan5/raylib/wiki/Working-on-Windows).

##### Android

[Android example](https://git.lunr.sh/UnrealXR/raylib-go/tree/master/examples/others/android/example).

##### Wasm

For web bindings, refer to [Raylib-Go-Wasm](https://github.com/BrownNPC/Raylib-Go-Wasm); it should be largely compatible with this repository.


### Installation

    go get -v -u git.lunr.sh/UnrealXR/raylib-go/raylib

### Build tags

* `drm` - build for Linux native [DRM](https://en.wikipedia.org/wiki/Direct_Rendering_Manager) mode, including Raspberry Pi 4 and other devices (PLATFORM_DRM)
* `drm_disable_input` - disables keyboard input capabilities for the DRM backend. Requires the `drm` build tag as a prerequisite
* `drm_leasing` - enables Wayland DRM leasing capabilties. Requires the `drm` build tag as a prerequisite
* `sdl` - build for [SDL](https://github.com/libsdl-org/SDL) backend (PLATFORM_DESKTOP_SDL)
* `sdl3` - build for [SDL3](https://github.com/libsdl-org/SDL) backend (PLATFORM_DESKTOP_SDL3)
* `rgfw` - build for [RGFW](https://github.com/ColleagueRiley/RGFW) backend (PLATFORM_DESKTOP_RGFW)
* `noaudio` - disables audio functions
* `opengl43` - uses OpenGL 4.3 backend
* `opengl21` - uses OpenGL 2.1 backend (default is 3.3 on desktop)
* `opengl11` - uses OpenGL 1.1 backend (pseudo OpenGL 1.1 style)
* `es2` - uses OpenGL ES 2.0 backend (can be used to link against [Google's ANGLE](https://github.com/google/angle))
* `es3` - experimental support for OpenGL ES 3.0
* `x11` - force X11 compatibility mode on Wayland (PLATFORM_DESKTOP/GLFW)
* `wayland` - force Wayland only mode (PLATFORM_DESKTOP/GLFW)

### Documentation

Documentation on [GoDoc](https://godoc.org/git.lunr.sh/UnrealXR/raylib-go/raylib). Also check raylib [cheatsheet](http://www.raylib.com/cheatsheet/cheatsheet.html). If you have problems or need assistance there is an active community in the #raylib-go channel of the [Raylib Discord Server](https://discord.gg/raylib) that can help.

### Example

```go
package main

import rl "git.lunr.sh/UnrealXR/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
```

Check more [examples](https://git.lunr.sh/UnrealXR/raylib-go/tree/master/examples) organized by raylib modules.

### Cross-compile (Linux)

To cross-compile for Windows install [MinGW](https://www.mingw-w64.org/) toolchain.

```
$ CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags "-s -w"
$ file basic_window.exe
basic_window.exe: PE32+ executable (console) x86-64 (stripped to external PDB), for MS Windows, 11 sections

$ CGO_ENABLED=1 CC=i686-w64-mingw32-gcc GOOS=windows GOARCH=386 go build -ldflags "-s -w"
$ file basic_window.exe
basic_window.exe: PE32 executable (console) Intel 80386 (stripped to external PDB), for MS Windows, 9 sections
```

To cross-compile for macOS install [OSXCross](https://github.com/tpoechtrager/osxcross) toolchain.

```
$ CGO_ENABLED=1 CC=x86_64-apple-darwin21.1-clang GOOS=darwin GOARCH=amd64 go build -ldflags "-linkmode external -s -w '-extldflags=-mmacosx-version-min=10.15'"
$ file basic_window
basic_window: Mach-O 64-bit x86_64 executable, flags:<NOUNDEFS|DYLDLINK|TWOLEVEL>

$ CGO_ENABLED=1 CC=aarch64-apple-darwin21.1-clang GOOS=darwin GOARCH=arm64 go build -ldflags "-linkmode external -s -w '-extldflags=-mmacosx-version-min=12.0.0'"
$ file basic_window
basic_window: Mach-O 64-bit arm64 executable, flags:<NOUNDEFS|DYLDLINK|TWOLEVEL|PIE>
```

### License

raylib-go is licensed under an unmodified zlib/libpng license. View [LICENSE](https://git.lunr.sh/UnrealXR/raylib-go/blob/master/LICENSE).
