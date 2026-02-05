package main

import (
	"gitea.bytedev.duckdns.org/tetris/internal/keyboard"
	"gitea.bytedev.duckdns.org/tetris/internal/types"
	"gitea.bytedev.duckdns.org/tetris/internal/window"
)

func main() {
	window := window.NewWindow()
	err := window.Init(480,720,types.NewColor(1,1,1,1))
	if err != nil {
		return
	}
	window.SetKeyCallBack(keyboard.KeyBoard)
	err = window.Update()
}