package main

import (
	"fmt"

	"github.com/YuruDeveloper/tetris/internal/keyboard"
	"github.com/YuruDeveloper/tetris/internal/renderer"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/YuruDeveloper/tetris/internal/window"
)

func main() {
	window := window.NewWindow()
	err := window.Init(480,720,types.NewColor(255,255,255,255))
	if err != nil {
		fmt.Print(err)
		return
	}
	shaders := renderer.NewShaders()
	err = shaders.LoadFiles()
	if err != nil {
		fmt.Print(err)
		return
	}
	err = shaders.CompileShaders()
	if err != nil {
		fmt.Print(err)
		return
	}
	window.SetKeyCallBack(keyboard.KeyBoard)
	err = window.Update()
	if err != nil {
		fmt.Print(err)
		return
	}
}