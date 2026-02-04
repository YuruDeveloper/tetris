package main

import (
	"log"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {

	if err := glfw.Init() ; err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	// opnegl version 
	glfw.WindowHint(glfw.ContextVersionMajor,4)
	glfw.WindowHint(glfw.ContextVersionMinor,6)
	glfw.WindowHint(glfw.OpenGLProfile,glfw.OpenGLCoreProfile)
	// 리사이즈 불가능 하게 하기
	glfw.WindowHint(glfw.Resizable,glfw.False)
	// 테두리 없에기
	//glfw.WindowHint(glfw.Decorated,glfw.False)
	// 투명 배경
	// glfw.WindowHint(glfw.TransparentFramebuffer,glfw.True)
	//  촤상위 로 띄우기
	// glfw.WindowHint(glfw.Floating,glfw.True)
  
	window , err := glfw.CreateWindow(480,480,"Tetris",nil,nil)
	if err != nil {
		log.Println(err)
		return
	}
	window.MakeContextCurrent()
	
	if err := gl.Init(); err != nil {
		panic(err)
	}

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}