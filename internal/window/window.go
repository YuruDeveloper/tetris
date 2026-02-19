package window

import (
	packagederror "github.com/YuruDeveloper/tetris/internal/packagedError"
	"github.com/YuruDeveloper/tetris/internal/ports"
	"github.com/YuruDeveloper/tetris/internal/types"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/google/uuid"
)



func NewWindow() *Window {
	return &Window{}
}

type Window struct {
	window     *glfw.Window
	renderer ports.Renderer
}

func (instance *Window) Init(width int, height int, color types.Color,renderer ports.Renderer) error {
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
	instance.window.MakeContextCurrent()
	if err := gl.Init(); err != nil {
		glfw.Terminate()
		return packagederror.NewError(packagederror.FailGLInitError, err.Error())
	}
	instance.renderer = renderer
	glfw.SwapInterval(0)
	floatColor := color.GetColor()
	gl.ClearColor(floatColor.Red, floatColor.Green, floatColor.Blue, floatColor.Alpha)
	renderer.Init(types.NewVector2(float32(width),float32(height)))
	return nil
}

func (instance *Window) NewObject(mesh *types.Reference[types.Mesh],material *types.Handle[types.Meterial],location,size types.Vector2) uuid.UUID {
	return instance.renderer.NewObject(mesh,material,location,size)
}

func (instance *Window) SetLocation(uuid uuid.UUID,location types.Vector2) {
	instance.renderer.SetLocation(uuid,location)
}

func (instance *Window) SetSize(uuid uuid.UUID,size types.Vector2) {
	instance.renderer.SetSize(uuid,size)
}

func (instance *Window) Update(deltaTime float64) {
	instance.renderer.Rendering(deltaTime)
	instance.window.SwapBuffers()
	glfw.PollEvents()
}

func (instance *Window) ShouldClose() bool {
	return instance.window.ShouldClose()
}

func (instance *Window) Close() {
	if instance.window == nil {
		return
	}
	instance.window.SetShouldClose(true)
	glfw.Terminate()
}

func (instance *Window) SetKeyCallBack(callback glfw.KeyCallback) {
	if instance.window == nil {
		return
	}
	instance.window.SetKeyCallback(callback)
}
