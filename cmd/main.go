package main

import (
	"fmt"

	"gitea.bytedev.duckdns.org/tetris/internal/keyboard"
	"gitea.bytedev.duckdns.org/tetris/internal/renderer"
	"gitea.bytedev.duckdns.org/tetris/internal/types"
	"gitea.bytedev.duckdns.org/tetris/internal/window"
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