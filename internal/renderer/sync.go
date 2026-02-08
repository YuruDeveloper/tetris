package renderer

import "github.com/go-gl/gl/v4.6-core/gl"

const TimeOut = 100

type Sync struct {
	fence uintptr
}

func NewSync() *Sync {
	return &Sync{}
}

func (instance *Sync) NewFence() {
	instance.fence = gl.FenceSync(gl.SYNC_FLUSH_COMMANDS_BIT, 0)
}

func (instance *Sync) WaitSync() {
	if instance.fence != 0 {
		gl.ClientWaitSync(instance.fence, gl.SYNC_FLUSH_COMMANDS_BIT, TimeOut)
		gl.DeleteSync(instance.fence)
		instance.fence = 0
	}
}

func (instance *Sync) Delete() {
	if instance.fence != 0 {
		gl.DeleteSync(instance.fence)
	}
}
