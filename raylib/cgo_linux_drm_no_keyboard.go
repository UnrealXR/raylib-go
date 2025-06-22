//go:build linux && drm && drm_disable_input && !drm_leasing && !rgfw && !sdl && !sdl3 && !android
// +build linux,drm,drm_disable_input,!drm_leasing,!rgfw,!sdl,!sdl3,!android

package rl

/*
#cgo pkg-config: libdrm gbm
#cgo linux,drm LDFLAGS: -lGLESv2 -lEGL -lpthread -lrt -lm -ldl
#cgo linux,drm CFLAGS: -DDISABLE_EVDEV_INPUT -DPLATFORM_DRM -DGRAPHICS_API_OPENGL_ES2 -DEGL_NO_X11
*/
import "C"
