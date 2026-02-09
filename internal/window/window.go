package window

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const TargetFPS = 144.0
const TargetDeltaTime = 1.0 / TargetFPS

func NewWindow() *Window {
	return &Window{}
}

type Window struct {
	window     *glfw.Window
	background types.Color
}

func (instance *Window) Init(width int, height int, color types.Color) error {
	if err := glfw.Init(); err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailGLFWInitError, err.Error())
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	window, err := glfw.CreateWindow(width, height, "Tetris", nil, nil)
	if err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailCreateWindow, err.Error())
	}
	instance.window = window
	instance.background = color
	instance.window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailGLInitError, err.Error())
	}
	glfw.SwapInterval(0)
	return nil
}

func (instance *Window) Update(renderer ports.Renderer) error {
	floatColor := instance.background.GetColor()
	gl.ClearColor(floatColor.Red, floatColor.Green, floatColor.Blue, floatColor.Alpha)
	var currentTime float64
	var lastTime float64
	var deltaTime float64
	lastTime = glfw.GetTime()
	for !instance.window.ShouldClose() {
		currentTime = glfw.GetTime()
		deltaTime = currentTime - lastTime
		if deltaTime > TargetDeltaTime {
			lastTime = currentTime
			renderer.Rendering(deltaTime)
			instance.window.SwapBuffers()
			glfw.PollEvents()
		}
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
