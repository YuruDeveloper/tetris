package window

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewWindow() *Window {
	return &Window{}
}

type Window  struct {
	window *glfw.Window
}

func (instance *Window) Init(width , heghit int) {
	if err := glfw.Init() ; err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	
	glfw.WindowHint(glfw.ContextVersionMajor,4)
	glfw.WindowHint(glfw.ContextVersionMinor,6)
	glfw.WindowHint(glfw.OpenGLProfile,glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable,glfw.False)
	window , err := glfw.CreateWindow(width,heghit,"Tetris",nil,nil)
	if err != nil {
		panic(err)
	}
	instance.window = window
	instance.window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		panic(err)
	}
	instance.Drawing()
}

func (instance *Window) Drawing() {
	for !instance.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		instance.window.SwapBuffers()
		glfw.PollEvents()
	}
}