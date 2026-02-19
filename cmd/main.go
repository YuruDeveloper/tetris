package main

import (
	"log"
	"runtime"

	asset "github.com/YuruDeveloper/tetris/internal/assets"
	"github.com/YuruDeveloper/tetris/internal/keyboard"
	"github.com/YuruDeveloper/tetris/internal/renderer"
	"github.com/YuruDeveloper/tetris/internal/resources"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/YuruDeveloper/tetris/internal/window"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const TargetFPS = float64(144)
const TargetDeltaTime = 1.0 / TargetFPS

func main() {
	runtime.LockOSThread()
	manager := asset.GetAssetManager()
	resource.Init(manager)
	window := window.NewWindow()
	renderer := renderer.NewRenderer(manager)
	err := window.Init(480, 720, types.NewColor(255, 255, 255, 255),renderer)
	if err != nil {
		log.Fatalln(err)
		return
	}
	object , err := window.NewObject(0,resource.DefaultMeshID,resource.DefaultMaterialID,types.NewVector2(0,0),types.NewVector2(10,10))
	if err != nil {
		window.Close()
		log.Fatalln(err)
		return
	}
	window.SetKeyCallBack(keyboard.KeyBoard)
	localX := float32(-250)
	old := glfw.GetTime()
	for !window.ShouldClose() {
		new := glfw.GetTime()
		delta := new - old
		if delta < TargetDeltaTime {
			continue
		}
		old = new
		window.Update(delta)
		if localX > 250 {
			localX = -250
		}
		localX += float32(delta) * 100
		window.SetLocation(object,types.NewVector2(localX,0))
	}
	window.Close()
}
