package sdlHelper

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle  = "OpenGL Shader"
	winWidth  = 640
	winHeight = 480
)

var window sdl.Window
var windowGLContext sdl.GLContext
var windowInitialized = false

func initWindow() {
	var window *sdl.Window
	//var event sdl.Event
	//var running bool
	var err error
	runtime.LockOSThread()
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}

	windowGLContext, err = sdl.GL_CreateContext(window)
	if err != nil {
		panic(err)
	}

	gl.Init()
	gl.Enable(gl.TEXTURE_2D)
}

//Window will return the primary sdl window, will init the window if needed
func Window() *sdl.Window {
	if !windowInitialized {
		initWindow()

	}

	return &window
}

//DeleteWindow delets the sdl window and gl context. Defer this function
func DeleteWindow() {
	sdl.Quit()
	window.Destroy()
	sdl.GL_DeleteContext(windowGLContext)
}
