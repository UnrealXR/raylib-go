//go:build linux && drm && !drm_leasing && !drm_disable_input && !rgfw && !sdl && !sdl3 && !android
// +build linux,drm,!drm_leasing,!drm_disable_input,!rgfw,!sdl,!sdl3,!android

package rl

/*
#cgo pkg-config: libdrm gbm
#cgo linux,drm LDFLAGS: -lGLESv2 -lEGL -lpthread -lrt -lm -ldl
#cgo linux,drm CFLAGS: -DSUPPORT_SSH_KEYBOARD_RPI -DPLATFORM_DRM -DGRAPHICS_API_OPENGL_ES2 -DEGL_NO_X11
*/
import "C"
