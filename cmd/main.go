package main

import (
	"log"
	"runtime"

	asset "github.com/YuruDeveloper/tetris/internal/assets"
	"github.com/YuruDeveloper/tetris/internal/keyboard"
	"github.com/YuruDeveloper/tetris/internal/resources"
	"github.com/YuruDeveloper/tetris/internal/renderer"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/YuruDeveloper/tetris/internal/window"
)

func main() {
	runtime.LockOSThread()
	resource.Init()
	window := window.NewWindow()
	err := window.Init(480, 720, types.NewColor(255, 255, 255, 255))
	if err != nil {
		log.Fatalln(err)
		return
	}

	shader, err := asset.GetAssetManager().ShaderAsset(resource.DefaultShaderID)
	if err != nil {
		log.Fatalln(err)
		return
	}
	texture, err := asset.GetAssetManager().TextureAsset(resource.DefaultTextureID)
	if err != nil {
		log.Fatalln(err)
		return
	}
	data, err := renderer.NewRenderObject([]types.Vector2{
		types.NewVector2(1, 1),
		types.NewVector2(1, -1),
		types.NewVector2(-1, 1),
		types.NewVector2(-1, -1),
	},
	[]uint32{
		0, 1, 2, 2, 3, 1,
	},
	[]types.Vector2{
		types.NewVector2(1.0,0.0),
		types.NewVector2(1.0,1.0),
		types.NewVector2(0.0,0.0),
		types.NewVector2(0.0,1.0),
	},
	types.NewVector2(10, 10),
	types.NewVector2(0, 10),
	types.NewVector2(480,720),
	renderer.NewMeterial(shader,texture),
	)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err := data.Init(0); err != nil {
		log.Fatalln(err)
		return
	}

	renderer := renderer.NewRenderer()
	renderer.Set(data)
	window.SetKeyCallBack(keyboard.KeyBoard)

	err = window.Update(renderer)

	if err != nil {
		log.Fatalln(err)
	}
	renderer.Delete()
}
