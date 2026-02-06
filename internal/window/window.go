package window

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func NewWindow() *Window {
	return &Window{}
}

type Window  struct {
	window *glfw.Window
	background types.Color
	program uint32
}

func (instance *Window) Init(width int, height int,color types.Color) error {
	if err := glfw.Init() ; err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailGLFWInitError,err.Error())
	}
	glfw.WindowHint(glfw.ContextVersionMajor,4)
	glfw.WindowHint(glfw.ContextVersionMinor,6)
	glfw.WindowHint(glfw.OpenGLProfile,glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable,glfw.False)
	window , err := glfw.CreateWindow(width,height,"Tetris",nil,nil)
	if err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailCreateWindow,err.Error())
	}
	instance.window = window
	instance.background = color
	instance.window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailGLInitError,err.Error())
	}	
	return nil
}

func (instance *Window) Update() error {
	floatColor := instance.background.GetColor()
	gl.ClearColor(floatColor.Red,floatColor.Green,floatColor.Blue,floatColor.Alpha)
	for !instance.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)
		if instance.program != 0 && gl.IsProgram(instance.program) {
			gl.UseProgram(instance.program)
		}
		instance.window.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
	return nil
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

func (instance *Window) SetProgram(program uint32) {
	instance.program = program
}