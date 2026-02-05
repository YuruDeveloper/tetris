package keyboard

import(
	"github.com/go-gl/glfw/v3.3/glfw"
)

func KeyBoard(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if mods == glfw.ModControl && key == glfw.KeyC {
			window.SetShouldClose(true)
		}	
}