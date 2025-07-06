module examples

go 1.21

replace git.lunr.sh/UnrealXR/raylib-go/raylib => ../raylib

replace git.lunr.sh/UnrealXR/raylib-go/raygui => ../raygui

replace git.lunr.sh/UnrealXR/raylib-go/easings => ../easings

replace git.lunr.sh/UnrealXR/raylib-go/physics => ../physics

require (
	git.lunr.sh/UnrealXR/raylib-go/easings v0.0.0-00010101000000-000000000000
	git.lunr.sh/UnrealXR/raylib-go/physics v0.0.0-00010101000000-000000000000
	git.lunr.sh/UnrealXR/raylib-go/raygui v0.0.0-00010101000000-000000000000
	git.lunr.sh/UnrealXR/raylib-go/raylib v0.0.0-20241202103652-5d50abe7c65b
	github.com/jakecoffman/cp v1.2.1
	github.com/neguse/go-box2d-lite v0.0.0-20170921151050-5d8ed9b7272b
)

require (
	github.com/ebitengine/purego v0.8.1 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/sys v0.27.0 // indirect
)
