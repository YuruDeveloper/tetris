package main

import (
	"gitea.bytedev.duckdns.org/tetris/internal/keyboard"
	"gitea.bytedev.duckdns.org/tetris/internal/types"
	"gitea.bytedev.duckdns.org/tetris/internal/window"
)

func main() {
	window := window.NewWindow()
	window.Init(480,720,types.NewColor(1,1,1,1))
	window.SetKeyCallBack(keyboard.KeyBoard)
	window.Update()
}