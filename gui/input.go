package gui

import "github.com/go-gl/glfw/v3.2/glfw"

// send typed runes straight through to the pty
func (gui *GUI) char(w *glfw.Window, r rune) {
	gui.terminal.Write([]byte(string(r)))
}

func (gui *GUI) key(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Repeat || action == glfw.Press {

		gui.logger.Debugf("KEY PRESS: key=0x%X scan=0x%X", key, scancode)

		switch true {
		case mods&glfw.ModControl > 0:

			if mods&glfw.ModAlt > 0 {
				// ctrl + alt +
				switch key {
				case glfw.KeyV:
					// todo handle both these errors
					if buf, err := gui.window.GetClipboardString(); err == nil {
						_ = gui.terminal.Write([]byte(buf))
					}
				}
			} else {
				// ctrl +
				switch key {
				case glfw.KeyC: // ctrl^c
					gui.logger.Debugf("Sending CTRL^C")
					gui.terminal.Write([]byte{0x3}) // send EOT
				}
			}
		}

		switch key {
		case glfw.KeyEnter:
			gui.terminal.Write([]byte{0x0a})
		case glfw.KeyBackspace:
			gui.terminal.Write([]byte{0x08})
		}

		//gui.logger.Debugf("Key pressed: 0x%X %q", key, string([]byte{byte(key)}))
		//gui.terminal.Write([]byte{byte(scancode)})
	}

}
