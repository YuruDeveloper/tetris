package window

import (
	"gitea.bytedev.duckdns.org/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewWindow() *Window {
	return &Window{}
}

type Window  struct {
	window *glfw.Window
	background types.Color
}

func (instance *Window) Init(width int, height int,color types.Color) {
	if err := glfw.Init() ; err != nil {
		panic(err)
	}
	glfw.WindowHint(glfw.ContextVersionMajor,4)
	glfw.WindowHint(glfw.ContextVersionMinor,6)
	glfw.WindowHint(glfw.OpenGLProfile,glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable,glfw.False)
	window , err := glfw.CreateWindow(width,height,"Tetris",nil,nil)
	if err != nil {
		panic(err)
	}
	instance.window = window
	instance.background = color
}

func (instance *Window) Update() {
	instance.window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		panic(err)
	}	
	floatColor := instance.background.GetColor()
	gl.ClearColor(floatColor.Red,floatColor.Green,floatColor.Blue,floatColor.Alpha)
	for !instance.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		instance.window.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
}

func (instance *Window) Close() {
	if instance.window == nil {
		return
	}
	instance.window.SetShouldClose(true)
}

func (instance *Window) SetKeyCallBack(callback glfw.KeyCallback) {
	if instance.window == nil {
		return
	}
	instance.window.SetKeyCallback(callback)
}