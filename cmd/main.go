package main

import "gitea.bytedev.duckdns.org/tetris/internal/window"

func main() {
	window := window.NewWindow()
	window.Init(480,720)
}