package libs

import (
	"github.com/gotranspile/cxgo/types"
)

const (
	glfw3H = "GLFW/glfw3.h"
)

func init() {
	RegisterLibrary(glfw3H, func(env *Env) *Library {
		intPtrT := env.PtrT(env.Go().Int())
		hintT := types.NamedTGo("Hint", "glfw.Hint", env.Go().Int())
		actionT := types.NamedTGo("Action", "glfw.Action", env.Go().Int())
		joystickHatStateT := types.NamedTGo("JoystickHatState", "glfw.JoystickHatState", env.Go().Int())
		keyT := types.NamedTGo("_GLFWKey", "glfw.Key", env.Go().Byte())
		mouseButtonT := types.NamedTGo("MouseButton", "glfw.MouseButton", env.Go().Int())
		modKeyT := types.NamedTGo("ModifierKey", "glfw.ModifierKey", env.Go().Int())
		inputModeT := types.NamedTGo("InputMode", "glfw.InputMode", env.Go().Int())
		perEventT := types.NamedTGo("PeripheralEvent", "glfw.PeripheralEvent", env.Go().Int())
		windowPtrT := env.PtrT(nil)
		keyCbT := types.NamedTGo("GLFWkeyfun", "glfw.KeyCallback", env.FuncTT(nil, windowPtrT, keyT, env.Go().Int(), actionT, modKeyT))
		charCbT := types.NamedTGo("GLFWcharfun", "glfw.CharCallback", env.FuncTT(nil, windowPtrT, env.Go().Rune()))
		frameBufCb := types.NamedTGo("GLFWframebuffersizefun", "glfw.FramebufferSizeCallback", env.FuncTT(nil, windowPtrT, env.Go().Rune()))
		errorCb := types.NamedTGo("GLFWerrorfun", "glfw.ErrorCallback", env.FuncTT(nil, env.Go().Int(), env.Go().String()))
		videoModeT := types.NamedTGo("GLFWvidmode", "glfw.VidMode", types.StructT([]*types.Field{
			{Name: types.NewIdentGo("width", "Width", env.Go().Int())},
			{Name: types.NewIdentGo("height", "Height", env.Go().Int())},
			{Name: types.NewIdentGo("redBits", "RedBits", env.Go().Int())},
			{Name: types.NewIdentGo("greenBits", "GreenBits", env.Go().Int())},
			{Name: types.NewIdentGo("blueBits", "BlueBits", env.Go().Int())},
			{Name: types.NewIdentGo("refreshRate", "RefreshRate", env.Go().Int())},
		}))
		monitorT := types.NamedTGo("GLFWmonitor", "glfw.Monitor", env.MethStructT(map[string]*types.FuncType{
			"GetVideoMode":    env.FuncTT(env.PtrT(videoModeT), nil),
			"GetPos":          env.FuncTT(nil, env.PtrT(env.Go().Int()), env.PtrT(env.Go().Int())),
			"GetPhysicalSize": env.FuncTT(nil, env.PtrT(env.Go().Int()), env.PtrT(env.Go().Int())),
			"GetName":         env.FuncTT(env.Go().String(), nil),
		}))
		monitorCb := types.NamedTGo("GLFWmonitorfun", "glfw.MonitorCallback", env.FuncTT(nil, env.PtrT(monitorT), perEventT))
		joystickT := types.NamedTGo("_GLFWJoystick", "glfw.Joystick", env.MethStructT(map[string]*types.FuncType{
			"GetAxes":        env.FuncTT(types.SliceT(types.FloatT(4)), env.PtrT(env.Go().Int())),
			"GetButtons":     env.FuncTT(types.SliceT(actionT), env.PtrT(env.Go().Int())),
			"GetHats":        env.FuncTT(types.SliceT(joystickHatStateT), env.PtrT(env.Go().Int())),
			"GetName":        env.FuncTT(env.Go().String(), nil),
			"GetGamepadName": env.FuncTT(env.Go().String(), nil),
			"IsGamepad":      env.FuncTT(env.Go().Bool(), nil),
			"GetGUID":        env.FuncTT(env.Go().String(), nil),
		}))
		joystickCb := types.NamedTGo("GLFWjoystickfun", "glfw.JoystickCallback", env.FuncTT(nil, joystickT, perEventT))
		posCb := types.NamedTGo("GLFWwindowposfun", "glfw.PosCallback", env.FuncTT(nil, windowPtrT, env.Go().Int(), env.Go().Int()))
		sizeCb := types.NamedTGo("GLFWwindowsizefun", "glfw.SizeCallback", env.FuncTT(nil, windowPtrT, env.Go().Int(), env.Go().Int()))
		scaleCb := types.NamedTGo("GLFWwindowcontentscalefun", "glfw.ContentScaleCallback", env.FuncTT(nil, windowPtrT, env.C().Float(), env.C().Float()))
		closeCb := types.NamedTGo("GLFWwindowclosefun", "glfw.CloseCallback", env.FuncTT(nil, windowPtrT))
		refreshCb := types.NamedTGo("GLFWwindowrefreshfun", "glfw.RefreshCallback", env.FuncTT(nil, windowPtrT))
		focusCb := types.NamedTGo("GLFWwindowfocusfun", "glfw.FocusCallback", env.FuncTT(nil, windowPtrT, env.Go().Bool()))
		windowT := types.NamedTGo("GLFWwindow", "glfw.Window", env.MethStructT(map[string]*types.FuncType{
			"MakeContextCurrent": env.FuncTT(nil, nil),
			"ShouldClose":        env.FuncTT(env.Go().Bool(), nil),
			"SwapBuffers":        env.FuncTT(nil, nil),
			"GetKey":             env.FuncTT(actionT, keyT),
			"GetInputMode":       env.FuncTT(env.Go().Int(), inputModeT),
			"SetInputMode":       env.FuncTT(nil, inputModeT, env.Go().Int()),
			"SetShouldClose":     env.FuncTT(nil, env.Go().Bool()),
			"GetFramebufferSize": env.FuncTT(nil, intPtrT, intPtrT),
			"Destroy":            env.FuncTT(nil, nil),
			"Focus":              env.FuncTT(nil, nil),
			"Maximize":           env.FuncTT(nil, nil),
			"Show":               env.FuncTT(nil, nil),
			"Hide":               env.FuncTT(nil, nil),
			"Iconify":            env.FuncTT(nil, nil),
			"Restore":            env.FuncTT(nil, nil),
			"SetTitle":           env.FuncTT(nil, env.C().String()),
			"SetSize":            env.FuncTT(nil, env.Go().Int(), env.Go().Int()),
			"GetSize":            env.FuncTT(nil, env.PtrT(env.Go().Int()), env.PtrT(env.Go().Int())),
			"SetPos":             env.FuncTT(nil, env.Go().Int(), env.Go().Int()),
			"GetPos":             env.FuncTT(nil, env.PtrT(env.Go().Int()), env.PtrT(env.Go().Int())),
			"GetUserPointer":     env.FuncTT(env.PtrT(nil), nil),
			"SetUserPointer":     env.FuncTT(nil, env.PtrT(nil)),
			"GetAttrib":          env.FuncTT(env.Go().Int(), env.Go().Int()),
			"SetAttrib":          env.FuncTT(nil, env.Go().Int(), env.Go().Int()),
			"GetMonitor":         env.FuncTT(env.PtrT(monitorT)),
			"SetMonitor":         env.FuncTT(nil, windowPtrT, env.PtrT(monitorT), env.Go().Int(), env.Go().Int(), env.Go().Int(), env.Go().Int(), env.Go().Int()),
			// callbacks
			"SetKeyCallback":             env.FuncTT(keyCbT, keyCbT),
			"SetCharCallback":            env.FuncTT(charCbT, charCbT),
			"SetFramebufferSizeCallback": env.FuncTT(frameBufCb, frameBufCb),
			"SetPosCallback":             env.FuncTT(posCb, posCb),
			"SetSizeCallback":            env.FuncTT(sizeCb, sizeCb),
			"SetContentScaleCallback":    env.FuncTT(scaleCb, scaleCb),
			"SetCloseCallback":           env.FuncTT(closeCb, closeCb),
			"SetRefreshCallback":         env.FuncTT(refreshCb, refreshCb),
			"SetFocusCallback":           env.FuncTT(focusCb, focusCb),
		}))
		windowPtrT.SetElem(windowT)
		l := &Library{
			Imports: map[string]string{
				"glfw": RuntimeOrg + "/glfw",
			},
			Types: map[string]types.Type{
				"_GLFWkey":      keyT,
				"GLFWwindow":    windowT,
				"GLFWmonitor":   monitorT,
				"GLFWvidmode":   videoModeT,
				"_GLFWjoystick": joystickT,
				"_Action":       actionT,
			},
			Idents: map[string]*types.Ident{
				// functions
				"glfwWindowHint": types.NewIdentGo("glfwWindowHint", "glfw.WindowHint", env.FuncTT(nil, hintT, env.Go().Int())),
				"glfwGetKeyName": types.NewIdentGo("glfwGetKeyName", "glfw.GetKeyName", env.FuncTT(env.Go().String(), keyT, env.Go().Int())),
				// constants
				"GLFW_OPENGL_CORE_PROFILE":   types.NewIdentGo("GLFW_OPENGL_CORE_PROFILE", "glfw.OpenGLCoreProfile", env.Go().Int()),
				"GLFW_OPENGL_ANY_PROFILE":    types.NewIdentGo("GLFW_OPENGL_ANY_PROFILE", "glfw.OpenGLAnyProfile", env.Go().Int()),
				"GLFW_OPENGL_COMPAT_PROFILE": types.NewIdentGo("GLFW_OPENGL_COMPAT_PROFILE", "glfw.OpenGLCompatProfile", env.Go().Int()),
				"GLFW_TRUE":                  types.NewIdentGo("GLFW_TRUE", "glfw.True", env.Go().Int()),
				"GLFW_FALSE":                 types.NewIdentGo("GLFW_FALSE", "glfw.False", env.Go().Int()),
				// key and button actions
				"GLFW_PRESS":   types.NewIdentGo("GLFW_PRESS", "glfw.Press", actionT),
				"GLFW_RELEASE": types.NewIdentGo("GLFW_RELEASE", "glfw.Release", actionT),
				"GLFW_REPEAT":  types.NewIdentGo("GLFW_REPEAT", "glfw.Repeat", actionT),
				/* The unknown key */
				"GLFW_KEY_UNKNOWN": types.NewIdentGo("GLFW_KEY_UNKNOWN", "glfw.KeyUnknown", keyT),
				/* Printable keys */
				"GLFW_KEY_SPACE":         types.NewIdentGo("GLFW_KEY_SPACE", "glfw.KeySpace", keyT),
				"GLFW_KEY_APOSTROPHE":    types.NewIdentGo("GLFW_KEY_APOSTROPHE", "glfw.KeyApostrophe", keyT),
				"GLFW_KEY_COMMA":         types.NewIdentGo("GLFW_KEY_COMMA", "glfw.KeyComma", keyT),
				"GLFW_KEY_MINUS":         types.NewIdentGo("GLFW_KEY_MINUS", "glfw.KeyMinus", keyT),
				"GLFW_KEY_PERIOD":        types.NewIdentGo("GLFW_KEY_PERIOD", "glfw.KeyPeriod", keyT),
				"GLFW_KEY_SLASH":         types.NewIdentGo("GLFW_KEY_SLASH", "glfw.KeySlash", keyT),
				"GLFW_KEY_0":             types.NewIdentGo("GLFW_KEY_0", "glfw.Key0", keyT),
				"GLFW_KEY_1":             types.NewIdentGo("GLFW_KEY_1", "glfw.Key1", keyT),
				"GLFW_KEY_2":             types.NewIdentGo("GLFW_KEY_2", "glfw.Key2", keyT),
				"GLFW_KEY_3":             types.NewIdentGo("GLFW_KEY_3", "glfw.Key3", keyT),
				"GLFW_KEY_4":             types.NewIdentGo("GLFW_KEY_4", "glfw.Key4", keyT),
				"GLFW_KEY_5":             types.NewIdentGo("GLFW_KEY_5", "glfw.Key5", keyT),
				"GLFW_KEY_6":             types.NewIdentGo("GLFW_KEY_6", "glfw.Key6", keyT),
				"GLFW_KEY_7":             types.NewIdentGo("GLFW_KEY_7", "glfw.Key7", keyT),
				"GLFW_KEY_8":             types.NewIdentGo("GLFW_KEY_8", "glfw.Key8", keyT),
				"GLFW_KEY_9":             types.NewIdentGo("GLFW_KEY_9", "glfw.Key9", keyT),
				"GLFW_KEY_SEMICOLON":     types.NewIdentGo("GLFW_KEY_SEMICOLON", "glfw.KeySemicolon", keyT),
				"GLFW_KEY_EQUAL":         types.NewIdentGo("GLFW_KEY_EQUAL", "glfw.KeyEqual", keyT),
				"GLFW_KEY_A":             types.NewIdentGo("GLFW_KEY_A", "glfw.KeyA", keyT),
				"GLFW_KEY_B":             types.NewIdentGo("GLFW_KEY_B", "glfw.KeyB", keyT),
				"GLFW_KEY_C":             types.NewIdentGo("GLFW_KEY_C", "glfw.KeyC", keyT),
				"GLFW_KEY_D":             types.NewIdentGo("GLFW_KEY_D", "glfw.KeyD", keyT),
				"GLFW_KEY_E":             types.NewIdentGo("GLFW_KEY_E", "glfw.KeyE", keyT),
				"GLFW_KEY_F":             types.NewIdentGo("GLFW_KEY_F", "glfw.KeyF", keyT),
				"GLFW_KEY_G":             types.NewIdentGo("GLFW_KEY_G", "glfw.KeyG", keyT),
				"GLFW_KEY_H":             types.NewIdentGo("GLFW_KEY_H", "glfw.KeyH", keyT),
				"GLFW_KEY_I":             types.NewIdentGo("GLFW_KEY_I", "glfw.KeyI", keyT),
				"GLFW_KEY_J":             types.NewIdentGo("GLFW_KEY_J", "glfw.KeyJ", keyT),
				"GLFW_KEY_K":             types.NewIdentGo("GLFW_KEY_K", "glfw.KeyK", keyT),
				"GLFW_KEY_L":             types.NewIdentGo("GLFW_KEY_L", "glfw.KeyL", keyT),
				"GLFW_KEY_M":             types.NewIdentGo("GLFW_KEY_M", "glfw.KeyM", keyT),
				"GLFW_KEY_N":             types.NewIdentGo("GLFW_KEY_N", "glfw.KeyN", keyT),
				"GLFW_KEY_O":             types.NewIdentGo("GLFW_KEY_O", "glfw.KeyO", keyT),
				"GLFW_KEY_P":             types.NewIdentGo("GLFW_KEY_P", "glfw.KeyP", keyT),
				"GLFW_KEY_Q":             types.NewIdentGo("GLFW_KEY_Q", "glfw.KeyQ", keyT),
				"GLFW_KEY_R":             types.NewIdentGo("GLFW_KEY_R", "glfw.KeyR", keyT),
				"GLFW_KEY_S":             types.NewIdentGo("GLFW_KEY_S", "glfw.KeyS", keyT),
				"GLFW_KEY_T":             types.NewIdentGo("GLFW_KEY_T", "glfw.KeyT", keyT),
				"GLFW_KEY_U":             types.NewIdentGo("GLFW_KEY_U", "glfw.KeyU", keyT),
				"GLFW_KEY_V":             types.NewIdentGo("GLFW_KEY_V", "glfw.KeyV", keyT),
				"GLFW_KEY_W":             types.NewIdentGo("GLFW_KEY_W", "glfw.KeyW", keyT),
				"GLFW_KEY_X":             types.NewIdentGo("GLFW_KEY_X", "glfw.KeyX", keyT),
				"GLFW_KEY_Y":             types.NewIdentGo("GLFW_KEY_Y", "glfw.KeyY", keyT),
				"GLFW_KEY_Z":             types.NewIdentGo("GLFW_KEY_Z", "glfw.KeyZ", keyT),
				"GLFW_KEY_LEFT_BRACKET":  types.NewIdentGo("GLFW_KEY_LEFT_BRACKET", "glfw.KeyLeftBracket", keyT),
				"GLFW_KEY_BACKSLASH":     types.NewIdentGo("GLFW_KEY_BACKSLASH", "glfw.KeyBackslash", keyT),
				"GLFW_KEY_RIGHT_BRACKET": types.NewIdentGo("GLFW_KEY_RIGHT_BRACKET", "glfw.KeyRightBracket", keyT),
				"GLFW_KEY_GRAVE_ACCENT":  types.NewIdentGo("GLFW_KEY_GRAVE_ACCENT", "glfw.KeyGraveAccent", keyT),
				"GLFW_KEY_WORLD_1":       types.NewIdentGo("GLFW_KEY_WORLD_1", "glfw.KeyWorld1", keyT),
				"GLFW_KEY_WORLD_2":       types.NewIdentGo("GLFW_KEY_WORLD_2", "glfw.KeyWorld2", keyT),
				// function key constants
				"GLFW_KEY_ESCAPE":        types.NewIdentGo("GLFW_KEY_ESCAPE", "glfw.KeyEscape", keyT),
				"GLFW_KEY_ENTER":         types.NewIdentGo("GLFW_KEY_ENTER", "glfw.KeyEnter", keyT),
				"GLFW_KEY_TAB":           types.NewIdentGo("GLFW_KEY_TAB", "glfw.KeyTab", keyT),
				"GLFW_KEY_BACKSPACE":     types.NewIdentGo("GLFW_KEY_BACKSPACE", "glfw.KeyBackspace", keyT),
				"GLFW_KEY_INSERT":        types.NewIdentGo("GLFW_KEY_INSERT", "glfw.KeyInsert", keyT),
				"GLFW_KEY_DELETE":        types.NewIdentGo("GLFW_KEY_DELETE", "glfw.KeyDelete", keyT),
				"GLFW_KEY_RIGHT":         types.NewIdentGo("GLFW_KEY_RIGHT", "glfw.KeyRight", keyT),
				"GLFW_KEY_LEFT":          types.NewIdentGo("GLFW_KEY_LEFT", "glfw.KeyLeft", keyT),
				"GLFW_KEY_DOWN":          types.NewIdentGo("GLFW_KEY_DOWN", "glfw.KeyDown", keyT),
				"GLFW_KEY_UP":            types.NewIdentGo("GLFW_KEY_UP", "glfw.KeyUp", keyT),
				"GLFW_KEY_PAGE_UP":       types.NewIdentGo("GLFW_KEY_PAGE_UP", "glfw.KeyPageUp", keyT),
				"GLFW_KEY_PAGE_DOWN":     types.NewIdentGo("GLFW_KEY_PAGE_DOWN", "glfw.KeyPageDown", keyT),
				"GLFW_KEY_HOME":          types.NewIdentGo("GLFW_KEY_HOME", "glfw.KeyHome", keyT),
				"GLFW_KEY_END":           types.NewIdentGo("GLFW_KEY_END", "glfw.KeyEnd", keyT),
				"GLFW_KEY_CAPS_LOCK":     types.NewIdentGo("GLFW_KEY_CAPS_LOCK", "glfw.KeyCapsLock", keyT),
				"GLFW_KEY_SCROLL_LOCK":   types.NewIdentGo("GLFW_KEY_SCROLL_LOCK", "glfw.KeyScrollLock", keyT),
				"GLFW_KEY_NUM_LOCK":      types.NewIdentGo("GLFW_KEY_NUM_LOCK", "glfw.KeyNumLock", keyT),
				"GLFW_KEY_PRINT_SCREEN":  types.NewIdentGo("GLFW_KEY_PRINT_SCREEN", "glfw.KeyPrintScreen", keyT),
				"GLFW_KEY_PAUSE":         types.NewIdentGo("GLFW_KEY_PAUSE", "glfw.KeyPause", keyT),
				"GLFW_KEY_F1":            types.NewIdentGo("GLFW_KEY_F1", "glfw.KeyF1", keyT),
				"GLFW_KEY_F2":            types.NewIdentGo("GLFW_KEY_F2", "glfw.KeyF2", keyT),
				"GLFW_KEY_F3":            types.NewIdentGo("GLFW_KEY_F3", "glfw.KeyF3", keyT),
				"GLFW_KEY_F4":            types.NewIdentGo("GLFW_KEY_F4", "glfw.KeyF4", keyT),
				"GLFW_KEY_F5":            types.NewIdentGo("GLFW_KEY_F5", "glfw.KeyF5", keyT),
				"GLFW_KEY_F6":            types.NewIdentGo("GLFW_KEY_F6", "glfw.KeyF6", keyT),
				"GLFW_KEY_F7":            types.NewIdentGo("GLFW_KEY_F7", "glfw.KeyF7", keyT),
				"GLFW_KEY_F8":            types.NewIdentGo("GLFW_KEY_F8", "glfw.KeyF8", keyT),
				"GLFW_KEY_F9":            types.NewIdentGo("GLFW_KEY_F9", "glfw.KeyF9", keyT),
				"GLFW_KEY_F10":           types.NewIdentGo("GLFW_KEY_F10", "glfw.KeyF10", keyT),
				"GLFW_KEY_F11":           types.NewIdentGo("GLFW_KEY_F11", "glfw.KeyF11", keyT),
				"GLFW_KEY_F12":           types.NewIdentGo("GLFW_KEY_F12", "glfw.KeyF12", keyT),
				"GLFW_KEY_F13":           types.NewIdentGo("GLFW_KEY_F13", "glfw.KeyF13", keyT),
				"GLFW_KEY_F14":           types.NewIdentGo("GLFW_KEY_F14", "glfw.KeyF14", keyT),
				"GLFW_KEY_F15":           types.NewIdentGo("GLFW_KEY_F15", "glfw.KeyF15", keyT),
				"GLFW_KEY_F16":           types.NewIdentGo("GLFW_KEY_F16", "glfw.KeyF16", keyT),
				"GLFW_KEY_F17":           types.NewIdentGo("GLFW_KEY_F17", "glfw.KeyF17", keyT),
				"GLFW_KEY_F18":           types.NewIdentGo("GLFW_KEY_F18", "glfw.KeyF18", keyT),
				"GLFW_KEY_F19":           types.NewIdentGo("GLFW_KEY_F19", "glfw.KeyF19", keyT),
				"GLFW_KEY_F20":           types.NewIdentGo("GLFW_KEY_F20", "glfw.KeyF20", keyT),
				"GLFW_KEY_F21":           types.NewIdentGo("GLFW_KEY_F21", "glfw.KeyF21", keyT),
				"GLFW_KEY_F22":           types.NewIdentGo("GLFW_KEY_F22", "glfw.KeyF22", keyT),
				"GLFW_KEY_F23":           types.NewIdentGo("GLFW_KEY_F23", "glfw.KeyF23", keyT),
				"GLFW_KEY_F24":           types.NewIdentGo("GLFW_KEY_F24", "glfw.KeyF24", keyT),
				"GLFW_KEY_F25":           types.NewIdentGo("GLFW_KEY_F25", "glfw.KeyF25", keyT),
				"GLFW_KEY_KP_0":          types.NewIdentGo("GLFW_KEY_KP_0", "glfw.KeyKP0", keyT),
				"GLFW_KEY_KP_1":          types.NewIdentGo("GLFW_KEY_KP_1", "glfw.KeyKP1", keyT),
				"GLFW_KEY_KP_2":          types.NewIdentGo("GLFW_KEY_KP_2", "glfw.KeyKP2", keyT),
				"GLFW_KEY_KP_3":          types.NewIdentGo("GLFW_KEY_KP_3", "glfw.KeyKP3", keyT),
				"GLFW_KEY_KP_4":          types.NewIdentGo("GLFW_KEY_KP_4", "glfw.KeyKP4", keyT),
				"GLFW_KEY_KP_5":          types.NewIdentGo("GLFW_KEY_KP_5", "glfw.KeyKP5", keyT),
				"GLFW_KEY_KP_6":          types.NewIdentGo("GLFW_KEY_KP_6", "glfw.KeyKP6", keyT),
				"GLFW_KEY_KP_7":          types.NewIdentGo("GLFW_KEY_KP_7", "glfw.KeyKP7", keyT),
				"GLFW_KEY_KP_8":          types.NewIdentGo("GLFW_KEY_KP_8", "glfw.KeyKP8", keyT),
				"GLFW_KEY_KP_9":          types.NewIdentGo("GLFW_KEY_KP_9", "glfw.KeyKP9", keyT),
				"GLFW_KEY_KP_DECIMAL":    types.NewIdentGo("GLFW_KEY_KP_DECIMAL", "glfw.KeyKPDecimal", keyT),
				"GLFW_KEY_KP_DIVIDE":     types.NewIdentGo("GLFW_KEY_KP_DIVIDE", "glfw.KeyKPDivide", keyT),
				"GLFW_KEY_KP_MULTIPLY":   types.NewIdentGo("GLFW_KEY_KP_MULTIPLY", "glfw.KeyKPMultiply", keyT),
				"GLFW_KEY_KP_SUBTRACT":   types.NewIdentGo("GLFW_KEY_KP_SUBTRACT", "glfw.KeyKPSubtract", keyT),
				"GLFW_KEY_KP_ADD":        types.NewIdentGo("GLFW_KEY_KP_ADD", "glfw.KeyKPAdd", keyT),
				"GLFW_KEY_KP_ENTER":      types.NewIdentGo("GLFW_KEY_KP_ENTER", "glfw.KeyKPEnter", keyT),
				"GLFW_KEY_KP_EQUAL":      types.NewIdentGo("GLFW_KEY_KP_EQUAL", "glfw.KeyKPEqual", keyT),
				"GLFW_KEY_LEFT_SHIFT":    types.NewIdentGo("GLFW_KEY_LEFT_SHIFT", "glfw.KeyLeftShift", keyT),
				"GLFW_KEY_LEFT_CONTROL":  types.NewIdentGo("GLFW_KEY_LEFT_CONTROL", "glfw.KeyLeftControl", keyT),
				"GLFW_KEY_LEFT_ALT":      types.NewIdentGo("GLFW_KEY_LEFT_ALT", "glfw.KeyLeftAlt", keyT),
				"GLFW_KEY_LEFT_SUPER":    types.NewIdentGo("GLFW_KEY_LEFT_SUPER", "glfw.KeyLeftSuper", keyT),
				"GLFW_KEY_RIGHT_SHIFT":   types.NewIdentGo("GLFW_KEY_RIGHT_SHIFT", "glfw.KeyRightShift", keyT),
				"GLFW_KEY_RIGHT_CONTROL": types.NewIdentGo("GLFW_KEY_RIGHT_CONTROL", "glfw.KeyRightControl", keyT),
				"GLFW_KEY_RIGHT_ALT":     types.NewIdentGo("GLFW_KEY_RIGHT_ALT", "glfw.KeyRightAlt", keyT),
				"GLFW_KEY_RIGHT_SUPER":   types.NewIdentGo("GLFW_KEY_RIGHT_SUPER", "glfw.KeyRightSuper", keyT),
				"GLFW_KEY_MENU":          types.NewIdentGo("GLFW_KEY_MENU", "glfw.KeyMenu", keyT),
				"GLFW_KEY_LAST":          types.NewIdentGo("GLFW_KEY_LAST", "glfw.KeyLast", keyT),
				// mouse buttons
				"GLFW_MOUSE_BUTTON_1":      types.NewIdentGo("GLFW_MOUSE_BUTTON_1", "glfw.MouseButton1", mouseButtonT),
				"GLFW_MOUSE_BUTTON_2":      types.NewIdentGo("GLFW_MOUSE_BUTTON_2", "glfw.MouseButton2", mouseButtonT),
				"GLFW_MOUSE_BUTTON_3":      types.NewIdentGo("GLFW_MOUSE_BUTTON_3", "glfw.MouseButton3", mouseButtonT),
				"GLFW_MOUSE_BUTTON_4":      types.NewIdentGo("GLFW_MOUSE_BUTTON_4", "glfw.MouseButton4", mouseButtonT),
				"GLFW_MOUSE_BUTTON_5":      types.NewIdentGo("GLFW_MOUSE_BUTTON_5", "glfw.MouseButton5", mouseButtonT),
				"GLFW_MOUSE_BUTTON_6":      types.NewIdentGo("GLFW_MOUSE_BUTTON_6", "glfw.MouseButton6", mouseButtonT),
				"GLFW_MOUSE_BUTTON_7":      types.NewIdentGo("GLFW_MOUSE_BUTTON_7", "glfw.MouseButton7", mouseButtonT),
				"GLFW_MOUSE_BUTTON_8":      types.NewIdentGo("GLFW_MOUSE_BUTTON_8", "glfw.MouseButton8", mouseButtonT),
				"GLFW_MOUSE_BUTTON_LAST":   types.NewIdentGo("GLFW_MOUSE_BUTTON_LAST", "glfw.MouseButtonLast", mouseButtonT),
				"GLFW_MOUSE_BUTTON_LEFT":   types.NewIdentGo("GLFW_MOUSE_BUTTON_LEFT", "glfw.MouseButtonLeft", mouseButtonT),
				"GLFW_MOUSE_BUTTON_RIGHT":  types.NewIdentGo("GLFW_MOUSE_BUTTON_RIGHT", "glfw.MouseButtonRight", mouseButtonT),
				"GLFW_MOUSE_BUTTON_MIDDLE": types.NewIdentGo("GLFW_MOUSE_BUTTON_MIDDLE", "glfw.MouseButtonMiddle", mouseButtonT),
				// modifier keys
				"GLFW_MOD_SHIFT":     types.NewIdentGo("GLFW_MOD_SHIFT", "glfw.ModShift", modKeyT),
				"GLFW_MOD_CONTROL":   types.NewIdentGo("GLFW_MOD_CONTROL", "glfw.ModControl", modKeyT),
				"GLFW_MOD_ALT":       types.NewIdentGo("GLFW_MOD_ALT", "glfw.ModAlt", modKeyT),
				"GLFW_MOD_SUPER":     types.NewIdentGo("GLFW_MOD_SUPER", "glfw.ModSuper", modKeyT),
				"GLFW_MOD_CAPS_LOCK": types.NewIdentGo("GLFW_MOD_CAPS_LOCK", "glfw.ModCapsLock", modKeyT),
				"GLFW_MOD_NUM_LOCK":  types.NewIdentGo("GLFW_MOD_NUM_LOCK", "glfw.ModNumLock", modKeyT),
				// input mode
				"GLFW_CURSOR":               types.NewIdentGo("GLFW_CURSOR", "glfw.Cursor", inputModeT),
				"GLFW_STICKY_KEYS":          types.NewIdentGo("GLFW_STICKY_KEYS", "glfw.StickyKeys", inputModeT),
				"GLFW_STICKY_MOUSE_BUTTONS": types.NewIdentGo("GLFW_STICKY_MOUSE_BUTTONS", "glfw.StickyMouseButtons", inputModeT),
				"GLFW_LOCK_KEY_MODS":        types.NewIdentGo("GLFW_LOCK_KEY_MODS", "glfw.LockKeyMods", inputModeT),
				"GLFW_RAW_MOUSE_MOTION":     types.NewIdentGo("GLFW_RAW_MOUSE_MOTION", "glfw.RawMouseMotion", inputModeT),
				// peripheral events
				"GLFW_CONNECTED":    types.NewIdentGo("GLFW_CONNECTED", "glfw.Connected", perEventT),
				"GLFW_DISCONNECTED": types.NewIdentGo("GLFW_DISCONNECTED", "glfw.Disconnected", perEventT),
				// hints
				"GLFW_FOCUSED":                  types.NewIdentGo("GLFW_FOCUSED", "glfw.Focused", hintT),
				"GLFW_ICONIFIED":                types.NewIdentGo("GLFW_ICONIFIED", "glfw.Iconified", hintT),
				"GLFW_RESIZABLE":                types.NewIdentGo("GLFW_RESIZABLE", "glfw.Resizable", hintT),
				"GLFW_VISIBLE":                  types.NewIdentGo("GLFW_VISIBLE", "glfw.Visible", hintT),
				"GLFW_DECORATED":                types.NewIdentGo("GLFW_DECORATED", "glfw.Decorated", hintT),
				"GLFW_AUTO_ICONIFY":             types.NewIdentGo("GLFW_AUTO_ICONIFY", "glfw.AutoIconify", hintT),
				"GLFW_FLOATING":                 types.NewIdentGo("GLFW_FLOATING", "glfw.Floating", hintT),
				"GLFW_MAXIMIZED":                types.NewIdentGo("GLFW_MAXIMIZED", "glfw.Maximized", hintT),
				"GLFW_CENTER_CURSOR":            types.NewIdentGo("GLFW_CENTER_CURSOR", "glfw.CenterCursor", hintT),
				"GLFW_TRANSPARENT_FRAMEBUFFER":  types.NewIdentGo("GLFW_TRANSPARENT_FRAMEBUFFER", "glfw.TransparentFramebuffer", hintT),
				"GLFW_HOVERED":                  types.NewIdentGo("GLFW_HOVERED", "glfw.Hovered", hintT),
				"GLFW_FOCUS_ON_SHOW":            types.NewIdentGo("GLFW_FOCUS_ON_SHOW", "glfw.FocusOnShow", hintT),
				"GLFW_MOUSE_PASSTHROUGH":        types.NewIdentGo("GLFW_MOUSE_PASSTHROUGH", "glfw.MousePassthrough", hintT),
				"GLFW_RED_BITS":                 types.NewIdentGo("GLFW_RED_BITS", "glfw.RedBits", hintT),
				"GLFW_GREEN_BITS":               types.NewIdentGo("GLFW_GREEN_BITS", "glfw.GreenBits", hintT),
				"GLFW_BLUE_BITS":                types.NewIdentGo("GLFW_BLUE_BITS", "glfw.BlueBits", hintT),
				"GLFW_ALPHA_BITS":               types.NewIdentGo("GLFW_ALPHA_BITS", "glfw.AlphaBits", hintT),
				"GLFW_DEPTH_BITS":               types.NewIdentGo("GLFW_DEPTH_BITS", "glfw.DepthBits", hintT),
				"GLFW_STENCIL_BITS":             types.NewIdentGo("GLFW_STENCIL_BITS", "glfw.StencilBits", hintT),
				"GLFW_ACCUM_RED_BITS":           types.NewIdentGo("GLFW_ACCUM_RED_BITS", "glfw.AccumRedBits", hintT),
				"GLFW_ACCUM_GREEN_BITS":         types.NewIdentGo("GLFW_ACCUM_GREEN_BITS", "glfw.AccumGreenBits", hintT),
				"GLFW_ACCUM_BLUE_BITS":          types.NewIdentGo("GLFW_ACCUM_BLUE_BITS", "glfw.AccumBlueBits", hintT),
				"GLFW_ACCUM_ALPHA_BITS":         types.NewIdentGo("GLFW_ACCUM_ALPHA_BITS", "glfw.AccumAlphaBits", hintT),
				"GLFW_AUX_BUFFERS":              types.NewIdentGo("GLFW_AUX_BUFFERS", "glfw.AuxBuffers", hintT),
				"GLFW_STEREO":                   types.NewIdentGo("GLFW_STEREO", "glfw.Stereo", hintT),
				"GLFW_SAMPLES":                  types.NewIdentGo("GLFW_SAMPLES", "glfw.Samples", hintT),
				"GLFW_SRGB_CAPABLE":             types.NewIdentGo("GLFW_SRGB_CAPABLE", "glfw.SrgbCapable", hintT),
				"GLFW_REFRESH_RATE":             types.NewIdentGo("GLFW_REFRESH_RATE", "glfw.RefreshRate", hintT),
				"GLFW_DOUBLEBUFFER":             types.NewIdentGo("GLFW_DOUBLEBUFFER", "glfw.Doublebuffer", hintT),
				"GLFW_CLIENT_API":               types.NewIdentGo("GLFW_CLIENT_API", "glfw.ClientApi", hintT),
				"GLFW_CONTEXT_VERSION_MAJOR":    types.NewIdentGo("GLFW_CONTEXT_VERSION_MAJOR", "glfw.ContextVersionMajor", hintT),
				"GLFW_CONTEXT_VERSION_MINOR":    types.NewIdentGo("GLFW_CONTEXT_VERSION_MINOR", "glfw.ContextVersionMinor", hintT),
				"GLFW_CONTEXT_REVISION":         types.NewIdentGo("GLFW_CONTEXT_REVISION", "glfw.ContextRevision", hintT),
				"GLFW_CONTEXT_ROBUSTNESS":       types.NewIdentGo("GLFW_CONTEXT_ROBUSTNESS", "glfw.ContextRobustness", hintT),
				"GLFW_OPENGL_FORWARD_COMPAT":    types.NewIdentGo("GLFW_OPENGL_FORWARD_COMPAT", "glfw.OpenglForwardCompat", hintT),
				"GLFW_CONTEXT_DEBUG":            types.NewIdentGo("GLFW_CONTEXT_DEBUG", "glfw.ContextDebug", hintT),
				"GLFW_OPENGL_DEBUG_CONTEXT":     types.NewIdentGo("GLFW_OPENGL_DEBUG_CONTEXT", "glfw.OpenglDebugContext", hintT),
				"GLFW_OPENGL_PROFILE":           types.NewIdentGo("GLFW_OPENGL_PROFILE", "glfw.OpenglProfile", hintT),
				"GLFW_CONTEXT_RELEASE_BEHAVIOR": types.NewIdentGo("GLFW_CONTEXT_RELEASE_BEHAVIOR", "glfw.ContextReleaseBehavior", hintT),
				"GLFW_CONTEXT_NO_ERROR":         types.NewIdentGo("GLFW_CONTEXT_NO_ERROR", "glfw.ContextNoError", hintT),
				"GLFW_CONTEXT_CREATION_API":     types.NewIdentGo("GLFW_CONTEXT_CREATION_API", "glfw.ContextCreationApi", hintT),
				"GLFW_SCALE_TO_MONITOR":         types.NewIdentGo("GLFW_SCALE_TO_MONITOR", "glfw.ScaleToMonitor", hintT),
				"GLFW_COCOA_RETINA_FRAMEBUFFER": types.NewIdentGo("GLFW_COCOA_RETINA_FRAMEBUFFER", "glfw.CocoaRetinaFramebuffer", hintT),
				"GLFW_COCOA_FRAME_NAME":         types.NewIdentGo("GLFW_COCOA_FRAME_NAME", "glfw.CocoaFrameName", hintT),
				"GLFW_COCOA_GRAPHICS_SWITCHING": types.NewIdentGo("GLFW_COCOA_GRAPHICS_SWITCHING", "glfw.CocoaGraphicsSwitching", hintT),
				"GLFW_X11_CLASS_NAME":           types.NewIdentGo("GLFW_X11_CLASS_NAME", "glfw.X11ClassName", hintT),
				"GLFW_X11_INSTANCE_NAME":        types.NewIdentGo("GLFW_X11_INSTANCE_NAME", "glfw.X11InstanceName", hintT),
				"GLFW_WIN32_KEYBOARD_MENU":      types.NewIdentGo("GLFW_WIN32_KEYBOARD_MENU", "glfw.Win32KeyboardMenu", hintT),
			},
		}
		l.Declare(
			// functions
			types.NewIdentGo("glfwInit", "glfw.Init", env.FuncTT(env.C().Int(), nil)),
			types.NewIdentGo("glfwTerminate", "glfw.Terminate", env.FuncTT(nil, nil)),
			types.NewIdentGo("glfwCreateWindow", "glfw.CreateWindow", env.FuncTT(env.PtrT(windowT), env.Go().Int(), env.Go().Int(), env.Go().String(), env.PtrT(monitorT), env.PtrT(windowT))),
			types.NewIdentGo("glfwGetProcAddress", "glfw.GetProcAddress", env.FuncTT(env.PtrT(nil), env.Go().String())),
			types.NewIdentGo("glfwPollEvents", "glfw.PollEvents", env.FuncTT(nil, nil)),
			types.NewIdentGo("glfwSwapInterval", "glfw.SwapInterval", env.FuncTT(nil, env.Go().Int())),
			types.NewIdentGo("glfwGetTime", "glfw.GetTime", env.FuncTT(env.C().Float(), nil)),
			types.NewIdentGo("glfwGetCurrentContext", "glfw.GetCurrentContext", env.FuncTT(env.PtrT(windowT), nil)),
			types.NewIdentGo("glfwSetClipboardString", "glfw.SetClipboardString", env.FuncTT(nil, env.Go().String())),
			types.NewIdentGo("glfwSetJoystickCallback", "glfw.SetJoystickCallback", env.FuncTT(joystickCb, joystickCb)),
			types.NewIdentGo("glfwSetErrorCallback", "glfw.SetErrorCallback", env.FuncTT(errorCb, errorCb)),
			types.NewIdentGo("glfwWaitEvents", "glfw.WaitEvents", env.FuncTT(nil, nil)),
			types.NewIdentGo("glfwSetMonitorCallback", "glfw.SetMonitorCallback", env.FuncTT(monitorCb, monitorCb)),
			types.NewIdentGo("glfwGetPrimaryMonitor", "glfw.GetPrimaryMonitor", env.FuncTT(env.PtrT(monitorT), nil)),
			types.NewIdentGo("glfwGetMonitors", "glfw.GetMonitors", env.FuncTT(env.PtrT(env.PtrT(monitorT)), env.PtrT(env.Go().Int()))),
		)
		return l
	})
}
